package src

import (
	"os"
	"testing"
)

func TestLocal_ParsePropertyPageHTML_Succeeds(t *testing.T) {
	html, err := os.ReadFile("resources/property_page.html")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = ParsePropertyPageHTML(html)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLocal_ParseRentSearchPageHTML_Succeeds(t *testing.T) {
	html, err := os.ReadFile("resources/rent_search_page.html")
	if err != nil {
		t.Fatalf(err.Error())
	}

	res, err := ParseRentSearchPageHTML(html)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Logf("%v", res)
}
