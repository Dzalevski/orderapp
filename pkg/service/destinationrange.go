package service

import (
	"djale/pkg/entities"
	"djale/pkg/store"
	"errors"
	geo "github.com/kellydunn/golang-geo"
)

type DestinationService interface {
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type DestinationServiceImpl struct {
	VanStore      store.VanStore
	OrderStore    store.OrderStore
	CustomerStore store.CustomerStore
}

// NewCustomerStore returns an instance of CustomerStore.
func NewDestinationService(vanStore store.VanStore, orderStore store.OrderStore, customerStore store.CustomerStore) *DestinationServiceImpl {
	return &DestinationServiceImpl{
		VanStore:      vanStore,
		OrderStore:    orderStore,
		CustomerStore: customerStore,
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

// for oder id u get slice of van ids that are okay for transferring the order
func (s *DestinationServiceImpl) WhichVanForCustomerOrder(orderID int) (*entities.VanRunResponse, error) {

	var validVans []int
	var resp entities.VanRunResponse

	order, err := s.OrderStore.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	customer, err := s.CustomerStore.GetCustomerByID(order.CustomerID)
	if err != nil {
		return nil, err
	}
	validVanID, err := s.GetDestinationBetweenTwoPoints(customer.Latitude, customer.Longitude)
	if err != nil {
		return nil, err
	}
	validVans = append(validVans, validVanID)

	resp.ConsID = order.ConsID
	resp.VanIDs = validVans

	return &resp, nil

}
