package main

import (
	"fmt"
	"go/src/ride"

	"gorm.io/gorm"

	"github.com/kataras/iris"
	"gorm.io/driver/sqlite"
)

func main() {
	// Start app
	app := New()

	// Open a new connection to our sqlite database.
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	db.AutoMigrate(&ride.Ride{})

	repository := ride.NewRepository(db)

	// Register APIs
	ride.NewController(app, repository)

	_ = app.Run(iris.Addr(fmt.Sprintf("%s:%d", "localhost", 8010)))
}

// New Initiate the HTTP application
// @TODO This is a shortcut to initiate iris app, further customization and configurations are required
// @Param
// @Success 200 Healthy
// @Failure 400
// @router /rides [get]
func New() *iris.Application {
	app := iris.New()
	// default router
	app.Get("/health", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"message": "Healthy",
		})
	})

	return app
}
