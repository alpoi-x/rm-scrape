package src

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/alpoi-x/rm-scrape/src/model"
	"github.com/sirupsen/logrus"
)

// NB: these will fail if a closing </script> tag exists in the JSON
var propertyPageRegex *regexp.Regexp = regexp.MustCompile(`window\.PAGE_MODEL\s*=\s*({.*?})\s*<\/script>`)
var rentSearchPageRegex *regexp.Regexp = regexp.MustCompile(`<script>window\.jsonModel = ({"properties":.*})\s*<\/script>`)

func ParsePropertyPageHTML(html []byte) (*model.PropertyPageModel, error) {
	b := propertyPageRegex.FindSubmatch(html)
	if b == nil {
		err := errors.New("no regex match")
		logrus.Errorf("error finding `window.PAGE_MODEL` in property page html: %s", err)
		return nil, err
	}

	var prop model.PropertyPageModel
	if err := json.Unmarshal(b[1], &prop); err != nil {
		logrus.Errorf("error unmarshalling property page model: %s", err)
		return nil, err
	}

	return &prop, nil
}

func ParseRentSearchPageHTML(html []byte) (*model.RentSearchPageModel, error) {
	b := rentSearchPageRegex.FindSubmatch(html)
	if b == nil {
		err := errors.New("no regex match")
		logrus.Errorf("error finding `window.jsonModel` in rent search page html: %s", err)
		return nil, err
	}

	var prop model.RentSearchPageModel
	if err := json.Unmarshal(b[1], &prop); err != nil {
		logrus.Errorf("error unmarshalling rent search page model: %s", err)
		return nil, err
	}

	return &prop, nil
}
