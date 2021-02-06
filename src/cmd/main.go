package main

import (
	"fmt"
	"go/src/note"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Start app
	app := New()

	// Open a new connection to our sqlite database.
	// Note the db configuration are hardcoded due to limit timeframe
	// db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Silent),
	// })
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	err = db.Debug().AutoMigrate(&note.Note{}, &note.Tag{})
	if err != nil {
		fmt.Printf("Failed to migrate db %v\n", err)
	}

	// newNote := &note.Note{Content: "Test"}
	// tag := note.Tag{Name: "Hihi"}
	// newNote.Tags = []note.Tag{tag}

	// db.Debug().Create(&newNote)
	// n := db.Debug().First(&newNote)
	// fmt.Printf("note %v\n", n)

	repository := note.NewRepository(db)

	// Register APIs
	note.NewController(app, repository)

	_ = app.Run(iris.Addr(fmt.Sprintf("%s:%d", "localhost", 8000)))
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
