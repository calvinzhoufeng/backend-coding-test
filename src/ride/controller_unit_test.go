package ride

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/shopspring/decimal"
)

type RepositoryMock struct{}

var ride = &Ride{
	StartLat:      decimal.NewFromInt(10),
	StartLong:     decimal.NewFromInt(1),
	EndLat:        decimal.NewFromInt(89),
	EndLong:       decimal.NewFromInt(81),
	RiderName:     "Jerry",
	DriverName:    "Tom",
	DriverVehicle: "SAL1234X",
}

func (r *RepositoryMock) CreateRide(ride Ride) error {
	return nil
}

func (r *RepositoryMock) GetRides(page int, pageSize int) ([]Ride, error) {
	return []Ride{*ride}, nil
}

func (r *RepositoryMock) GetRideById(id string) (Ride, error) {
	return *ride, nil
}

func TestController(t *testing.T) {
	repositoryMock := &RepositoryMock{}

	controller := NewController(iris.New(), repositoryMock)
	e := httptest.New(t, controller.app)

	// redirects to /admin without basic auth
	e.POST("/rides").WithJSON(ride).Expect().Status(httptest.StatusOK)

	obj := e.GET("/rides").Expect().Status(httptest.StatusOK).JSON().Object()

	obj.ContainsKey("data")
	obj.Value("data").Array().Element(0)
	t.Logf("returned data %v\n", obj.Value("data").Array().Element(0).Object())

	obj.Value("data").Array().Element(0).Object().ContainsKey("RiderName")
}
