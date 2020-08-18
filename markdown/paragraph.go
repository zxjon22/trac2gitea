package markdown

import "regexp"

var pageBreakRegexp = regexp.MustCompile(`\[\[[Bb][Rr]\]\]`)

func (converter *DefaultConverter) convertParagraphs(in string) string {
	// convert trac page breaks to HTML <br>s
	// - the alternative of "  \n" to force a newline doesn't work in the likes of table cells
	return pageBreakRegexp.ReplaceAllString(in, "<br>")
}
