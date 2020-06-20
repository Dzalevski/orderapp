package store

import (
	"database/sql"
	"djale/pkg/entities"
	"errors"
	"log"
	"time"
)

type ConsignmentsStore interface {
	GetConsignmentByID(id int) (*entities.Consignment, error)
	DeleteConsignmentByID(id int) error
	InsertConsignment(barcode string, linkToSupplier string, ) error
	UpdateConsignmentByID(barcode string, linkToSupplier string, id int) error
}

// CustomerStoreImpl is the concrete implementation of the CustomerStore interface.
type ConsignmentsStoreImpl struct {
	db *sql.DB
}

// NewCustomerStore returns an instance of CustomerStore.
func NewConsignmentsStore(db *sql.DB) *ConsignmentsStoreImpl {
	return &ConsignmentsStoreImpl{
		db: db,
	}
}

func (s *ConsignmentsStoreImpl) GetConsignmentByID(id int) (*entities.Consignment, error) {

	var cons entities.Consignment

	err := s.db.QueryRow("SELECT barcode,link_to_supplier, returned_at FROM consignments WHERE id=$1",
		id).Scan(&cons.Barcode, &cons.LinkToSupplier, &cons.ReturnedAt)
	if err != nil {
		return nil, errors.New("failed to get consignments by id")
	}
	return &cons, nil

}

func (s *ConsignmentsStoreImpl) DeleteConsignmentByID(id int) error {
	_, err := s.db.Exec("DELETE FROM consignments WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete consignments by id")
	}
	return nil
}

func (s *ConsignmentsStoreImpl) InsertConsignment(barcode string, linkToSupplier string, ) error {
	_, err := s.db.Exec("INSERT INTO consignments(barcode,link_to_supplier,returned_at) values($1,$2,NOW())", barcode, linkToSupplier)
	if err != nil {
		log.Println(err)
		return errors.New("failed to insert consignments")
	}
	return nil
}

func (s *ConsignmentsStoreImpl) UpdateConsignmentByID(barcode string, linkToSupplier string, id int) error {
	_, err := s.db.Exec("UPDATE consignments SET barcode=$1, link_to_supplier=$2 , returned_at=$3 WHERE id=$4", barcode, linkToSupplier, time.Now(), id)
	if err != nil {
		return errors.New("failed to update consignments by id")
	}
	return nil
}
