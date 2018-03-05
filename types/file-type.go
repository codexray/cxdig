package types

import (
	"path"
	"strings"
)

// FileType identifies the type of a file in the repository
type FileType string

func (id *FileType) String() string {
	return string(*id)
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
)

func identifyBuildSystemFiles(fileName string) LanguageID {
	fileName = strings.ToLower(fileName)

	switch fileName {
	case "conanfile.txt", "conanfile.py", "conanenv.txt":
		return "Conan"
	case "sconstruct":
		return "SCons" // SCons / Cuppa
	case "premake4.lua", "premake5.lua":
		return "Premake"
	case "gulp.js":
		return "Gulp"
	case "zeusfile.yml": // https://github.com/dreadl0ck/zeus
		// TODO: handle zeus/scripts path
		return "ZEUS"
	case "bam.lua": // http://matricks.github.io/bam/
		return "bam"
	case "meson.build":
		return "Meson"
	case "huntergate.cmake", "hunter.cmake":
		return "Hunter"
	case "requirements.txt":
		return "cget"
	case "meta.yaml":
		return "Conda"
	case "build.hs":
		return "Shake"
	case "gemfile":
		return "Gemfile"
	case "package.json":
		return "npm"
	case "webpack.config.js":
		return "webpack"
	case "bower.json":
		return "Bower"
	case "pom.xml":
		return "Maven"
	}

	// CMake
	if fileName == "cmakelists.txt" ||
		strings.HasSuffix(fileName, ".cmake") ||
		strings.HasSuffix(fileName, ".cmake.in") {
		return "CMake"
	}

	// Make
	if fileName == "makefile" ||
		strings.HasSuffix(fileName, ".make") ||
		strings.HasSuffix(fileName, ".mkfile") ||
		// NMake
		strings.HasSuffix(fileName, ".mak") ||
		// dmake
		strings.HasSuffix(fileName, ".mk") {
		return "Makefile"
	}

	// QMake
	if strings.HasSuffix(fileName, ".pro") ||
		strings.HasSuffix(fileName, ".pri") {
		return "QMake"
	}

	// VC++
	if strings.HasSuffix(fileName, ".sln") ||
		strings.HasSuffix(fileName, ".vcxproj") ||
		strings.HasSuffix(fileName, ".vcproj") ||
		strings.HasSuffix(fileName, ".props") {
		return "Visual Studio"
	}

	// XCode
	if strings.HasSuffix(fileName, ".xcconfig") ||
		strings.HasSuffix(fileName, ".pbxproj") ||
		strings.HasSuffix(fileName, ".xcworkspacedata") {
		return "XCode"
	}

	// Automake
	if strings.HasSuffix(fileName, ".am") {
		return "Automake"
	}

	// Qbs
	if strings.HasSuffix(fileName, ".qbs") {
		return "QBS"
	}

	// Ninja
	if strings.HasSuffix(fileName, ".ninja") {
		return "Ninja"
	}
	// Vcpkg
	if strings.Contains(fileName, ".vcpkg") {
		return "Vcpkg"
	}

	// Boost.Jam
	if strings.HasSuffix(fileName, ".jam") {
		return "Boost.Jam"
	}

	// Gradle
	if strings.HasSuffix(fileName, ".gradle") {
		return "Gradle"
	}

	// Tup
	if strings.HasSuffix(fileName, ".tup") {
		return "Tup"
	}

	// Bazel
	if strings.HasSuffix(fileName, ".bzl") {
		return "Bazel"
	}

	// GYP
	if strings.HasSuffix(fileName, ".gyp") || strings.HasSuffix(fileName, ".gypi") {
		return "GYP"
	}

	// ----------- INSTALLERS -----------
	// Inno Setup
	if strings.HasSuffix(fileName, ".iss") {
		return "Inno Setup"
	}
	// WiX
	if strings.HasSuffix(fileName, ".wixproj") || strings.HasSuffix(fileName, ".wxs") {
		return "WiX"
	}
	// NSIS
	if strings.HasSuffix(fileName, ".nsh") || strings.HasSuffix(fileName, ".nsi") {
		return "NSIS"
	}
	// ---------------------------------

	// GENie
	//strings.HasSuffix(fileName, ".lua") ||

	return LanguageUnknown
}

func identifyEnvConfigFiles(fileName string) LanguageID {
	fileName = strings.ToLower(fileName)

	if strings.HasPrefix(fileName, ".eslintrc.") {
		return "ESLint"
	}

	// TODO: handle ".github/*" files (issue_template.md, pull_request_template.md)
	// TODO: try to handle "*.pc" files as pkg-config files

	switch fileName {
	case ".travis.yml":
		return "Travis"
	case "appveyor.yml":
		return "AppVeyor"
	case ".gitlab-ci.yml":
		return "Gitlab"
	case "circle.yml":
		// TODO: handle ".circleci/config.yml" path
		return "CircleCI"
	case ".clang-format":
		return "ClangFormat"
	// http://editorconfig.org/
	case ".clang_complete":
		return "clang_complete" // plugin for Vim
	case ".editorconfig":
		return "EditorConfig"
	case ".gdbinit":
		return "gdbinit"
	case ".yardopts":
		return "YARD"
	case ".istanbul.yml":
		return "Istanbul" // JS code coverage tool
	case "codecov.yml":
	case ".codecov.yml":
		return "Codecov.io" // code coverage
	case ".pylintrc":
		return "Pylint"
	case ".flake8":
		return "Flake8" // Python pep8 linter
	case ".dir-locals.el":
		return "Emacs.dir-locals"
	case "doxygen.config":
		return "Doxygen"
	case ".mention-bot":
		return "mention-bot" // Automatically mention potential reviewers on pull requests
	}

	return LanguageUnknown
}

func identifyLicenseFiles(fileName string) LanguageID {
	//fileName = strings.ToLower(fileName)

	switch fileName {
	case "APACHE-2.0.txt":
		return "Apache-2.0"
	case "GNU-AGPL-3.0.txt":
		return "AGPL-3.0"
	}

	return LanguageUnknown
}

// IdentifyFileTypeAndLanguage tries to identify the type and eventual language from a given file name
func IdentifyFileTypeAndLanguage(fileName string) (FileType, LanguageID) {
	fileName = strings.ToLower(fileName)
	ext := path.Ext(fileName)

	if ext == ".in" {
		// for files generated by configure ("MD5.h.in", FStream.hxx.in", ...)
		// use the second level file extension
		_, lang := IdentifyFileTypeAndLanguage(fileName[:len(fileName)-3])
		return FileTypeGenerator, lang
	}

	if lang := identifySourceFiles(fileName); lang != LanguageUnknown {
		return FileTypeSource, lang
	}

	if lang := identifyGeneratorFiles(fileName); lang != LanguageUnknown {
		return FileTypeGenerator, lang
	}

	if lang := identifyBuildSystemFiles(fileName); lang != LanguageUnknown {
		return FileTypeBuildSystem, lang
	}

	if lang := identifyEnvConfigFiles(fileName); lang != LanguageUnknown {
		return FileTypeEnvConfig, lang
	}

	return FileTypeUnknown, LanguageUnknown
}
