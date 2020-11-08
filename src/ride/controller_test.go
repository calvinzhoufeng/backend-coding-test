package ride

import (
	"os"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// TestMain will exec each test, one by one
func TestMain(m *testing.M) {
	// exec setUp function
	setUp()
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// exec tearDown function
	tearDown()
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

// setUp function, add a number to numbers slice
func setUp() {
	if db == nil {
		db, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

		db.AutoMigrate(&Ride{})
	}
}

// tearDown function
func tearDown() {
}

func TestNewApp(t *testing.T) {

	controller := NewController(iris.New(), db)
	e := httptest.New(t, controller.app)

	ride := &Ride{
		StartLat:      decimal.NewFromInt(10),
		StartLong:     decimal.NewFromInt(1),
		EndLat:        decimal.NewFromInt(89),
		EndLong:       decimal.NewFromInt(81),
		RiderName:     "Jerry",
		DriverName:    "Tom",
		DriverVehicle: "SAL1234X",
	}
	// redirects to /admin without basic auth
	e.POST("/rides").WithJSON(ride).Expect().Status(httptest.StatusOK)

	obj := e.GET("/rides").Expect().Status(httptest.StatusOK).JSON().Object()

	obj.ContainsKey("data")
	obj.Value("data").Array().Element(0)
	t.Logf("returned data %v\n", obj.Value("data").Array().Element(0).Object())

	obj.Value("data").Array().Element(0).Object().ContainsKey("RiderName")
}
