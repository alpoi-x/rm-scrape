package model

import (
	"errors"
	"fmt"
	"strings"
)

var MaxDaysSinceAdded []string = []string{"", "1", "3", "7", "14"}

var DisplayPropertyTypes []string = []string{
	"",
	"houses",
	"flats",
	"bungalows",
	"land",
	"commercial",
	"other",
}

var MustHave []string = []string{
	"garden",
	"parking",
	"houseShare",
	"retirement",
	"student",
}

var DontShow []string = []string{
	"houseShare",
	"retirement",
	"student",
}

var FurnishTypes []string = []string{
	"furnished",
	"partFurnished",
	"unfurnished",
}

var PropertyTypes []string = []string{
	"bungalow",
	"detached",
	"flat",
	"land",
	"park-home",
	"private-halls",
	"semi-detached",
	"terraced",
}

func Contains[T comparable](list []T, val T) bool {
	for _, lv := range list {
		if lv == val {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](list []T, other []T) bool {
	for _, ov := range other {
		if !Contains(list, ov) {
			return false
		}
	}
	return true
}

func ConvertRadius(r float32) string {
	switch {
	case r > 40.0:
		return "40.0"
	case r > 30.0:
		return "30.0"
	case r > 20.0:
		return "20.0"
	case r > 15.0:
		return "15.0"
	case r > 10.0:
		return "10.0"
	case r > 5.0:
		return "5.0"
	case r > 3.0:
		return "3.0"
	case r > 1.0:
		return "1.0"
	case r > 0.5:
		return "0.5"
	case r > 0.25:
		return "0.25"
	default:
		return "0.0"
	}
}

type RentSearchParameters struct {
	LocationIdentifier  string
	Radius              float32
	MinPrice            int
	MaxPrice            int
	MinBedrooms         *int
	MaxBedrooms         *int
	DisplayPropertyType string
	MaxDaysSinceAdded   string
	IncludeLetAgreed    bool
	MustHave            []string
	DontShow            []string
	FurnishTypes        []string
	PropertyTypes       []string
}

func (r RentSearchParameters) ValidateAndBuildQueryString() (*string, error) {
	// error if minPrice is more than maxPrice
	if r.MinPrice > r.MaxPrice {
		return nil, errors.New("minPrice must be smaller than maxPrice")
	}

	// validate and convert minPrice for query
	var minPrice string
	if r.MinPrice < 0 {
		return nil, errors.New("minPrice must be greater than zero")
	} else if r.MinPrice == 0 {
		minPrice = ""
	} else if r.MinPrice > 40000 {
		minPrice = "40000"
	} else {
		minPrice = fmt.Sprint(r.MinPrice)
	}

	// validate and convert maxPrice for query
	var maxPrice string
	if r.MaxPrice < 0 {
		return nil, errors.New("maxPrice must be greater than zero")
	} else if r.MaxPrice == 0 {
		maxPrice = ""
	} else if r.MaxPrice > 40000 {
		maxPrice = "40000"
	} else {
		maxPrice = fmt.Sprint(r.MaxPrice)
	}

	// validate and convert minBedrooms for query
	var minBedrooms string
	if r.MinBedrooms == nil {
		minBedrooms = ""
	} else if *r.MinBedrooms < 0 {
		return nil, errors.New("minBedrooms, if specified, must be greater than zero")
	} else if *r.MinBedrooms > 10 {
		minBedrooms = "10"
	} else {
		minBedrooms = fmt.Sprint(*r.MinBedrooms)
	}

	// validate and convert maxBedroom for query
	var maxBedrooms string
	if r.MaxBedrooms == nil {
		maxBedrooms = ""
	} else if *r.MaxBedrooms < 0 {
		return nil, errors.New("maxBedrooms, if specified, must be greater than zero")
	} else if *r.MaxBedrooms > 10 {
		maxBedrooms = "10"
	} else {
		maxBedrooms = fmt.Sprint(*r.MaxBedrooms)
	}

	// validate displayPropertyType for query
	var displayPropertyType string
	if !Contains(DisplayPropertyTypes, r.DisplayPropertyType) {
		return nil, errors.New(
			fmt.Sprintf(
				"displayPropertyType \"%s\" invalid. Valid values are: \"%s\"",
				r.DisplayPropertyType,
				strings.Join(DisplayPropertyTypes, "\", \""),
			),
		)
	} else {
		displayPropertyType = r.DisplayPropertyType
	}

	// validate maxDaysSinceAdded for query
	var maxDaysSinceAdded string
	if !Contains(MaxDaysSinceAdded, r.MaxDaysSinceAdded) {
		return nil, errors.New(
			fmt.Sprintf(
				"maxDaysSinceAdded \"%s\" invalid. Valid values are \"%s\"",
				r.MaxDaysSinceAdded,
				strings.Join(MaxDaysSinceAdded, "\", \""),
			),
		)
	} else {
		maxDaysSinceAdded = r.MaxDaysSinceAdded
	}

	// validate and join mustHave for query
	var mustHave string
	if len(r.MustHave) == 0 {
		mustHave = ""
	} else if !ContainsAll(MustHave, r.MustHave) {
		return nil, errors.New(
			fmt.Sprintf(
				"one of mustHave values is invalid. Valid values are \"%s\"",
				strings.Join(MustHave, "\", \""),
			),
		)
	} else {
		mustHave = strings.Join(r.MustHave, ",")
	}

	// validate and join dontShow for query
	var dontShow string
	if len(r.DontShow) == 0 {
		dontShow = ""
	} else if !ContainsAll(DontShow, r.DontShow) {
		return nil, errors.New(
			fmt.Sprintf(
				"one of dontShow values is invalid. Valid values are \"%s\"",
				strings.Join(DontShow, "\", \""),
			),
		)
	} else {
		dontShow = strings.Join(r.DontShow, ",")
	}

	// validate and join furnishTypes for query
	var furnishTypes string
	if len(r.FurnishTypes) == 0 {
		furnishTypes = ""
	} else if !ContainsAll(FurnishTypes, r.FurnishTypes) {
		return nil, errors.New(
			fmt.Sprintf(
				"one of furnishTypes values is invalid. Valid values are \"%s\"",
				strings.Join(FurnishTypes, "\", \""),
			),
		)
	} else {
		furnishTypes = strings.Join(r.FurnishTypes, ",")
	}

	// validate and join propertyTypes for query
	var propertyTypes string
	if len(r.PropertyTypes) == 0 {
		propertyTypes = ""
	} else if !ContainsAll(PropertyTypes, r.PropertyTypes) {
		return nil, errors.New(
			fmt.Sprintf(
				"one of propertyTypes values is invalid. Valid values are \"%s\"",
				strings.Join(PropertyTypes, "\", \""),
			),
		)
	} else {
		propertyTypes = strings.Join(r.PropertyTypes, ",")
	}

	// turn includeLetAgreed to string for query
	var includeLetAgreed string
	if r.IncludeLetAgreed {
		includeLetAgreed = "true"
	} else {
		includeLetAgreed = "false"
	}

	qp := []string{
		"searchType=RENT",
		"locationIdentifier=" + r.LocationIdentifier,
		"radius=" + ConvertRadius(r.Radius),
		"minPrice=" + minPrice,
		"maxPrice=" + maxPrice,
		"minBedrooms=" + minBedrooms,
		"maxBedrooms=" + maxBedrooms,
		"displayPropertyType=" + displayPropertyType,
		"maxDaysSinceAdded=" + maxDaysSinceAdded,
		"includeLetAgreed=" + includeLetAgreed,
		"mustHave=" + mustHave,
		"dontShow=" + dontShow,
		"furnishTypes=" + furnishTypes,
		"propertyTypes=" + propertyTypes,
	}

	qs := strings.Join(qp, "&")

	return &qs, nil
}
