package model

// plenty of things not modelled here that aren't useful to me

type PropertyPageModel struct {
	PropertyData PropertyData `json:"propertyData"`
}

type PropertyData struct {
	Id              string                      `json:"id"`
	Status          PropertyDataStatus          `json:"status"`
	Prices          PropertyDataPrices          `json:"prices"`
	Address         PropertyDataAddress         `json:"address"`
	KeyFeatures     []string                    `json:"keyFeatures"`
	Images          []PropertyDataImage         `json:"images"`
	Floorplans      []PropertyDataFloorplan     `json:"floorPlans"`
	Location        PropertyDataLocation        `json:"location"`
	NearestStations []PropertyDataStation       `json:"nearestStations"`
	Bedrooms        *int                        `json:"bedrooms"`
	Bathrooms       *int                        `json:"bathrooms"`
	TransactionType *string                     `json:"transactionType"` // enum?
	PropertySubType *string                     `json:"propertySubType"` // enum?
	SharedOwnership PropertyDataSharedOwnership `json:"sharedOwnership"`
}

type PropertyDataStatus struct {
	Published bool `json:"published"`
	Archived  bool `json:"archived"`
}

type PropertyDataPrices struct {
	PrimaryPrice          *string `json:"primaryPrice"`
	SecondaryPrice        *string `json:"secondaryPrice"`
	DisplayPriceQualifier *string `json:"displayPriceQualifier"`
	PricePerSqFt          *string `json:"pricePerSqFt"`
	Message               *string `json:"message"`
	ExchangeRate          *string `json:"exchangeRate"`
}

type PropertyDataAddress struct {
	DisplayAddress  *string `json:"displayAddress"`
	CountryCode     *string `json:"countryCode"`
	DeliveryPointId *string `json:"deliveryPointId"`
	UkCountry       *string `json:"ukCountry"`
	Outcode         *string `json:"outcode"`
	Incode          *string `json:"incode"`
}

type PropertyDataImage struct {
	Url     *string `json:"url"`
	Caption *string `json:"caption"`
}

type PropertyDataFloorplan struct {
	Url     *string `json:"url"`
	Caption *string `json:"caption"`
	Type    *string `json:"type"` // enum?
}

type PropertyDataLocation struct {
	Latitude  *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
	PinType   *string  `json:"pinType"` // enum?
}

type PropertyDataStation struct {
	Name     *string  `json:"name"`
	Types    []string `json:"types"` // []enum?
	Distance *float64 `json:"distance"`
	Unit     *string  `json:"unit"` // enum?
}

type PropertyDataSharedOwnership struct {
	SharedOwnershipFlag bool    `json:"sharedOwnershipFlag"`
	OwnershipPercentage *string `json:"ownershipPercentage"` // TODO look up the real type for this
	RentPrice           *string `json:"rentPrice"`           // TODO look up the real type for this
	RentFrequency       *string `json:"rentFrequency"`
}
