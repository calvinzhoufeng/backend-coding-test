package ride

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Repository is the interface can be used in DI
type Repository interface {
	CreateRide(ride Ride) error
	GetRides(page int, pageSize int) ([]Ride, error)
	GetRideById(id string) (Ride, error)
}

// RepositoryImpl is the default implementation of Repository
type RepositoryImpl struct {
	DB *gorm.DB
}

// NewNewRepository is the constructor of RepositoryImpl
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

// CreateRide is to create a new Ride
func (r *RepositoryImpl) CreateRide(ride Ride) error {
	log.Debug().Msgf("to be added %v", ride)

	if dbc := r.DB.Create(&ride); dbc.Error != nil {
		return dbc.Error
	}

	log.Info().Msg("Insert ride successfully")
	return nil
}

// GetRides Get all rides from db per page number and pagesize
func (r *RepositoryImpl) GetRides(page int, pageSize int) ([]Ride, error) {
	var rides []Ride
	r.DB.Scopes(Paginate(page, pageSize)).Find(&rides)

	return rides, r.DB.Error
}

// GetRide by Id from db
func (r *RepositoryImpl) GetRideById(id string) (Ride, error) {
	var ride Ride
	r.DB.First(&ride, "id=?", id)
	return ride, r.DB.Error
}

// Paginate The generic pagination function
// @Param	page - optional - page number of the query
// 			pageSize - optional - page size of the query
// @Success object array
// @Failure DB.error
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
