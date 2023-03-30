// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package markdown

import (
	"regexp"
)

var singleLineCodeBlockRegexp = regexp.MustCompile(`{{{([^\n]+?)}}}`)
var multiLineCodeBlockRegexp = regexp.MustCompile(`(?m)^{{{(?s)(.+?)^}}}`)
var nonCodeBlockRegexp = regexp.MustCompile(`(?m)(?:}}}$|\A)(?s)(.+?)(?:^{{{|\z)`)
var commitTicketRefRegexp = regexp.MustCompile(`(?m)\x60\x60\x60\n?#!Commit.*\n`)
var langRegexp = regexp.MustCompile(`(?m)\x60\x60\x60\n?#!(c|c\+\+|ps1|php|py|sh|cpp|pl)\n`)
var langMap = map[string]string{"c++": "cpp"}

func (converter *DefaultConverter) convertCodeBlocks(in string) string {
	// convert single line {{{...}}} to `...`
	out := singleLineCodeBlockRegexp.ReplaceAllString(in, "`$1`")

	// convert multi-line {{{...}}} to ```-delimited lines
	// - we leave in place any Trac '#!...' sequences following the opening '{{{'
	//   since we have no easy way of dealing with these and they are best left in place
	//   as a reminder to review them in the Gitea world
	out = multiLineCodeBlockRegexp.ReplaceAllStringFunc(out, func(match string) string {
		text := multiLineCodeBlockRegexp.ReplaceAllString(match, `$1`)
		return "```" + text + "```"
	})

	// but remove #!CommitTicketReference repository="" revision=""
	out = commitTicketRefRegexp.ReplaceAllString(out, "```\n")

	// and fixup some common languages
	out = langRegexp.ReplaceAllStringFunc(out, func(match string) string {
		if matches := langRegexp.FindStringSubmatch(match); matches != nil {
			if lang, found := langMap[matches[1]]; found {
				return langRegexp.ReplaceAllString(match, "```"+lang+"\n")
			} else {
				return langRegexp.ReplaceAllString(match, "```$1\n")
			}
		}

		return match
	})

	return out
}

func (converter *DefaultConverter) convertNonCodeBlocks(in string, convertFn func(string) string) string {
	return nonCodeBlockRegexp.ReplaceAllStringFunc(in, convertFn)
}
