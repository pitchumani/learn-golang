package main

// check defer function

import "fmt"
import "os"
import "time"

func createFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// close the file after this function returns
	defer f.Close()

	_, err = f.WriteString("Hello World\n")
	if err != nil {
		return err
	}

	return nil
}


func checkDeferTimings() {
	fmt.Println("checkDeferTimings")
	now := time.Now()
	// this is called after the function returns, but the argument is
	// evaluated now, see the duration (that doesn't include the sleep ms)
	defer fmt.Printf("duration: %s\n", time.Since(now))
	fmt.Println("Sleeping for 50ms...")
	time.Sleep(50 * time.Millisecond)
}

// use anonymous function to evaluate time correctly
func checkDeferTimings1() {
	fmt.Println("checkDeferTimings1")
	now := time.Now()
	
	defer func(now time.Time) {
		fmt.Printf("duration: %s\n", time.Since(now))
	}(now)
	fmt.Println("Sleeping for 50ms...")
	time.Sleep(50 * time.Millisecond)
}

func main() {
	fmt.Println("Deferring function calls")

	// prints goodbyte after main return
	defer fmt.Println("goodbye")

	fmt.Println("Hello")

	// demonstrate defer function
	createFile("test-file.txt")

	// executed in reverse
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	// panic is executed at the end
	defer fmt.Println("four")
	defer panic("five")
	defer fmt.Println("six")

	checkDeferTimings()
	checkDeferTimings1()
}
