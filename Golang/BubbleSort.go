package main
import "fmt"

func BubbleSort(arr []int, n int){
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
				if(arr[j]<arr[i]){
					arr[j],arr[i]=arr[i],arr[]
				}
		}l
	}
}