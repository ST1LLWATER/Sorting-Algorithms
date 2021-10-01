package main

func insertionSortAlpha(arr []int) []int {
	//  Insertion sort: Works from left to right, moving the current index value to the correct position in each pass.
	//  Time Complexity: Best - O(n) Average - O(n^2) Worst - O(n^2)

	//  Declare that i, j, current are type int
	var i, j, current int

	//  Loop starting at position 1 as the value at position 0 is theoretically in a "sorted" portion of the array
	for i = 1; i < len(arr); i++ {

		//  Define the current value to sort around
		current = arr[i]

		//  Set the second index to be current - 1
		j = i - 1

		//	While loops are not implemented in Go so an alternative approach must be used
		for j >= 0 && arr[j] > current {

			//  Perform swap
			arr[j+1] = arr[j]

			//  Decrement index value of j
			j -= 1
		}

		//  Move current value into second indexer + 1 to place in correct position
		arr[j+1] = current
	}

	return arr
}

func insertionSortBeta(arr []int) []int {
	//  Insertion sort: Works from left to right, moving the current index value to the correct position in each pass.
	//  Time Complexity: Best - O(n) Average - O(n^2) Worst - O(n^2)

	//  Loop starting at position 1 as the value at position 0 is theoretically in a "sorted" portion of the array
	for i := 1; i < len(arr); i++ {

		j := i

		//	While loops are not implemented in Go so an alternative approach must be used
		for j > 0 {
			//	If the value in position j is less than the value proceeding it, swap the values
			if arr[j] < arr[j-1] {
				//	Swap the values of index j and j-1
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}

			//	Decrement j
			j -= 1
		}
	}
	//	Return sorted array
	return arr
}
