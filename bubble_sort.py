# Python program for implementation of Bubble Sort

def bubbleSort(arr):
	n = len(arr)

	# Traverse through all array elements
	for i in range(n-1):
		# range(n) also work but outer loop will repeat one time more than needed.

		# Last i elements are already in place
		for j in range(0, n-i-1):

			# traverse the array from 0 to n-i-1
			# Swap if the element found is greater
			# than the next element
			if arr[j] > arr[j + 1]:
				arr[j], arr[j + 1] = arr[j + 1], arr[j]


# Driver code to test above
arr = [64, 34, 25, 12, 22, 11, 90]
print("Array given is:-")
# used list comprehension to print it in one line
[print(i, end=" ") for i in arr]
print()
bubbleSort(arr)
print("Sorted array is:-")
# used list comprehension to print it in one line
[print(i, end=" ") for i in arr]
