package ride

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

type RideController struct {
	app *iris.Application
}

func NewRideController(app *iris.Application, db *gorm) {

	c := &RideController{
		app: app,
	}

	// List down all services exposed
	c.app.Get("/health", c.HealthCheck)
	c.app.Get("/rides", c.GetRides)
	c.app.Get("/rides/{id}", c.GetRideById)

	c.app.Post("rides", c.AddRides)

	log.Info().Msg("Ride controller initialized successfully")

	return c
}

/**
 * Get all rides
 * @method GET
 */
 func (c *RideController) GetRides(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&api.Response{
		Success: true,
		Data: "Healthy",
	})
	return
}

/**
 * Get a ride by Id
 * @method GET
 */
 func (c *RideController) GetRideById(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&api.Response{
		Success: true,
		Data: "Healthy",
	})
	return
}

/**
 * Add a ride
 * @method POST
 */
 func (c *RideController) AddRides(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&api.Response{
		Success: true,
		Data: "Healthy",
	})
	return
}

/**
 * Health check
 * @method GET
 */
func (c *RideController) HealthCheck(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&api.Response{
		Success: true,
		Data: "Healthy",
	})
	return
}