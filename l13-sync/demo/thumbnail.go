package demo

import (
	"fmt"
	"time"
)

/*
func generateThumbnail(image string, size int) {
	thumb := fmt.Sprintf("%s@%dx.png", image, size)
	fmt.Println("Generating thumbnail:", thumb)

	time.Sleep(time.Millisecond * time.Duration(size))

	fmt.Println("Finished generating thumbnail:", thumb)
}
*/
func generateThumbnail(image string, size int) error {
	if size % 5 == 0 {
		return fmt.Errorf("%d is divisible by 5", size)
	}
	
	thumb := fmt.Sprintf("%s@%dx.png", image, size)

	fmt.Println("Generating thumbnail:", thumb)

	time.Sleep(time.Millisecond * time.Duration(size))

	fmt.Println("Finished generating thumbnail:", thumb)

	return nil
}

