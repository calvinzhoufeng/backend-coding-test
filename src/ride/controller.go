package ride

import (
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"

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
	c.app.Get("/rides/{id}", c.GetRideById)

	c.app.Post("rides", c.AddRides)

	log.Info().Msg("Ride controller initialized successfully")

	return c
}

// @Title GetRides
// @Description Get all rides
// @Param
// @Success 200 []Rides
// @Failure 400
// @router /rides [get]
func (c *RideController) GetRides(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    "Healthy",
	})
	return
}

/**
 * Get a ride by Id
 * @method GET
 */
func (c *RideController) GetRideById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var ride Ride
	c.DB.First(&ride, "id=?", id)

	if ride.Id == 0 {
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

	log.Debug().Int64("rideId", ride.Id).Msg("Found a ride")

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    "Healthy",
	})
	return
}

/**
 * Add a ride
 * @method POST
 */
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
	// if createRideRequest.startLat < -90 || createRideRequest.endLat > 90 || crecreateRideRequest.startLong < -180 || createRideRequest.startLong > 180 {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	_, _ = ctx.JSON(&Response{
	// 		Success: false,
	// 		Data: Err{
	// 			Code:    "003",
	// 			Message: "Start latitude and longitude must be between -90 - 90 and -180 to 180 degrees respectively",
	// 		},
	// 	})
	// 	return
	// }

	ride := Ride{
		startLat:  createRideRequest.startLat,
		endLat:    createRideRequest.endLat,
		startLong: createRideRequest.startLong,
		endLong:   createRideRequest.endLong,

		driverName: createRideRequest.driverName,
		riderName:  createRideRequest.riderName,
	}

	c.repo.CreateRide(ride)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&Response{
		Success: true,
		Data:    "Healthy",
	})
	return
}
