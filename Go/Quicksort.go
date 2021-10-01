package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice := []int{100, 2, 5, 7, 8, 34, 64}
	fmt.Println("Unsorted Array: ", slice)
	quicksort(slice)
	fmt.Println("\nSorted Array: ", slice, "\n")
}
 
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    left, right := 0, len(a)-1
     
    pivot := rand.Int() % len(a)
     
    a[pivot], a[right] = a[right], a[pivot]
     
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
     
    a[left], a[right] = a[right], a[left]
     
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}