package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var titleRe *regexp.Regexp = regexp.MustCompile("^# (.+) - (.+)")
var subsectionRe *regexp.Regexp = regexp.MustCompile("^## (.+)")

func ReadFile(r io.Reader) (ManPage, error) {
	scanner := bufio.NewReader(r)

	var mp ManPage
	var line string
	var err error
	var curSection Subsection
	for {
		line, err = scanner.ReadString('\n')
		switch err {
		case io.EOF:
			if curSection.Name != "" {
				curSection.Content = strings.TrimSpace(curSection.Content)
				mp.Sections = append(mp.Sections, curSection)
			}
			return mp, nil
		case nil:
			// Nothing special
		default:
			if curSection.Name != "" {
				curSection.Content = strings.TrimSpace(curSection.Content)

				mp.Sections = append(mp.Sections, curSection)
			}
			return mp, err
		}
		line = strings.TrimSuffix(line, "\n")
		if strings.TrimSpace(line) == "" {
			// If there's multiple blank lines, collapse them into
			// one.
			// Blank lines at the beginning or end are collapsed
			// before appending the subsection with strings.TrimSpace
			c := curSection.Content
			if len(c) > 2 && c[len(c)-2:] == "\n\n" {
				continue
			}
		}
		if matches := titleRe.FindStringSubmatch(line); matches != nil {
			mp.Name = matches[1]
			mp.Title = matches[2]
			curSection = Subsection{}
		} else if matches := subsectionRe.FindStringSubmatch(line); matches != nil {
			if curSection.Name != "" {
				curSection.Content = strings.TrimSpace(curSection.Content)
				mp.Sections = append(mp.Sections, curSection)
			}
			curSection = Subsection{Name: matches[1]}
		} else {
			if curSection.Name != "" {
				curSection.Content += line + "\n"
			}
		}
	}
}
