package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const (
		resultSize = 470
	)

	if len(result.Requests) != resultSize {
		t.Errorf("result shoud have %d request, but had %d request",
			resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result shoud have %d request, but had %d request",
			resultSize, len(result.Items))
	}
}
