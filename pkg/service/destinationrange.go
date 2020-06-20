package service

import (
	"djale/pkg/store"
	"errors"
	geo "github.com/kellydunn/golang-geo"
)

type DestinationService interface {
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type DestinationServiceImpl struct {
	VanStore store.VanStore
}

// NewCustomerStore returns an instance of CustomerStore.
func NewDestinationService(vanStore store.VanStore) *DestinationServiceImpl {
	return &DestinationServiceImpl{
		VanStore: vanStore,
	}
}

// retrieves all vans  from db and calculate the min distance between the vans point and requested point and returns which van ID should take the ride.
func (s *DestinationServiceImpl) GetDestinationBetweenTwoPoints(latitude float64, longitude float64) (int, error) {
	vans, err := s.VanStore.GetAllVans()
	if err != nil {
		return 0, errors.New("failed to get all vans")
	}

	vanID := 0

	p := geo.NewPoint(latitude, longitude)

	// this will give us the first vanID which will have 500km radius of requested destination and destination that goes to
	for _, v := range vans {
		p2 := geo.NewPoint(v.Latitude, v.Longitude)
		dist := p.GreatCircleDistance(p2)
		if dist < 500 {
			vanID = v.ID
		}
	}
	return vanID, nil

}
