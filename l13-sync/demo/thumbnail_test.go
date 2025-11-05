package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_ThumbnailGenerator1(t *testing.T) {
	t.Parallel()

	const image = "foo.png"

	for i := 0; i < 5; i++ {
		go generateThumbnail(image, i+1)
	}

	fmt.Println("Waiting for thumbnails to be generated")
}

// example of using waitgroup
func Test_WaitGroup(t *testing.T) {
	t.Parallel()

	var completed bool
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Millisecond * 10)

		fmt.Println("done with waitgroup")

		completed = true
		// Done to be called for each go routine
		wg.Done()
	}(&wg)

	fmt.Println("waiting for waitgroup to unblock")

	wg.Wait()

	fmt.Println("waitgroup is unblocked")

	if !completed {
		t.Fatal("waigroup is not completed")
	}
}

// Test_ThumbnailGenerator - updated with right waitgroup
func Test_ThumbnailGenerator2(t *testing.T) {
	t.Parallel()

	const image = "foo.png"
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			// done is called after completing the generateThumbnail function
			defer wg.Done()
			generateThumbnail(image, i+1)
		}(i)
	}

	fmt.Println("Waiting for thumbnails to be generated")
	wg.Wait()
	fmt.Println("Finished generate all thumbnails")
}

