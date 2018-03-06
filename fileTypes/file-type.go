package filestypes

import (
	"strings"

	"github.com/pkg/errors"

	"codexray/cxdig/core"
)

// FileType identifies the type of a file in the repository
type FileType string
type LanguageID string

func (id *LanguageID) String() string {
	return string(*id)
}
func (id *FileType) String() string {
	return string(*id)
}

// FileTypeJSON is the struct used to extract FileType Information from a JSON
type FileTypeJSON struct {
	Language   string   `json:"language"`
	Type       FileType `json:"type"`
	FileName   []string `json:"filesNames"`
	FilePrefix []string `json:"filesPrefixes"`
	FileSuffix []string `json:"filesSuffixes"`
}

type FileTypeInfos struct {
	Language string
	Type     FileType
}

const (
	// FileTypeSource tags a file as being source code in a given language
	FileTypeSource FileType = "Source"
	// FileTypeGenerator tags a file as being part of a source file generator tool
	FileTypeGenerator FileType = "Generator"
	// FileTypeBuildSystem tags a file as being part of a build system tool
	FileTypeBuildSystem FileType = "BuildSystem"
	// FileTypeEnvConfig tags a file as being used to configure some development tools
	FileTypeEnvConfig FileType = "EnvConfig"
	// FileTypeLicense tags is used to tag license files
	FileTypeLicense FileType = "License"
	// FileTypeUnknown tags unrecognized files
	FileTypeUnknown FileType = ""
	// LanguageUnknown is used to identify unknown languages
	LanguageUnknown LanguageID = ""
)

var (
	nameAndSuffixMap map[string]FileTypeInfos
	prefixMap        map[string]FileTypeInfos
)

func InitFileTypeMaps() error {
	var fileTypesJSON []FileTypeJSON
	err := core.ReadJSONFile("./filetypes.json", &fileTypesJSON)
	if err != nil {
		return err
	}

	nameAndSuffixMap = make(map[string]FileTypeInfos)
	prefixMap = make(map[string]FileTypeInfos)

	for _, value := range fileTypesJSON {
		var err error
		for _, name := range value.FileName {
			if strings.ToLower(name) == name {
				err = findPossibleKeyConflicts(name)
				if err != nil {
					return errors.Wrap(err, "conflict was found while creating fileTypes maps")
				}
				nameAndSuffixMap[strings.ToLower(name)] = FileTypeInfos{value.Language, value.Type}
			} else {
				return errors.New("filesNames need to be in lower case")
			}
		}
		for _, suffix := range value.FileSuffix {
			if strings.ToLower(suffix) == suffix {
				err = findPossibleKeyConflicts(suffix)
				if err != nil {
					return errors.Wrap(err, "conflict was found while creating fileTypes maps")
				}
				nameAndSuffixMap[strings.ToLower(suffix)] = FileTypeInfos{value.Language, value.Type}
			} else {
				return errors.New("filesSuffixes need to be in lower case")
			}
		}
		for _, prefix := range value.FilePrefix {
			if strings.ToLower(prefix) == prefix {
				err = findPossibleKeyConflicts(prefix)
				if err != nil {
					return errors.Wrap(err, "conflict was found while creating fileTypes maps")
				}
				prefixMap[strings.ToLower(prefix)] = FileTypeInfos{value.Language, value.Type}
			} else {
				return errors.New("filesPrefixes need to be in lower case")
			}
		}
	}
	return nil
}

// IdentifyFileTypeAndLanguage tries to identify the type and eventual language from a given file name
func IdentifyFileTypeAndLanguage(fileName string) (FileType, LanguageID) {
	fileName = strings.ToLower(fileName)

	// try first with file name
	if _, exist := nameAndSuffixMap[fileName]; exist {
		return FileType(nameAndSuffixMap[fileName].Type), LanguageID(nameAndSuffixMap[fileName].Language)
	}

	for key, value := range prefixMap {
		if strings.HasPrefix(fileName, key) {
			return FileType(value.Type), LanguageID(value.Language)
		}
	}

	for i := 0; i < len(fileName); i++ {
		if _, ok := nameAndSuffixMap[fileName[i:]]; ok {
			return FileType(nameAndSuffixMap[fileName[i:]].Type), LanguageID(nameAndSuffixMap[fileName[i:]].Language)
		}
	}

	return FileTypeUnknown, LanguageUnknown
}

func findPossibleKeyConflicts(key string) error {
	if _, exist := nameAndSuffixMap[key]; exist {
		return errors.New(key + " already exists in nameAndSuffixMap")
	}
	if _, exist := prefixMap[key]; exist {
		return errors.New(key + " already exists in prefixMap")
	}
	return nil
}

func findPossiblePrefixConflicts() error {
	for key := range prefixMap {
		for key2 := range prefixMap {
			if (strings.HasPrefix(key, key2) || strings.HasPrefix(key2, key)) && key != key2 {
				return errors.New(key + " and " + key2 + " has prefix in common")
			}
		}
	}
	return nil
}
