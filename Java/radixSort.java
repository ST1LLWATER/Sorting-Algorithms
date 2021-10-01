/*

Introduction:

  Radix sort is a sorting algorithm that sorts numbers by the positions of their digits.
  The way it works is that it uses the place value of the digits in a number (from least significant digit to most significant digit)
  In contrast to other sorting algorithms, Merge Sort, Insertion Sort, Bubble Sort,
  the number are not compared.

  It can also be applied to string with just a little change of the place value used (here we go from most significant to least significant digit).

  To sort the digits, Radix sort uses a stable sorting algorithm as a subroutine.
  Using a variation of cing sort, I have sorted the digits using the radix in every position.

*/


/*
Time and Space Complexity:

  Time Complexity: O(d * (n + b));
  	
    Best Case: Ω((n+b)*d)
    Average Case: θ((n+b)*d)
    Worst Case: O((n+b)*d)

    here: d -> no. of digit/char of the largest element in the array.
          n -> no. of elements in the array.
          b -> base of the elements(for example: 10 for decimal no.s, 26 for alphabets)

  Space Complexity: O(n + b);
    here: n -> no. of elements in the array.
          b -> base of the elements(for example: 10 for decimal no.s, 26 for alphabets)
    
  The performance of Radix Sort depends on the stable sorting algorithm chosen to sort the digits.
*/


// Code:

import java.util.*;

public class radixSort {

	int maxElement(int arr[], int n)
	{
		int mx = arr[0];
		for (int i = 1; i < n; i++)
			if (arr[i] > mx)
				mx = arr[i];
		return mx;
	}

	void cSort(int arr[], int n, int exp)
	{
		int output[] = new int[n];
		int i;
		int c[] = new int[10];
		Arrays.fill(c, 0);

		// storing count of elements' occurences
		for (i = 0; i < n; i++)
			c[(arr[i] / exp) % 10]++;

		// creating actual position of this digit in output[]
		for (i = 1; i < 10; i++)
			c[i] += c[i - 1];

		// creating the output array
		for (i = n - 1; i >= 0; i--) {
			output[c[(arr[i] / exp) % 10] - 1] = arr[i];
			c[(arr[i] / exp) % 10]--;
		}

    // copying sorted output array to original array "arr"
		for (i = 0; i < n; i++)
			arr[i] = output[i];
	}

	void radixsort(int arr[], int n)
	{
		// fetching no. of digits of largest element in array.
		int m = maxElement(arr, n);

	  // performing count sort for every digit.
		// Moreoverm, instead of passing digit number, exp is passed.
		// exp is 10^i where i is current digit number
		for (int exp = 1; m / exp > 0; exp *= 10)
			cSort(arr, n, exp);
	}


	public static void main(String[] args)
	{
    Scanner sc = new Scanner(System.in);

    System.out.print("Enter no. of elements in the array: ");
		int n = sc.nextInt();
    int[] arr = new int[n];
    System.out.println("Enter elements in the array: ");
    for(int i=0;i<n;i++){
      arr[i] = sc.nextInt();
    }

    // making object of the class
    radixSort obj = new radixSort();

    // calling the sort function
		obj.radixsort(arr, n);

    // printing the sorted array
    for (int i = 0; i < n; i++)
			System.out.print(arr[i] + " ");
	}
}

