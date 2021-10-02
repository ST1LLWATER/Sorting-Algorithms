/*
Introduction:
  Radix sort is a sorting algorithm that sorts numbers by the positions of their digits.
  The way it works is that it uses the place value of the digits in a number (from least significant digit to most significant digit)
  In contrast to other sorting algorithms, Merge Sort, Insertion Sort, Bubble Sort,
  the number are not compared.
  It can also be applied to string with just a little change of the place value used (here we go from most significant to least significant digit).
  To sort the digits, Radix sort uses a stable sorting algorithm as a subroutine.
  Using a variation of cing sort, I have sorted the digits using the radix in every position.

Time: O(nk), where n is the number of elements and k is the number of bits per element
Space: O(n + k), where n is the number of elements and k is the number of bits per element
*/


import (
	"fmt"
	"strconv"
)

// Finds the largest number in an array
func findLargestNum(array []int) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

// Radix Sort
func radixSort(array []int) []int {

  fmt.Println("Running Radix Sort on Unsorted List\n")

  // Base 10 is used
	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for largestNum / significantDigit > 0 {
  
    fmt.Println("\tSorting: " + strconv.Itoa(significantDigit) + "'s place", array)

		bucket := [10]int{0}

		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(array[i] / significantDigit) % 10]++
		}

		// Add the count of the previous buckets
    // Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i - 1] 
		}

		// Use the bucket to fill a "semiSorted" array
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i] / significantDigit) % 10]--
			semiSorted[bucket[(array[i] / significantDigit) % 10]] = array[i]
		}

    // Replace the current array with the semisorted array
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
    
    fmt.Println("\n\tBucket: ", bucket)
  
    // Move to next significant digit
		significantDigit *= 10 
	}

	return array
}

func main() {

  fmt.Println("\n\nRunning Radix Sort Example in Go!")
  fmt.Println("----------------------------------\n")

	unsortedList :=[]int{10, 2, 303, 4021, 293, 1, 0, 429, 480, 92, 2999, 14}
	fmt.Println("Unsorted List:", unsortedList, "\n")

	sortedList := radixSort(unsortedList)

	fmt.Println("\nSorted List:", sortedList, "\n")
}
