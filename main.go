package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ManPage struct {
	Name, Title string
	Section     uint8
	Sections    []Subsection
}

type Subsection struct {
	Name    string
	Content string
}

var fnameRe *regexp.Regexp = regexp.MustCompile(`(.+)\.([1-9])\.md`)

func guessSection(f *os.File) uint8 {
	if matches := fnameRe.FindStringSubmatch(f.Name()); matches != nil {
		n, err := strconv.Atoi(matches[2])
		if err == nil {
			return uint8(n)
		}

	}
	return 1
}
func (mp ManPage) String() string {
	r := "NAME\n\t" + mp.Name

	if mp.Title != "" {
		r += " - " + mp.Title
	}
	for _, s := range mp.Sections {
		indented := strings.Replace(s.Content, "\n", "\n\t", -1)

		// This isn't really necessary, but just for my sanity.
		//
		// Get rid of any blank lines we just created which only contain
		// a tab.
		indented = strings.Replace(indented, "\t\n", "\n", -1)
		r += "\n\n" + strings.ToUpper(s.Name) + "\n\t" + indented
	}
	return r
}

func main() {
	troff := flag.Bool("t", false, "Output in troff format")
	section := flag.Int("section", 0, "Specify section of manual")
	install := flag.String("install", "", "Gzip troff output and copy to dir")

	flag.Parse()
	files := flag.Args()
	if len(files) < 1 || (*section > 9) {
		flag.Usage()
		os.Exit(2)
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Println(err)
			continue
		}
		md, err := ReadFile(f)
		if err != nil {
			log.Println(err)
		}
		f.Close()

		if *section > 0 {
			md.Section = uint8(*section)
		} else {
			md.Section = guessSection(f)
		}

		if *install != "" {
			err := md.InstallTo(*install)
			if err != nil {
				log.Println(err)
			}
		} else if *troff {
			fmt.Println(md.TroffString())
		} else {
			fmt.Println(md.String())
		}
	}
}
