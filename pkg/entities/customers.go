package entities

type Customers struct {
	ID int `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Postcode    string  `json:"postcode"`
	GeoLocation string  `json:"geo_location"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
