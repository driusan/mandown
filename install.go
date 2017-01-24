package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func (mp ManPage) InstallTo(dir string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s.%d.gz", dir, mp.Name, mp.Section))
	if err != nil {
		return err
	}
	defer f.Close()
	w := gzip.NewWriter(f)
	defer w.Close()
	fmt.Fprintf(w, "%s", mp.TroffString())
	return nil
}
