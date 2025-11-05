package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("../../l13-sync")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("-> ", file.Name())
			continue
		}
		fmt.Println(file.Name())
	}
}
