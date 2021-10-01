/* INTRODUCTION
 * ------------
 * Bubble sort, sometimes referred to as sinking sort, is a simple
 * sorting algorithm that repeatedly steps through the list,
 * compares adjacent elements and swaps them if they are in the wrong order.
 * The pass through the list is repeated until the list is sorted.
 * The algorithm, which is a comparison sort, is named for the way smaller
 * or larger elements "bubble" to the top of the list.
 *
 * // TIME COMPLEXITIES
 * ------------------------
 * // Worst-case performance	O(n^2) comparisons and swaps
 * // Best-case performance	O(n) comparisons and O(1) swaps
 * // Average performance	O(n^2) comparisons and swaps
 * ------------------------
 *
 *
 */

#include <iostream>

void BubbleSort(int *array) {
  for (int i = 0; i < 10; i++) { // go through all elements to the end
    for (int j = i + 1; j < 10;
         j++) { // from the already passed elements to the end
      if (array[j] < array[i]) { // if next element < current element. We can use > for downting sorting
        int temp;
        temp = array[i]; // then changing their places
        array[i] = array[j];
        array[j] = temp;
      }
    }
  }
}

int main() {
  int arr[10] = {2, 5, 7, 1, 3,
                 9, 4, 6, 8, 10}; // creating array of 10 non-sorted elements
  for (int i = 0; i < 10; i++)
    std::cout << arr[i] << " "; // just printing created array

  BubbleSort(arr); // calling the function, which will BubbleSort the array
  std::cout << "Sorted array: " << std::endl;
  for (int i = 0; i < 10; i++)
    std::cout << arr[i] << " "; // printing sorted array
  return 0;
}
