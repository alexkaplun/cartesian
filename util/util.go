package util

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/alexkaplun/cartesian/model"
)

// Load the pints data from a json file into the PointList model
func LoadPointListFromCsv(fileName string) (*model.PointList, error) {
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var pointList model.PointList
	if err = json.Unmarshal(data, &pointList); err != nil {
		return nil, err
	}
	return &pointList, nil
}
