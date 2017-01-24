package main

import (
	"fmt"
	"strings"
)

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
		r += strings.Replace(s.Content, "\n\n", "\n.PP\n", -1) + "\n"

	}
	return r
}
