package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Local_TestValidateAndBuildQueryString_ValidInput_ExpectedOutput(t *testing.T) {
	q := RentSearchParameters{
		LocationIdentifier:  "STATION%5E6332",
		Radius:              0.0,
		MinPrice:            500,
		MaxPrice:            1500,
		MinBedrooms:         AsPointer(0),
		MaxBedrooms:         AsPointer(1),
		DisplayPropertyType: "flats",
		MaxDaysSinceAdded:   "",
		IncludeLetAgreed:    false,
		MustHave:            []string{},
		DontShow:            []string{"houseShare", "retirement"},
		FurnishTypes:        []string{},
		PropertyTypes:       []string{"flat"},
	}

	qs, err := q.ValidateAndBuildQueryString()
	if err != nil {
		t.Fatalf(err.Error())
	}

	ex := "searchType=RENT&locationIdentifier=STATION%5E6332&radius=0.0&minPrice=500&maxPrice=1500&minBedrooms=0&maxBedrooms=1&displayPropertyType=flats&maxDaysSinceAdded=&includeLetAgreed=false&mustHave=&dontShow=houseShare,retirement&furnishTypes=&propertyTypes=flat"
	assert.Equal(t, ex, *qs)
}
