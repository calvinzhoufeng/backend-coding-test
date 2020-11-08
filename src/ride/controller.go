package ride

import (
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type RideController struct {
	app  *iris.Application
	DB   *gorm.DB
	repo *Repository
}

func NewController(app *iris.Application, db *gorm.DB) *RideController {
	repository := NewRepository(db)

	c := &RideController{
		app:  app,
		DB:   db,
		repo: repository,
	}

	// List down all services exposed
	c.app.Get("/rides", c.GetRides)
	c.app.Get("/ride/{id}", c.GetRideById)

	c.app.Post("rides", c.AddRides)

	log.Info().Msg("Ride controller initialized successfully")

	return c
}

// GetRides Get all rides
// @Param
// @Success 200 []Rides
// @Failure 400
// @router /rides [get]
func (c *RideController) GetRides(ctx iris.Context) {
	page, _ := ctx.Params().GetInt("page")
	pageSize, _ := ctx.Params().GetInt("pageSize")

	rides, err := c.repo.GetRides(page, pageSize)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Response{
			Success: false,
			Data: Err{
				Code:    "SERVER_ERROR",
				Message: "Unknown error",
			},
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    rides,
	})
	return
}

// GetRideById  Get a ride by ID
// @Param rideId
// @Success 200 Ride
// @Failure 400
// @router /rides [get]
func (c *RideController) GetRideById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var ride Ride
	c.DB.First(&ride, "id=?", id)

	if ride.ID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Response{
			Success: false,
			Data: Err{
				Code:    "001",
				Message: "Invalid id",
			},
		})
		return
	}

	log.Debug().Int64("rideId", ride.ID).Msg("Found a ride")

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    "Healthy",
	})
	return
}

// AddRides Add a new ride
// @Param &Ride{}
// @Success 200 Ride
// @Failure 400
// @router /rides [post]
func (c *RideController) AddRides(ctx iris.Context) {
	var createRideRequest CreateRideRequest
	if err := ctx.ReadJSON(&createRideRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Response{
			Success: false,
			Data: Err{
				Code:    "002",
				Message: "Bad request body",
			},
		})
		return
	}

	// TODO: Other validations are skipped due to time constraits
	if createRideRequest.StartLat.LessThan(decimal.NewFromInt(-90)) || createRideRequest.EndLat.GreaterThan(decimal.NewFromInt(90)) {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Response{
			Success: false,
			Data: Err{
				Code:    "003",
				Message: "Start latitude and longitude must be between -90 - 90 and -180 to 180 degrees respectively",
			},
		})
		return
	}

	ride := Ride{
		StartLat:  createRideRequest.StartLat,
		EndLat:    createRideRequest.EndLat,
		StartLong: createRideRequest.StartLong,
		EndLong:   createRideRequest.EndLong,

		DriverName: createRideRequest.DriverName,
		RiderName:  createRideRequest.RiderName,
	}
	log.Debug().Msgf("controller to be added %v %s", ride, ride.RiderName)
	log.Debug().Msgf("controller to be added %s", ride.RiderName)

	c.repo.CreateRide(ride)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    "Healthy",
	})
	return
}
