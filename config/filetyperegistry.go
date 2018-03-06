package config

import (
	"strings"

	"github.com/pkg/errors"

	"codexray/cxdig/core"
	"codexray/cxdig/types"
)

// FileTypeRegistry is used to find the type of a file from its name
type FileTypeRegistry struct {
	nameAndSuffixMap map[string]fileLanguageType
	prefixMap        map[string]fileLanguageType
}

// NewFileTypeRegistry create a new fileTypeRegistry
func NewFileTypeRegistry() *FileTypeRegistry {
	return &FileTypeRegistry{
		nameAndSuffixMap: make(map[string]fileLanguageType),
		prefixMap:        make(map[string]fileLanguageType),
	}
}

type fileLanguageType struct {
	Language types.LanguageID
	Type     types.FileType
}

// LoadJSONFile loads the types definition from a JSON file
func (r *FileTypeRegistry) LoadJSONFile(filePath string) error {
	var json []types.FileTypeInfo
	err := core.ReadJSONFile(filePath, &json)
	if err != nil {
		return err
	}

	for _, value := range json {
		// process fileNames
		for _, name := range value.FileName {
			if name == "" || strings.ToLower(name) != name {
				return errors.New("invalid fileName value: must be in lower case and not empty")
			}
			err := r.findPossibleKeyConflicts(name)
			if err != nil {
				return errors.Wrap(err, "conflict was found while creating fileTypes maps")
			}
			r.nameAndSuffixMap[strings.ToLower(name)] = fileLanguageType{value.Language, value.Type}
		}

		// process fileSuffixes
		for _, suffix := range value.FileSuffix {
			if suffix == "" || strings.ToLower(suffix) != suffix {
				return errors.New("invalid fileSuffix value: must be in lower case and not empty")
			}
			err := r.findPossibleKeyConflicts(suffix)
			if err != nil {
				return errors.Wrap(err, "conflict was found while creating fileTypes maps")
			}
			r.nameAndSuffixMap[strings.ToLower(suffix)] = fileLanguageType{value.Language, value.Type}
		}

		// process filePrefixes
		for _, prefix := range value.FilePrefix {
			if prefix == "" || strings.ToLower(prefix) != prefix {
				return errors.New("invalid filePrefix value: must be in lower case and not empty")
			}
			err := r.findPossibleKeyConflicts(prefix)
			if err != nil {
				return errors.Wrap(err, "conflict was found while creating fileTypes maps")
			}
			r.prefixMap[strings.ToLower(prefix)] = fileLanguageType{value.Language, value.Type}
		}
	}
	return r.findPossiblePrefixConflicts()
}

// GetFileTypeAndLanguage tries to identify the type and eventual language from a given file name
func (r *FileTypeRegistry) GetFileTypeAndLanguage(fileName string) (types.FileType, types.LanguageID) {
	fileName = strings.ToLower(fileName)

	// try first with file name
	if _, exist := r.nameAndSuffixMap[fileName]; exist {
		return types.FileType(r.nameAndSuffixMap[fileName].Type), types.LanguageID(r.nameAndSuffixMap[fileName].Language)
	}

	for key, value := range r.prefixMap {
		if strings.HasPrefix(fileName, key) {
			return types.FileType(value.Type), types.LanguageID(value.Language)
		}
	}

	for i := 0; i < len(fileName); i++ {
		if _, ok := r.nameAndSuffixMap[fileName[i:]]; ok {
			return types.FileType(r.nameAndSuffixMap[fileName[i:]].Type), types.LanguageID(r.nameAndSuffixMap[fileName[i:]].Language)
		}
	}

	return types.FileTypeUnknown, types.LanguageUnknown
}

func (r *FileTypeRegistry) findPossibleKeyConflicts(key string) error {
	if _, exist := r.nameAndSuffixMap[key]; exist {
		return errors.New("duplicate entry error on value " + key)
	}
	if _, exist := r.prefixMap[key]; exist {
		return errors.New("duplicate entry error on value " + key)
	}
	return nil
}

func (r *FileTypeRegistry) findPossiblePrefixConflicts() error {
	for key := range r.prefixMap {
		for key2 := range r.prefixMap {
			if (strings.HasPrefix(key, key2) || strings.HasPrefix(key2, key)) && key != key2 {
				return errors.Errorf("%s conflicts with %s", key, key2)
			}
		}
	}
	return nil
}
