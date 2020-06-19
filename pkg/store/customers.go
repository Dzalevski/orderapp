package store

import (
	"database/sql"
	"djale/pkg/entities"
	"errors"
	"log"
)

type CustomerStore interface {
	GetCustomerByID(id int) (*entities.Customers, error)
	DeleteCustomerByID(id int) error
	InsertCustomer(name string, address string, postcode string, geoLocation string) error
	UpdateCustomerByID(name string, address string, postocde string, geoLocation string, id int) error
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type CustomerStoreImpl struct {
	db *sql.DB
}

// NewCustomerStore returns an instance of CustomerStore.
func NewCustomerStore(db *sql.DB) *CustomerStoreImpl {
	return &CustomerStoreImpl{
		db: db,
	}
}

func (s *CustomerStoreImpl) GetCustomerByID(id int) (*entities.Customers, error) {

	var cust entities.Customers

	err := s.db.QueryRow("SELECT name,address, postcode,geo_location FROM customers WHERE id=$1",
		id).Scan(&cust.Name, &cust.Address, &cust.Postcode, &cust.GeoLocation)
	if err != nil {
		return nil, errors.New("failed to get customer by id")
	}
	return &cust, nil

}

func (s *CustomerStoreImpl) DeleteCustomerByID(id int) error {
	_, err := s.db.Exec("DELETE FROM customers WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete customer by id")
	}
	return nil
}

func (s *CustomerStoreImpl) InsertCustomer(name string, address string, postcode string, geoLocation string) error {
	_, err := s.db.Exec("INSERT INTO customers(name,address,postcode,geo_location) values($1,$2,$3,$4)", name, address, postcode, geoLocation)
	if err != nil {
		log.Println(err)
		return errors.New("failed to insert customer")
	}
	return nil
}

func (s *CustomerStoreImpl) UpdateCustomerByID(name string, address string, postocde string, geoLocation string, id int) error {
	_, err := s.db.Exec("UPDATE customers SET name=$1, address=$2 , postcode=$3,geo_location=$4 WHERE id=$5", name, address, postocde, geoLocation, id)
	if err != nil {
		return errors.New("failed to update customer by id")
	}
	return nil
}
