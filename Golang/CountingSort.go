package main

import "fmt"

/*COUNTING SORT ALGORITHM
-------------------------
 Introduction
- It is a sorting algorithm that is used for sorting
  elements within a specific range and is based on the 
  frequency/count of each element to be sorted.
--------------------------
Time complexity
- Best case: O(n)
- Average case: O(N+K)
- Worst case: O(n^2)
---------------------------
Space complexity - O(K)
*/


func countSort(arr []int) []int {
    var max = arr[0]
    var i = 1
    for i < len(arr) {
        if arr[i] > max {
            max = arr[i]
        }
        i++
    }

    var indices = make([]int, max + 1)
    var j = 0
    for j < len(arr) {
        indices[arr[j]]++
        j++
    }

    var k = 1
    for k < len(indices) {
        indices[k] += indices[k - 1]
        k++
    }

    var result = make([]int, len(arr))
    var m = 0
    for m < len(arr) {
        result[indices[arr[m]] - 1] = arr[m]
        indices[arr[m]]--
        m++
    }
    return result
}

func main() {
    var n = []int{1, 39, 2, 9, 7, 54, 11}
    fmt.Println(countSort(n))
}