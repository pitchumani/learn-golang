package main

import "fmt"

func selectionSort(arr *[]int) int {
	n_iterations := 0
	var arr_sz = len(*arr)
	
	for i := 0; i < arr_sz; i++ {
		var min_idx int = i
		for j := i+1; j < arr_sz; j++ {
			if (*arr)[j] < (*arr)[min_idx] {
				min_idx = j
			}
			n_iterations++
		}
		// swap the min value to correct place (min_idx)
		tmp_val := (*arr)[i]
		(*arr)[i] = (*arr)[min_idx]
		(*arr)[min_idx] = tmp_val
	}
	return n_iterations
}

func main() {
	var arr = []int{45,67,23,4,78,1,0}
	fmt.Printf("Array before sort:")
	for _, i := range arr {
		fmt.Printf(" %v", i)
	}
	fmt.Println("")
	n_iterations := selectionSort(&arr)
	fmt.Printf("Array after sort:")
	for _, i := range arr {
		fmt.Printf(" %v", i)
	}
	fmt.Printf("\nSelection sort took %v iterations to sort %v elements\n",
		n_iterations, len(arr))
}
