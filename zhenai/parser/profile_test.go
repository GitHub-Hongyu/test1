package parser

import (
	"testing"
	"crawler/model"
	"io/ioutil"

)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "情深义重 ")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was &v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:   "情深义重",
		Age:    45,
		Height: 168,
	}
	if profile != expected {
		t.Errorf("expected %v: but was %v", expected, result.Items[0])
	}
}
