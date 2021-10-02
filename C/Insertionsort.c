#include<stdio.h>
void insertion_sort(int n,int book[]){
    for(int i=1;i<n;i++){
        int j=i-1; 
        int key =book[i];  //storing current element whose left side is checked for its correct position 
        
        while(book[j]>key && j>=0){ // check whether the adjacent element in left side is greater or less than the current element
           
            book[j+1]=book[j]; // moving the left side element to one position forward
            j--;
        }
        book[j+1]=key;   // moving current element to its  correct position.
    }
}
int main(){
    int n;
    int a[n];
     printf("Enter the number of numbers-->");
     scanf("%d",&n);
    for(int i=0;i<n;i++){
        printf("Enter the Number %d --> ",i+1);
        scanf("%d",&a[i]);
    }
    insertion_sort(n,a);
    printf("Sorted Order of Numbers is : \n");
    for(int i=0;i<n;i++){
      printf("Position %d- %d  \n",i+1,a[i]);
    }
}