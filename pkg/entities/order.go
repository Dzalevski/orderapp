package entities

type Order struct {
	ID int `json:"id"`
	CustomerID int `json:"customer_id"`
	ConsID int `json:"cons_id"`
}

type VanRunResponse struct {
	ConsID int `json:"cons_id"`
	VanIDs []int `json:"van_id"`
}