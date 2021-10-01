package main

import "fmt"

/*
	QUICK SORT ALOGRITHM
TIME COMPLEXITIES
------------------------
 Worst-case performance	          O(n^2)
 Best-case performance	  		  O(n*log(n))
 Average performance	          O(n*log(n))
------------------------

 SPACE COMPLEXITIES
------------------------
 Basic Approach        		O(n)
 Tail Recursion Approach    O(logn)
*/

//---------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------

/*
	partition function
-----------------------
	parms
		-array
		-lower bound
		-upper bound
	return type int
-----------------------
function explanation:
The function takes last element as pivot, places
the pivot element at its correct position in sorted
array, and places all smaller elements compared to pivot
to left of pivot and all greater elements to right
----------------------------------------------------------------
*/
func partition(arr []int, l, h int) int {
	pivot := arr[h]
	i := l
	for j := l; j < h; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[h] = arr[h], arr[i]
	return i
}

/*
	quick sort function
----------------------------------------
	parms
		-array
		-lower bound
		-upper bound
----------------------------------------
*/
func quickSort(arr []int, l int, h int) {
	for l < h {
		/* pi is partitioning index, arr[p] is now
		   at right place */
		pi := partition(arr, l, h)
		/* If left part is smaller, then recur for left
		part and handle right part iteratively*/
		if pi-l < h-pi {
			quickSort(arr, l, pi-1)
			l = pi + 1
		} else {
			//Else recur for right part
			quickSort(arr, pi+1, h)
			h = pi - 1
		}
	}
}

func main() {
	var size int
	_, err := fmt.Scanln(&size)
	if err != nil {
		return
	}
	input := make([]int, size)
	for i := 0; i < size; i++ {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			return
		}
	}
	quickSort(input, 0, size-1)
	fmt.Println(input)
}
