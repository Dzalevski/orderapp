package store

import (
	"database/sql"
	"djale/pkg/entities"
	"errors"
)

type VanStore interface {
	GetAllVans() ([]entities.Van, error)
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type VanStoreImpl struct {
	db *sql.DB
}

// NewCustomerStore returns an instance of CustomerStore.
func NewVanStore(db *sql.DB) *VanStoreImpl {
	return &VanStoreImpl{
		db: db,
	}
}

func (s *VanStoreImpl) GetAllVans() ([]entities.Van, error) {
	rows, err := s.db.Query("SELECT id,name,latitude,longitude FROM van ")

	if err != nil {
		return nil, errors.New("failed to get all vans")
	}

	defer rows.Close()

	var vans []entities.Van

	for rows.Next() {
		var v entities.Van
		if err := rows.Scan(&v.ID, &v.Name, &v.Latitude, &v.Longitude); err != nil {
			return nil, errors.New("failed to scan rows")
		}
		vans = append(vans, v)
	}

	return vans, nil
}
