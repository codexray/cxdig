package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifyBuildSystemFiles(t *testing.T) {
	ident := identifyBuildSystemFiles("cmakelists.txt")
	assert.Equal(t, LanguageID("CMake"), ident)

	ident = identifyBuildSystemFiles("makefile")
	assert.Equal(t, LanguageID("Makefile"), ident)

	ident = identifyBuildSystemFiles("test.pro")
	assert.Equal(t, LanguageID("QMake"), ident)

	ident = identifyBuildSystemFiles("test.sln")
	assert.Equal(t, LanguageID("Visual Studio"), ident)

	ident = identifyBuildSystemFiles("test.xcconfig")
	assert.Equal(t, LanguageID("XCode"), ident)

	ident = identifyBuildSystemFiles("test.am")
	assert.Equal(t, LanguageID("Automake"), ident)

	ident = identifyBuildSystemFiles("test.qbs")
	assert.Equal(t, LanguageID("QBS"), ident)

	ident = identifyBuildSystemFiles("test.ninja")
	assert.Equal(t, LanguageID("Ninja"), ident)

	ident = identifyBuildSystemFiles("test.vcpkg")
	assert.Equal(t, LanguageID("Vcpkg"), ident)

	ident = identifyBuildSystemFiles("test.jam")
	assert.Equal(t, LanguageID("Boost.Jam"), ident)

	ident = identifyBuildSystemFiles("test.gradle")
	assert.Equal(t, LanguageID("Gradle"), ident)

	ident = identifyBuildSystemFiles("test.Tup")
	assert.Equal(t, LanguageID("Tup"), ident)

	ident = identifyBuildSystemFiles("test.bzl")
	assert.Equal(t, LanguageID("Bazel"), ident)

	ident = identifyBuildSystemFiles("test.gyp")
	assert.Equal(t, LanguageID("GYP"), ident)

	ident = identifyBuildSystemFiles("test.iss")
	assert.Equal(t, LanguageID("Inno Setup"), ident)

	ident = identifyBuildSystemFiles("test.wixproj")
	assert.Equal(t, LanguageID("WiX"), ident)

	ident = identifyBuildSystemFiles("test.nsh")
	assert.Equal(t, LanguageID("NSIS"), ident)

	ident = identifyBuildSystemFiles("conanfile.txt")
	assert.Equal(t, LanguageID("Conan"), ident)

	ident = identifyBuildSystemFiles("sconstruct")
	assert.Equal(t, LanguageID("SCons"), ident)

	ident = identifyBuildSystemFiles("premake4.lua")
	assert.Equal(t, LanguageID("Premake"), ident)

	ident = identifyBuildSystemFiles("gulp.js")
	assert.Equal(t, LanguageID("Gulp"), ident)

	ident = identifyBuildSystemFiles("zeusfile.yml")
	assert.Equal(t, LanguageID("ZEUS"), ident)

	ident = identifyBuildSystemFiles("bam.lua")
	assert.Equal(t, LanguageID("bam"), ident)

	ident = identifyBuildSystemFiles("meson.build")
	assert.Equal(t, LanguageID("Meson"), ident)

	ident = identifyBuildSystemFiles("huntergate.cmake")
	assert.Equal(t, LanguageID("Hunter"), ident)

	ident = identifyBuildSystemFiles("requirements.txt")
	assert.Equal(t, LanguageID("cget"), ident)

	ident = identifyBuildSystemFiles("meta.yaml")
	assert.Equal(t, LanguageID("Conda"), ident)

	ident = identifyBuildSystemFiles("build.hs")
	assert.Equal(t, LanguageID("Shake"), ident)

	ident = identifyBuildSystemFiles("gemfile")
	assert.Equal(t, LanguageID("Gemfile"), ident)

	ident = identifyBuildSystemFiles("package.json")
	assert.Equal(t, LanguageID("npm"), ident)

	ident = identifyBuildSystemFiles("bower.json")
	assert.Equal(t, LanguageID("Bower"), ident)

	ident = identifyBuildSystemFiles("pom.xml")
	assert.Equal(t, LanguageID("Maven"), ident)

	ident = identifyBuildSystemFiles("TestOgUnknownLanguage")
	assert.Equal(t, LanguageUnknown, ident)
}

func TestIdentifyEnvConfigFiles(t *testing.T) {
	ident := identifyEnvConfigFiles(".eslintrc.test")
	assert.Equal(t, LanguageID("ESLint"), ident)

	ident = identifyEnvConfigFiles(".travis.yml")
	assert.Equal(t, LanguageID("Travis"), ident)

	ident = identifyEnvConfigFiles("appveyor.yml")
	assert.Equal(t, LanguageID("AppVeyor"), ident)

	ident = identifyEnvConfigFiles(".gitlab-ci.yml")
	assert.Equal(t, LanguageID("Gitlab"), ident)

	ident = identifyEnvConfigFiles(".clang-format")
	assert.Equal(t, LanguageID("ClangFormat"), ident)

	ident = identifyEnvConfigFiles(".clang_complete")
	assert.Equal(t, LanguageID("clang_complete"), ident)

	ident = identifyEnvConfigFiles(".editorconfig")
	assert.Equal(t, LanguageID("EditorConfig"), ident)

	ident = identifyEnvConfigFiles(".gdbinit")
	assert.Equal(t, LanguageID("gdbinit"), ident)

	ident = identifyEnvConfigFiles(".yardopts")
	assert.Equal(t, LanguageID("YARD"), ident)

	ident = identifyEnvConfigFiles(".istanbul.yml")
	assert.Equal(t, LanguageID("Istanbul"), ident)

	ident = identifyEnvConfigFiles(".pylintrc")
	assert.Equal(t, LanguageID("Pylint"), ident)

	ident = identifyEnvConfigFiles(".flake8")
	assert.Equal(t, LanguageID("Flake8"), ident)

	ident = identifyEnvConfigFiles(".dir-locals.el")
	assert.Equal(t, LanguageID("Emacs.dir-locals"), ident)

	ident = identifyEnvConfigFiles("doxygen.config")
	assert.Equal(t, LanguageID("Doxygen"), ident)

	ident = identifyEnvConfigFiles(".mention-bot")
	assert.Equal(t, LanguageID("mention-bot"), ident)

	ident = identifyEnvConfigFiles("unknown type")
	assert.Equal(t, LanguageUnknown, ident)
}

func TestIdentifyLicenseFiles(t *testing.T) {
	ident := identifyLicenseFiles("APACHE-2.0.txt")
	assert.Equal(t, LanguageID("Apache-2.0"), ident)

	ident = identifyLicenseFiles("GNU-AGPL-3.0.txt")
	assert.Equal(t, LanguageID("AGPL-3.0"), ident)

	ident = identifyLicenseFiles("unknown type")
	assert.Equal(t, LanguageUnknown, ident)
}

func TestIdentifyFileTypeAndLanguage(t *testing.T) {
	fileType, languageID := IdentifyFileTypeAndLanguage("test.c")
	assert.Equal(t, FileTypeSource, fileType)
	assert.Equal(t, LanguageID("C"), languageID)

	fileType, languageID = IdentifyFileTypeAndLanguage("test.lex")
	assert.Equal(t, FileTypeGenerator, fileType)
	assert.Equal(t, LanguageID("Lex"), languageID)

	fileType, languageID = IdentifyFileTypeAndLanguage("makefile")
	assert.Equal(t, FileTypeBuildSystem, fileType)
	assert.Equal(t, LanguageID("Makefile"), languageID)

	fileType, languageID = IdentifyFileTypeAndLanguage(".gitlab-ci.yml")
	assert.Equal(t, FileTypeEnvConfig, fileType)
	assert.Equal(t, LanguageID("Gitlab"), languageID)

	fileType, languageID = IdentifyFileTypeAndLanguage("test.c.in")
	assert.Equal(t, FileTypeGenerator, fileType)
	assert.Equal(t, LanguageID("C"), languageID)

	fileType, languageID = IdentifyFileTypeAndLanguage("unknownfile")
	assert.Equal(t, FileTypeUnknown, fileType)
	assert.Equal(t, LanguageUnknown, languageID)
}
