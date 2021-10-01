import java.util.*;
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
