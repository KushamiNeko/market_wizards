package minify

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"regexp"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	multiStartingSpace    *regexp.Regexp
	multiEmptyLine        *regexp.Regexp
	htmlSingleLineComment *regexp.Regexp
	cssSingleLineComment  *regexp.Regexp
	jsSingleLineComment   *regexp.Regexp
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {

	multiStartingSpace = regexp.MustCompile(`(?m)^\s+`)

	multiEmptyLine = regexp.MustCompile(`(?m)\s*\n+`)

	htmlSingleLineComment = regexp.MustCompile(`(\s|\n)*<!--[^\r]*?-->`)

	cssSingleLineComment = regexp.MustCompile(`(?m)^(\s|\n)*/\*[^\r]*?\*/`)

	jsSingleLineComment = regexp.MustCompile(`(\s|\n)//.*\n`)

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Minify(content []byte) []byte {

	content = jsSingleLineComment.ReplaceAll(content, []byte(""))

	content = htmlSingleLineComment.ReplaceAll(content, []byte(""))

	content = cssSingleLineComment.ReplaceAll(content, []byte(""))

	content = multiEmptyLine.ReplaceAll(content, []byte("\n"))

	content = multiStartingSpace.ReplaceAll(content, []byte(""))

	return content
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
