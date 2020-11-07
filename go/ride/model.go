package ride

import (
	"time"
)

type Rides struct {
	Id string `json:"id" xorm:"'id' pk"`

	startLat  double `json:"startLat"xorm:"'startLat'"`
	startLong double `json:"startLong"xorm:"'startLong'"`
	endLat    double `json:"endLat"xorm:"'endLat'"`
	endLong   double `json:"endLong"xorm:"'endLong'"`

	riderName     string    `json:"riderName"xorm:"'riderName'"`
	driverName    string    `json:"driverName"xorm:"'driverName'"`
	driverVehicle string    `json:"driverVehicle"xorm:"'driverVehicle'"`
	created       time.Time `json:"created"xorm:"'created'"`
}

func (m Model) TableName() string {
	return "Rides"
}
