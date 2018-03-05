package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

// WriteJSONFile writes the given object to a json file
func WriteJSONFile(fileName string, obj interface{}) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return err
	}
	f.Write(jsonData)
	return nil
}
