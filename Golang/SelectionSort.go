package main

import "fmt"

/*
 Insertion Sort
 ---------------------------------
 Time Complexity - O(n^2)
 ---------------------------------
 Space Complexity - O(1)
 ---------------------------------
*/

/*
Function explanatin:
We check if preceding element is greater than then current element
and we swap them if the conditon is true
*/
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&input[i])
	}
	SelectionSort(input)
	fmt.Println(input)
}
