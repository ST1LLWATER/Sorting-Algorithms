/* INTRODUCTION
------------------
//Insertion sort is a simple sorting algorithm
// The array is virtually split into a sorted and an unsorted part.
// Values from the unsorted part are picked and placed at the correct position in the sorted part.
// It is similar to selection sort where we first find the minimum element and place the minimum element at the beginning. 


// TIME COMPLEXITIES
------------------------
// Worst-case performance	O(n^2)
// Best-case performance	O(n^2)
// Average performance	    O(n)   */
------------------------

    
import java.util.*;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Solution {  

    public static void insertionSort(int[] arr) {
    	int i,j,temp,n=arr.length;
        for(i=0;i<n;i++){
                j=i;
                while(j>0){
                    if(arr[j]<arr[j-1]){
                        temp=arr[j-1];
                        arr[j-1]=arr[j];
                        arr[j]=temp;
                    }
                    j--;
                }
        }
    }

}

public class Runner {

    static BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    
    public static int[] takeInput() throws IOException {
        int size = Integer.parseInt(br.readLine().trim());
        int[] input = new int[size];

        if (size == 0) {
            return input;
        }
        
        String[] strNums; 
        strNums = br.readLine().split("\\s");
        

        for (int i = 0; i < size; ++i) {
            input[i] = Integer.parseInt(strNums[i]);
        }

        return input;
    }

    public static void printArray(int[] arr) {
        for (int element : arr) {
            System.out.print(element + " ");
        }

        System.out.println();
    }

    public static void main(String[] args) throws NumberFormatException, IOException {
        int t = Integer.parseInt(br.readLine().trim());

        while(t > 0) {

            int[] input = takeInput();
            Solution.insertionSort(input);
            printArray(input);

            t -= 1;
        }
    }
}
