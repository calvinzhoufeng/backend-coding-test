package ride

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Ride struct {
	gorm.Model
	Id int64 `json:"id" xorm:"'id' pk"`

	startLat  decimal.Decimal `json:"startLat"xorm:"'startLat'"`
	startLong decimal.Decimal `json:"startLong"xorm:"'startLong'"`
	endLat    decimal.Decimal `json:"endLat"xorm:"'endLat'"`
	endLong   decimal.Decimal `json:"endLong"xorm:"'endLong'"`

	riderName     string    `json:"riderName"xorm:"'riderName'"`
	driverName    string    `json:"driverName"xorm:"'driverName'"`
	driverVehicle string    `json:"driverVehicle"xorm:"'driverVehicle'"`
	created       time.Time `json:"created"xorm:"'created'"`
}

type CreateRideRequest struct {
	startLat  decimal.Decimal `json:"startLat"`
	startLong decimal.Decimal `json:"startLong"`
	endLat    decimal.Decimal `json:"endLat"`
	endLong   decimal.Decimal `json:"endLong"`

	riderName     string `json:"riderName"`
	driverName    string `json:"driverName"`
	driverVehicle string `json:"driverVehicle"`
}

/**
 * Note: this is a shortcut due limited time frame
 */
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type Err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
