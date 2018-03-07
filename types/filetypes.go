package types

// FileType identifies the type of a file in the repository
type FileType string

func (id *LanguageID) String() string {
	return string(*id)
}

// LanguageID identifies the language used for a source file
type LanguageID string

func (id *FileType) String() string {
	return string(*id)
}

// FileTypeInfo is the struct used to extract FileType Information from a JSON
type FileTypeInfo struct {
	Language   LanguageID `json:"language"`
	Type       FileType   `json:"type"`
	FileName   []string   `json:"fileNames"`
	FilePrefix []string   `json:"filePrefixes"`
	FileSuffix []string   `json:"fileSuffixes"`
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
