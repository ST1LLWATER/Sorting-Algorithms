package main

import "fmt"

/*
	MergeSort Algorithm
	-------------------------------
	Time complexity - O(nlongn)
	-------------------------------
	Space complexity - O(n)
	------------------------------
*/

/*
Algorithm explantion:
We divide the given array into two half recursively and then we call the function merge to merge the array
in sorted manner */

func mergeSort(a []int) []int {
    if len(a) < 2 {
        return a
    }
    m := (len(a)) / 2

    f := mergeSort(a[:m])
    s := mergeSort(a[m:])

    return merge(f, s)
}

func merge(f []int, s []int) []int {
    var i, j int
    size := len(f) + len(s)

    a := make([]int, size, size)

    for z := 0; z < size; z++ {
        lenF := len(f)
        lenS := len(s)

        if i > lenF-1 && j <= lenS-1 {
            a[z] = s[j]
            j++
        } else if j > lenS-1 && i <= lenF-1 {
            a[z] = f[i]
            i++
        } else if f[i] < s[j] {
            a[z] = f[i]
            i++
        } else {
            a[z] = s[j]
            j++
        }
    }

    return a
}

func main() {
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&input[i])
	}
   fmt.Println(mergeSort(input))
}
