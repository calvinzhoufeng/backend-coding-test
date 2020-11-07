package main

import (
	"fmt"

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

	db.AutoMigrate(&Ride{})

	// Register APIs
	NewController(app, db)

	_ = app.Run(iris.Addr(fmt.Sprintf("%s:%d", "localhost", 8010)))
}

func New() *iris.Application {
	app := iris.New()
	// app.Use(recover.New())
	// handel request log
	// reqLogger := requestLogger.New(requestLogger.Config{
	// 	Status:             true,
	// 	IP:                 true,
	// 	Method:             true,
	// 	Path:               true,
	// 	Query:              false,
	// 	Columns:            false,
	// 	MessageContextKeys: nil,
	// 	MessageHeaderKeys:  nil,
	// 	LogFunc:            logger.RequestLogFunc,
	// 	Skippers:           nil,
	// })
	// app.Use(reqLogger)

	// request rate limit
	// if config.RateLimit != 0 {
	// 	app.Use(ratelimit.NewRateLimitHandler(config.RateLimit))
	// }

	// base metrics export setup
	// metrics := appmetric.New()
	// app.Use(metrics.ServeHTTP)
	// if config.PromPassword != "" {
	// 	log.Debug().Str("username", "prom").
	// 		Str("password", config.PromPassword).Msg("/metrics protected by basic auth")
	// 	promAuth := basicauth.New(basicauth.Config{
	// 		Users:   map[string]string{"prom": config.PromPassword},
	// 		Realm:   "Authorization Required", // defaults to "Authorization Required"
	// 		Expires: time.Duration(30) * time.Minute,
	// 	})
	// 	app.Get("/metrics", promAuth, iris.FromStd(promhttp.Handler()))
	// } else {
	// 	app.Get("/metrics", iris.FromStd(promhttp.Handler()))
	// }

	// handel error
	// app.OnAnyErrorCode(reqLogger, metrics.ServeHTTP)
	// handel iris internal log
	// app.Logger().Handle(logger.IrisLoggerHandler)
	// Turn on iris debug log
	// if config.IrisDebug == true {
	// 	app.Logger().SetLevel("debug")
	// }
	// default router
	app.Get("/health", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"message": "Healthy",
		})
	})

	// CORS setup
	// crs := cors.AllowAll()
	// app.AllowMethods(iris.MethodOptions)
	// app.Use(crs)

	return app
}
