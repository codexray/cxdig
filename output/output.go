package output

import (
	"codexray/cxdig/repos"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func NewXRayFilePath(r repos.Repository, fileName string) (string, error) {
	dirPath, filePath, err := getFilePath(r, fileName)
	if err != nil {
		return "", nil
	}

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func WriteJSONFile(r repos.Repository, fileName string, obj interface{}) error {
	filePath, err := NewXRayFilePath(r, fileName)
	if err != nil {
		return errors.Wrap(err, "path for the json could not be created")
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return errors.Wrap(err, "failed to save json in path "+filePath)
	}
	_, err = f.Write(jsonData)
	if err != nil {
		return errors.Wrap(err, "failed to write in json file")
	}
	return nil
}

func ReadJSONFile(r repos.Repository, fileName string, obj interface{}) error {
	_, filePath, err := getFilePath(r, fileName)
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

func CheckFileExistence(r repos.Repository, fileName string) (bool, error) {
	_, filePath, err := getFilePath(r, fileName)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(filePath)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func getFilePath(r repos.Repository, fileName string) (string, string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	if cwd == r.GetAbsPath() {
		cwd = path.Dir(cwd)
	}
	repoName := r.Name()
	cxRootDir := filepath.Join(cwd, repoName.String()+".cxray")

	filePath := filepath.Join(cxRootDir, fileName)
	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return "", "", err
	}

	// Make sure parentDir exists and is inside cwd
	dirPath := filepath.Dir(filePath)
	if !strings.HasPrefix(dirPath, cxRootDir) {
		return "", "", errors.Errorf("invalid output file path: must be inside %s (%s)", cxRootDir, filePath)
	}

	return dirPath, filePath, nil
}
