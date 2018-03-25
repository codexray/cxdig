package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FileExists returns true if the given file path exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

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
