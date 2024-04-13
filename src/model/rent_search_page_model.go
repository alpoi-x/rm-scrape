package model

// again, plenty of stuff not modelled here

type RentSearchPageModel struct {
	Properties  []RentSearchProperty `json:"properties"`
	ResultCount string               `json:"resultCount"`
}

type RentSearchProperty struct {
	Id             int                        `json:"id"`
	Bedrooms       int                        `json:"bedrooms"`
	Bathrooms      int                        `json:"bathrooms"`
	NumberOfImages int                        `json:"numberOfImages"`
	DisplayAddress string                     `json:"displayAddress"`
	Location       RentSearchPropertyLocation `json:"location"`
	Price          RentSearchPropertyPrice    `json:"price"`
}

type RentSearchPropertyLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type RentSearchPropertyPrice struct {
	Amount       int    `json:"amount"`
	Frequency    string `json:"frequency"`
	CurrencyCode string `json:"currencyCode"`
}
