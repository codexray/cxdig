package gitlog

import "testing"

import "github.com/stretchr/testify/assert"

func getGitLogSample() []string {
	text := []string{}
	text = append(text, "commit 18f9a796892229fbd04f762f01d6b01d159bd3e7\n")
	text = append(text, "Author: John Doe <john.doe@example.com>\n")
	text = append(text, "Date:   Sat, 5 Feb 2017 03:50:46 +0100\n")
	text = append(text, "\n")
	text = append(text, "    Decode json text from parsing module\n")
	text = append(text, "\n")
	text = append(text, "M       go-scripts/projects-parser/ParsingTypes.go\n")
	text = append(text, "M       go-scripts/projects-parser/cppparser.go\n")
	text = append(text, "M       go-scripts/projects-parser/main.go\n")
	text = append(text, "M       src/CppParser/ParserApi/CppParserTypes.cpp\n")
	text = append(text, "M       src/CppParser/app/main.cpp\n")
	return text
}

func TestTimeDecoding(t *testing.T) {
	decoder := GitLogDecoder{}
	decoder.DecodeDateTime("Date: Mon, 5 Dec 2016 11:23:39 +0100")
}

func TestAuthorDecoding(t *testing.T) {
	decoder := GitLogDecoder{}
	author := decoder.DecodeAuthorInfo("John Doe<john.doe@example.com> ")
	assert.Equal(t, "John Doe", author.Name)
	assert.Equal(t, "john.doe@example.com", author.Email)

	decoder.Clear()
	decoder.DecodeAuthorInfo("John Doe <> ")
	assert.Error(t, decoder.Err())

	decoder.Clear()
	decoder.DecodeAuthorInfo("John Doe ")
	assert.Error(t, decoder.Err())

	author = decoder.DecodeAuthorInfo("< john.doe@example.com > ")
	assert.Equal(t, "", author.Name)
	assert.Equal(t, "john.doe@example.com", author.Email)
}
