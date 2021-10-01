// INSERTION SORT- used for small size arrays

// Here you divide the array into 2 parts, sorted and unsorted. 

// Then you start picking up elements from unsorted part and place them in the proper place inside 
// the sorted part, making sure that the sorted part stays sorted. 

//  You continue this untill your array is completely sorted.

//Algorithm (ascending)
// 1: Iterate from the second element to the last element in the array
// 2: Compare current element (key) with the element before it
// 3: If key<predecessor, compare key with the elements before. 
//    ( We're moving the greater elements one position up to make space for the swapped element.


#include <math.h>
#include <stdio.h>

// Time Complexity = O(n^2) , worst case - array is sorted in reverse order, best case, array sorted.
// Auxiliary Space Needed = O(1)

void insertionSort(int arr[], int n)
{
	int i, key, j;
	for (i = 1; i < n; i++) {
		key = arr[i];
		j = i - 1;
		while (j >= 0 && arr[j] > key) 
		{
			arr[j + 1] = arr[j];
			j = j - 1;
		}
		arr[j + 1] = key;
	}
}



int main()
{
	int m;
	printf("Enter size of array : ");
	scanf("%d",&m);
	int arr[m];
	printf("\nEnter array : ");
	for(int i =0; i<m; i++){
	    scanf("%d",&arr[i]);
	}
	
	int n = sizeof(arr) / sizeof(arr[0]);

	insertionSort(arr, n);
	for ( int i = 0; i < n; i++)
		printf("%d ", arr[i]);
	printf("\n");

	return 0;
}
