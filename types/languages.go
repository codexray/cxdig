package types

import (
	"path"
)

// LanguageID identifies a programming language
type LanguageID string

func (id *LanguageID) String() string {
	return string(*id)
}

const (
	// LanguageUnknown is used to identify unknown languages
	LanguageUnknown LanguageID = ""
)

// based on https://gitlab.com/godeps/go-linguist
var knownSourceExtensions = map[string]LanguageID{
	// Assembly
	".asm": "Assembly",
	".s":   "Assembly",
	// C
	".c": "C",
	// C++
	".cpp": "C++",
	".cc":  "C++",
	".cxx": "C++",
	".c++": "C++",
	".h":   "C++",
	".hh":  "C++",
	".hpp": "C++",
	".hxx": "C++",
	".h++": "C++",
	".inc": "C++",
	".inl": "C++",
	".ipp": "C++",
	".tcc": "C++",
	".tpp": "C++",
	".txx": "C++",
	".moc": "C++", // Qt
	// C#
	".cs":  "C#",
	".csx": "C#",
	// D
	".d":  "D",
	".dd": "D",
	".di": "D",
	// FORTRAN
	".f":   "Fortran",
	".f03": "Fortran",
	".f08": "Fortran",
	".f77": "Fortran",
	".f90": "Fortran",
	".f95": "Fortran",
	".for": "Fortran",
	".fpp": "Fortran",
	// Go
	".go": "Go",
	// Java
	".java": "Java",
	// Pascal / Delphi
	".dfm":    "Pascal",
	".dpr":    "Pascal",
	".lpr":    "Pascal",
	".pas":    "Pascal",
	".pascal": "Pascal",
	// Perl
	".al":   "Perl",
	".perl": "Perl",
	".ph":   "Perl",
	".pl":   "Perl",
	".plx":  "Perl",
	".pm":   "Perl",
	".psgi": "Perl",
	// PHP
	".aw":    "PHP",
	".ctp":   "PHP",
	".php":   "PHP",
	".php3":  "PHP",
	".php4":  "PHP",
	".php5":  "PHP",
	".phps":  "PHP",
	".phpt":  "PHP",
	".phtml": "PHP",
	// Python
	".bzl":  "Python",
	".gyp":  "Python",
	".gypi": "Python",
	".lmi":  "Python",
	".py":   "Python",
	".py3":  "Python",
	".pyde": "Python",
	".pyp":  "Python",
	".pyt":  "Python",
	".pyw":  "Python",
	".pyx":  "Cython",
	".xpy":  "Python",
	// Groovy
	".groovy": "Groovy",
	".grt":    "Groovy",
	".gsp":    "Groovy",
	".gtpl":   "Groovy",
	".gvy":    "Groovy",
	// QML
	".qml": "QML",
	// R
	".r":   "R",
	".rd":  "R",
	".rsx": "R",
	// Ruby
	".eye":     "Ruby",
	".gemspec": "Ruby",
	".god":     "Ruby",
	".irbrc":   "Ruby",
	".rabl":    "Ruby",
	".rake":    "Ruby",
	".rb":      "Ruby",
	".rbuild":  "Ruby",
	".rbw":     "Ruby",
	".rbx":     "Ruby",
	".ru":      "Ruby",
	".ruby":    "Ruby",
	// Rust
	".rs":    "Rust",
	".rs.in": "Rust",
	// SQL
	".cql":   "SQL",
	".ddl":   "SQL",
	".mysql": "SQL",
	".prc":   "SQL",
	".sql":   "SQL",
	".tab":   "SQL",
	".udf":   "SQL",
	".viw":   "SQL",
	// VB
	".vb":  "Visual Basic",
	".vba": "Visual Basic",
	".vbs": "Visual Basic",
}

var knownGeneratorExtensions = map[string]LanguageID{
	// SWIG
	".i": "SWIG",
	// IDL
	".idl": "IDL",
	// protobuf / gRPC
	".proto": "Protocol Buffer",
	// Thrift
	".thrift": "Thrift",
	// FlatBuffers
	".fbs": "FlatBuffers",
	// Cap’n Proto
	".capnp": "Cap'n Proto",
	// Lex-Yacc
	".l":    "Lex",
	".lex":  "Lex",
	".ll":   "Lex",
	".y":    "Yacc",
	".yacc": "Yacc",
	".yxx":  "Yacc", // To be validated
	// GNU M4
	".m4": "M4",
}

func identifySourceFiles(fileName string) LanguageID {
	ext := path.Ext(fileName)
	if lang, ok := knownSourceExtensions[ext]; ok {
		return lang
	}

	return LanguageUnknown
}

func identifyGeneratorFiles(fileName string) LanguageID {
	ext := path.Ext(fileName)
	if lang, ok := knownGeneratorExtensions[ext]; ok {
		return lang
	}

	return LanguageUnknown
}
