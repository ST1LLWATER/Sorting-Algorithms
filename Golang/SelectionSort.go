/*
Two parts are maintained - Sorted Part and Unsorted Part.
In each iteration, the minimal element is picked from the unsorted part and then is put at the end of the sorted part.
In the end, the entire array is sorted.

Time Complexity - O(n*n)

Space Complexity - O(1)
*/

package main

import "fmt"

func selectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
}

func main() {
	sample := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	fmt.Println("Before sorting")
	fmt.Println(sample)
	selectionSort(sample)
	fmt.Println("After sorting")
	fmt.Println(sample)
}
