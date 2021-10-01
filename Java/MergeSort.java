// INTRODUCTION
------------------
// Merge Sort is a sorting technique which works on Divide and Conquer Algorithm.
// We take first and second parts(divide from mid) and lets say we have sorted both the arrays,
// now we compare each element from the other element in the 2nd list and add it in order.
------------------




// TIME COMPLEXITIES
------------------------
//The time complexity of merge sort is same for all the cases(best,worst,average)
// -- O(n*log n).
------------------------






// CODE
---------------------------------------------------------------
// Java program for implementation of Merge Sort

public class MergeSort {

    static void merge( int [] arr , int l , int mid , int r ){

        int SIZE = arr.length;
        int [] b = new int[SIZE];

        int i = l;
        int j = mid+1;
        int k = l;

        while ( i <= mid && j <= r){

            if ( arr[i] < arr[j]){
                b[k] = arr[i];
                i++;
            }
            else{
                b[k] = arr[j];
                j++;
            }
            k++;
        }

        if ( i > mid){
            while ( j<= r){
                b[k] = arr[j];
                k++; j++;
            }
        }
        else{
            while ( i<= mid){
                b[k] = arr[i];
                k++; j++;
            }
        }

        for ( int x = l ; x < r ; x++){
            arr[x] = b[x];
        }

    }


    static void mSort( int [] arr , int l , int r ){

        if ( l < r ){

            int mid = (l+r)/2;

            mSort(arr, l , mid);
            mSort(arr , mid+1 , r);
            merge(arr, l , mid , r );
        }
    }

}
