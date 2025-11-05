package demo

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func Walk() ([]string, error) {
	var entries []string

	err := filepath.WalkDir("data", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			entries = append(entries, path)
			return nil
		}

		name := d.Name()

		switch {
		case name == ".":
			return nil
		case name == "..":
			return nil
		case name == "testdata":
			return fs.SkipDir
		case strings.HasPrefix(name, "."):
			return fs.SkipDir
		case strings.HasPrefix(name, "_"):
			return fs.SkipDir
		}

		return nil
	})

	return entries, err
}

