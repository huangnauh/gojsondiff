package tests

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
)

func LoadFixture(file string) map[string]interface{} {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		Fail("Fixture file '" + file + "' not found.")
	}
	var result map[string]interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		Fail("Unmarshaling JSON of '" + file + "' failed: " + err.Error())
	}
	return result
}

func LoadFixtureAsArray(file string) []interface{} {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		Fail("Fixture file '" + file + "' not found.")
	}
	var result []interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		Fail("Unmarshaling JSON of '" + file + "' failed: " + err.Error())
	}
	return result
}
