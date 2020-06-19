package entities

type Customers struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Postcode string `json:"postcode"`
	GeoLocation string `json:"geo_location"`
}
