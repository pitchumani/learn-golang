package demo

import (
	"golang.org/x/sync/errgroup"
	"fmt"
	//"sync"
	"testing"
	//"time"
)

// use error group instead of wait group
func Test_ThumbnailGenerator2(t *testing.T) {
	t.Parallel()

	const image = "foo.png"
	var wg errgroup.Group

	for i := 0; i < 5; i++ {
		i := i
		wg.Go(func() error {
			return generateThumbnail(image, i+1)
		})
	}

	fmt.Println("Waiting for thumbnails to be generated")
	err := wg.Wait()

	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println("Finished generate all thumbnails")
}

