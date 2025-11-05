package demo

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTestData(t testing.TB) {
	t.Helper()

	if err := os.RemoveAll("data"); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir("data", 0755); err != nil {
		t.Fatal(err)
	}
	list := []string{
		"data/.hidden/d.txt",
		"data/a.txt",
		"data/b.txt",
		"data/e/f/_ignore/i.txt",
		"data/e/f/g.txt",
		"data/e/f/h.txt",
		"data/e/j.txt",
		"data/testdata/c.txt",
	}
	for _, path := range list {
		if ext := filepath.Ext(path); len(ext) > 0 {
			path = filepath.Dir(path)
		}
	    // use MkdirAll to create all subdirectories 
	    //if err := os.Mkdir(path, 0755); err != nil {
	    if err := os.MkdirAll(path, 0755); err != nil {
			// ignore if the directory already exists
			if !errors.Is(err, fs.ErrExist) {
				t.Fatal(err)
			}
		}

		fmt.Println("creating", path)
		f, err := os.Create(path)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Fprintf(f, "%s", strings.ToUpper(path))

		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
		
		
func Test_Walk(t *testing.T) {
	t.Parallel()
	createTestData(t)
	exp := []string{
		filepath.Join("data", "a.txt"),
		filepath.Join("data", "b.txt"),
		filepath.Join("data", "e", "f", "g.txt"),
		filepath.Join("data", "e", "f", "h.txt"),
		filepath.Join("data", "e", "j.txt"),
	}
	act, err := Walk()
	if err != nil {
		t.Fatal(err)
	}

	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if as != es {
		t.Fatalf("expected %s, got %s", es, as)
	}
}
