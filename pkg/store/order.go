package store

import (
	"database/sql"
	"djale/pkg/entities"
	"errors"
)

type OrderStore interface {
	GetAllOrders() ([]entities.Order, error)
	InsertOrder(customerID int, consID int) error
	GetOrderByID(id int) (*entities.Order, error)
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type OrderStoreImpl struct {
	db *sql.DB
}

// NewCustomerStore returns an instance of CustomerStore.
func NewOrderStore(db *sql.DB) *OrderStoreImpl {
	return &OrderStoreImpl{
		db: db,
	}
}

func (s *OrderStoreImpl) InsertOrder(customerID int, consID int) error {
	_, err := s.db.Exec("INSERT INTO order(customer_id,cons_id) values($1,$2)", customerID, consID)
	if err != nil {
		return errors.New("failed to insert consignments")
	}
	return nil
}

func (s *OrderStoreImpl) GetOrderByID(id int) (*entities.Order, error) {

	var o entities.Order

	err := s.db.QueryRow("SELECT customer_id, cons_id FROM order WHERE id=$1",
		id).Scan(&o.CustomerID, &o.ConsID)
	if err != nil {
		return nil, errors.New("failed to get customer by id")
	}
	return &o, nil

}

func (s *OrderStoreImpl) GetAllOrders() ([]entities.Order, error) {
	rows, err := s.db.Query("SELECT customer_id,cons_id FROM order ")

	if err != nil {
		return nil, errors.New("failed to get all vans")
	}

	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var o entities.Order
		if err := rows.Scan(&o.CustomerID, &o.ConsID); err != nil {
			return nil, errors.New("failed to scan rows")
		}
		orders = append(orders, o)
	}

	return orders, nil
}
