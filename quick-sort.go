package main

import (
	"fmt"
)

func partition(arr *[]int, low int, high int) int {
	var pivot = (*arr)[high]
	i := low - 1
	for j := low; j <= high - 1; j++ {
		if (*arr)[j] < pivot {
			i++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = tmp;
		}
	}
	// place the actual pivot in the pivot index
	tmp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[high]
	(*arr)[high] = tmp
	return i + 1
}

func quickSort(arr *[]int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high)
	}
}

func main() {
	var arr = []int{12,3,56,7,0,34,8,98}
	fmt.Print("array before sort: ")
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	fmt.Println();
	quickSort(&arr, 0, len(arr) - 1)
	fmt.Print("array after sort: ")
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	fmt.Println();
}
