package core

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSONFile(fileName string, obj interface{}) error {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, obj)
	if err != nil {
		return err
	}
	return nil
}
