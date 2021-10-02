package main

import "fmt"

/*

BUBBLE SORT ALGORITHM
--------------------------
Time complexity
-Best case  O(n)
-worst case O(n^2)
---------------------------
Space complexity - O(1)

*/

/*
Function explanation:
We simply iterate through array and repedtedly swap the elements if they are in wrong order
*/
func BubbleSort(arr []int, n int) {
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--
	}
}

func main() {
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&input[i])
	}
	BubbleSort(input, size)
	fmt.Println(input)
}
