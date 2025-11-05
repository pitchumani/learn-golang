package main

import "fmt"

func bubbleSort(arr *[]int) int {
	n_iterations := 0

	// for each element, find max value and move it to the end
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr) - i - 1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
			n_iterations++
		}
	}
	return n_iterations
}

func main() {
	var arr = []int{45,67,23,4,78,1, 0}
	fmt.Printf("Array before sort:")
	for _, i := range arr {
		fmt.Printf(" %v", i)
	}
	fmt.Println("")
	n_iterations := bubbleSort(&arr)
	fmt.Printf("Array after sort:")
	for _, i := range arr {
		fmt.Printf(" %v", i)
	}
	fmt.Printf("\nBubble sort took %v iterations to sort %v elements\n",
		n_iterations, len(arr))
}
