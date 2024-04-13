package src

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alpoi-x/rm-scrape/src/model"
	"github.com/sirupsen/logrus"
)

const propertiesUri string = "https://www.rightmove.co.uk/properties/"
const propertyToRentUri = "https://www.rightmove.co.uk/property-to-rent/find.html"
const propertyForSaleUri = "https://www.rightmove.co.uk/property-for-sale/find.html"

const headerAccept string = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
const headerAcceptEncoding string = "gzip, deflate, br, zstd"
const headerAcceptLanguage string = "en-GB,en-US;q=0.9,en;q=0.8"
const headerReferer string = "https://google.co.uk"

type RightMoveScraper struct {
	client http.Client
	uar    *UserAgentRepository
}

func CreateRightMoveScraper() *RightMoveScraper {
	client := http.Client{}
	return &RightMoveScraper{
		client: client,
		uar:    CreateUserAgentRepository(),
	}
}

func (s *RightMoveScraper) AddHeaders(req *http.Request) *http.Request {
	req.Header.Add("User-Agent", s.uar.Random())
	req.Header.Add("Accept", headerAccept)
	req.Header.Add("Accept-Encoding", headerAcceptEncoding)
	req.Header.Add("Accept-Language", headerAcceptLanguage)
	req.Header.Add("Referer", headerReferer)
	return req
}

func (s *RightMoveScraper) GetPropertyPage(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", propertiesUri, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("error constructing http request: %s", err)
		return nil, err
	}

	s.AddHeaders(req)

	res, err := s.client.Do(req)
	if err != nil {
		logrus.Errorf("error fetching property page: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("error reading response body: %s", err)
		return nil, err
	}

	return body, nil
}

func (s *RightMoveScraper) GetRentSearchPage(params *model.RentSearchParameters) ([]byte, error) {
	qs, err := params.ValidateAndBuildQueryString()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?%s", propertyToRentUri, *qs)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("error constructing http request: %s", err)
		return nil, err
	}

	s.AddHeaders(req)

	res, err := s.client.Do(req)
	if err != nil {
		logrus.Errorf("error fetching rent search page: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("error reading response body: %s", err)
		return nil, err
	}

	return body, nil
}

// searchType=SALE
// &locationIdentifier=STATION%5E19814
// &insId=1
// &radius=0.0
// &minPrice=
// &maxPrice=
// &minBedrooms=
// &maxBedrooms=
// &displayPropertyType=
// &maxDaysSinceAdded=
// &_includeSSTC=on
// &newHome=
// &auction=false
// func (s *RightMoveScraper) GetSaleSearchPage()
