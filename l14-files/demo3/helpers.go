package demo

import (
	"strings"
	"testing"
	"testing/fstest"
)

func createTestFS(t testing.TB) fstest.MapFS {
	t.Helper()

	cab := fstest.MapFS{}

	files := []string{
		".hidden/d.txt",
		"a.txt",
		"b.txt",
		"e/f/_ignore/i.txt",
		"e/f/g.txt",
		"e/f/h.txt",
		"e/j.txt",
		"testdata/c.txt",
	}

	for _, path := range files {
		cab[path] = &fstest.MapFile{
			Data: []byte(strings.ToUpper(path)),
		}
	}
	return cab
}
		
