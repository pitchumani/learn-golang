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

	fmt.Println("Mode\t\tSize\tName")
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\t%d\t\t%s\n", info.Mode(), info.Size(), info.Name())
	}
}
