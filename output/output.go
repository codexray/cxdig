package output

import (
	"codexray/cxdig/core"
	"codexray/cxdig/repos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// GetFilePath returns the full path of a generated output file
func GetFilePath(r repos.Repository, fileName string) (string, error) {
	outDir := getOutputDir(r)

	filePath, err := filepath.Abs(filepath.Join(outDir, fileName))
	core.DieOnError(err)

	// make sure the final absolute file path is located under the output dir
	if !strings.HasPrefix(filepath.Dir(filePath), outDir) {
		return "", errors.Errorf("invalid output file path: must be inside %s (%s)", outDir, filePath)
	}

	return filePath, nil
}

// MakeFilePath returns the full path of a generated output file and make sure its parent dir exists
func MakeFilePath(r repos.Repository, fileName string) (string, error) {
	filePath, err := GetFilePath(r, fileName)
	if err != nil {
		return "", err
	}

	dirPath := filepath.Dir(filePath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create '%s'", dirPath)
	}
	return filePath, nil
}

// FileExists returns true if the given file name exists in the output directory
func FileExists(r repos.Repository, fileName string) bool {
	filePath, err := GetFilePath(r, fileName)
	if err == nil {
		return core.FileExists(filePath)
	}
	return false
}

// WriteJSONFile saves the given data to a JSON file located in the output folder attached to the given repository
func WriteJSONFile(r repos.Repository, fileName string, obj interface{}) error {
	filePath, err := MakeFilePath(r, fileName)
	if err != nil {
		return err
	}

	core.Infof("Create JSON file '%s'", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON data")
	}
	_, err = f.Write(jsonData)
	if err != nil {
		return errors.Wrapf(err, "failed to write JSON data to file %s", filePath)
	}
	return nil
}

// ReadJSONFile reads data from the given JSON file located in the output folder attached to the given repository
func ReadJSONFile(r repos.Repository, fileName string, obj interface{}) error {
	filePath, err := GetFilePath(r, fileName)
	if err != nil {
		return nil
	}

	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, obj)
	if err != nil {
		return err
	}
	return nil
}

func getOutputDir(r repos.Repository) string {
	// IMPORANT: use the 'filepath' package and not 'path' to
	// be avoid surprises with '/' vs '\' (windows) and dir path
	// being replaced by '.' depending on the current dir

	// create the cxray dir in current dir except if we are inside the repository
	topDir, err := os.Getwd()
	core.DieOnError(errors.Wrap(err, "failed to get current working directory")) // not supposed to happen

	// cxray dir name to use
	cxdirName := fmt.Sprintf("%s.cxray", r.Name())

	// check we are not already in the cxray dir
	if filepath.Base(topDir) == cxdirName {
		return topDir
	}

	// are we inside the repo dir?
	if strings.HasPrefix(topDir, r.GetAbsPath()) {
		topDir = filepath.Dir(r.GetAbsPath())
	}

	return filepath.Join(topDir, cxdirName)
}
