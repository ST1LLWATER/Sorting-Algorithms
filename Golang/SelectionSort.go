package main

import "fmt"

/*
	SELECTION SORT ALGORITHM
-----------------------------------
Time complexity - O(n^2)
-----------------------------------
Space complexity - O(1)
*/

/*
Function explantation:
We find the min element present in the unsorted array and  swap it with "ith" elemnth in the array
*/
func SelectionSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		min_index := i
		for j := i; j < n; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
			arr[i], arr[min_index] = arr[min_index], arr[i]		
	}
}

func main() {
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&input[i])
	}
	SelectionSort(input, size)
	fmt.Println(input)

}
