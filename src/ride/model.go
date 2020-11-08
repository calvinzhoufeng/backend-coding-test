package ride

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Ride struct {
	gorm.Model
	ID int64 `gorm:"primaryKey;autoIncrement"`

	StartLat  decimal.Decimal `gorm:"type:decimal(7,6);"`
	StartLong decimal.Decimal `gorm:"type:decimal(7,6);"`
	EndLat    decimal.Decimal `gorm:"type:decimal(7,6);"`
	EndLong   decimal.Decimal `gorm:"type:decimal(7,6);"`

	RiderName     string
	DriverName    string
	DriverVehicle string
	Created       time.Time `gorm:"autoCreateTime"`
}

type CreateRideRequest struct {
	StartLat  decimal.Decimal `json:"startLat"`
	StartLong decimal.Decimal `json:"startLong"`
	EndLat    decimal.Decimal `json:"endLat"`
	EndLong   decimal.Decimal `json:"endLong"`

	RiderName     string `json:"riderName"`
	DriverName    string `json:"driverName"`
	DriverVehicle string `json:"driverVehicle"`
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
