// INTRODUCTION
------------------
// Merge Sort is a sorting technique which is a little more efficient than bubble sort.
// It swaps the minimum element (considering ascending order) from unsorted part and putting it at the beginning.
------------------




// TIME COMPLEXITIES
------------------------
//The time complexity of Selection sort is same for all the cases(best,worst,average)
// -- O(n^2).
------------------------






// CODE
---------------------------------------------------------------
// Java program for implementation of Selection

public class SelectionSort {
    public static void main(String[] args) {

        int [] a = { 5 ,2 , -1 , 6 , 4 , 3 , 1 , 0 };
        int n = a.length;


        for ( int i = 0 ; i < n - 1 ; i++ ){

            int minIndex = i;

            for ( int j = i + 1 ; j < n ; j++){

                if ( a[j] < a[minIndex]){
                    minIndex = j;
                }

            }
            int temp = a[i];
            a[i] = a[minIndex];
            a[minIndex] = temp;
        }





    }
}