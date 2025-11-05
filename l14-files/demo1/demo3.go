package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Mode\t\tSize\tName\n")
	err := filepath.WalkDir("../../l13-sync", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// error - may be the path does not exists
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		fmt.Printf("%10s%5d\t\t%s\n", info.Mode(), info.Size(), path)

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
