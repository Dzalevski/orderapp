package entities

import "time"

type Consignment struct {
	Barcode        string    `json:"barcode"`
	LinkToSupplier string    `json:"link_to_supplier"`
	ReturnedAt     time.Time `json:"returned_at"`
}
