package output

import (
	"codexray/cxdig/repos"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func NewXRayFilePath(r repos.Repository, fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	repoName := r.Name()
	cxRootDir := filepath.Join(cwd, repoName.String()+".cxray")

	filePath := filepath.Join(cxRootDir, fileName)
	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	// Make sure parentDir exists and is inside cwd
	dirPath := filepath.Dir(filePath)
	if !strings.HasPrefix(dirPath, cxRootDir) {
		return "", errors.Errorf("invalid output file path: must be inside %s (%s)", cxRootDir, filePath)
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
