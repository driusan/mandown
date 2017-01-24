package main

import (
	"fmt"
	"regexp"
	"strings"
)

var boldRe *regexp.Regexp = regexp.MustCompile(`\*(.+?)\*([ \t]*)`)
var culRe *regexp.Regexp = regexp.MustCompile(`__(.+?)__([ \t]*)`)
var ulRe *regexp.Regexp = regexp.MustCompile(`_(.+?)_([ \t]*)`)

func (mp ManPage) TroffString() string {
	r := fmt.Sprintf(".TH %s %d\n", strings.ToUpper(mp.Name), mp.Section)
	r += fmt.Sprintf(".SH NAME\n%s \\- %s\n", mp.Name, mp.Title)

	for _, s := range mp.Sections {
		// This is probably not the most efficient way to check if a
		// string contains a space, but it avoids importing another
		// package
		if len(strings.Fields(s.Name)) > 1 {
			r += fmt.Sprintf(".SH \"%s\"\n", strings.ToUpper(s.Name))
		} else {
			r += fmt.Sprintf(".SH %s\n", strings.ToUpper(s.Name))
		}

		// Do a couple transformations
		// Paragraphs
		content := strings.Replace(s.Content, "\n\n", "\n.PP\n", -1) + "\n"
		// *Bold*
		content = boldRe.ReplaceAllString(content, "\n.B $1\n")
		// __continuous underline__
		content = culRe.ReplaceAllString(content, "\n.cu\n$1\n")
		// _underline_
		content = ulRe.ReplaceAllString(content, "\n.ul\n$1\n")
		r += fmt.Sprintf(content)

	}
	return r
}
