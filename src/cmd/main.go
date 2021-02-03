package main

import (
	"fmt"
	"go/src/note"

	"gorm.io/gorm"

	"github.com/kataras/iris"
	"gorm.io/driver/sqlite"
)

func main() {
	// Start app
	app := New()

	// Open a new connection to our sqlite database.
	// Note the db configuration are hardcoded due to limit timeframe
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

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
