package src

import (
	"testing"

	"github.com/alpoi-x/rm-scrape/src/model"
)

func TestNonLocal_GetPropertyPage_AdHoc(t *testing.T) {
	s := CreateRightMoveScraper()

	_, err := s.GetPropertyPage("81948717")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNonLocal_GetRentSearchPage_AdHoc(t *testing.T) {
	s := CreateRightMoveScraper()
	q := &model.RentSearchParameters{
		LocationIdentifier:  "STATION%5E6332",
		Radius:              0.0,
		MinPrice:            500,
		MaxPrice:            1500,
		MinBedrooms:         model.AsPointer(0),
		MaxBedrooms:         model.AsPointer(1),
		DisplayPropertyType: "",
		MaxDaysSinceAdded:   "",
		IncludeLetAgreed:    false,
		MustHave:            []string{},
		DontShow:            []string{"houseShare", "retirement"},
		FurnishTypes:        []string{},
		PropertyTypes:       []string{},
	}

	_, err := s.GetRentSearchPage(q)
	if err != nil {
		t.Errorf(err.Error())
	}
}
