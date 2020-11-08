package ride

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateRide(ride Ride) error {
	log.Debug().Msgf("to be added %v", ride)

	if dbc := r.DB.Create(&ride); dbc.Error != nil {
		return dbc.Error
	}

	log.Info().Msg("Insert ride successfully")
	return nil
}

// GetRides Get all rides from db per page number and pagesize
func (r *Repository) GetRides(page int, pageSize int) ([]Ride, error) {
	var rides []Ride
	r.DB.Scopes(Paginate(page, pageSize)).Find(&rides)

	return rides, r.DB.Error
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
