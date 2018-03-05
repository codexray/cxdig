package types

import "testing"
import "github.com/stretchr/testify/assert"

func TestIdentifySourceFiles(t *testing.T) {
	assert.Equal(t, LanguageID("C++"), identifySourceFiles("abc/file.cpp"))
	assert.Equal(t, LanguageID("C++"), identifySourceFiles("abc/file.inl"))
	assert.Equal(t, LanguageID("C++"), identifySourceFiles("abc/file.moc"))
	assert.Equal(t, LanguageID("C#"), identifySourceFiles("abc/file.cs"))
	assert.Equal(t, LanguageID("D"), identifySourceFiles("abc/file.d"))
	assert.Equal(t, LanguageID("Fortran"), identifySourceFiles("abc/file.f"))
	assert.Equal(t, LanguageID("Java"), identifySourceFiles("abc/file.java"))
	assert.Equal(t, LanguageUnknown, identifySourceFiles("abc/file.cmake"))
}

func TestIdentifyGeneratorFile(t *testing.T) {
	assert.Equal(t, LanguageID("Lex"), identifyGeneratorFiles("abc/file.l"))
	assert.Equal(t, LanguageID("Thrift"), identifyGeneratorFiles("abc/file.thrift"))
	assert.Equal(t, LanguageID("IDL"), identifyGeneratorFiles("abc/file.idl"))
	assert.Equal(t, LanguageID("Lex"), identifyGeneratorFiles("abc/file.ll"))
	assert.Equal(t, LanguageID("M4"), identifyGeneratorFiles("abc/file.m4"))
	assert.Equal(t, LanguageUnknown, identifyGeneratorFiles("abc/file.unknown"))

}
