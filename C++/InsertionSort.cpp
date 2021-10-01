#include <iostream>

using std::cout;
    
// Insertion Sort Algorithm
void insertionSort(int arr[], int n)
{
    int i, key, j;
    for (i = 1; i < n; i++)
    {
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
    int arr[] = { 12, 11, 13, 5, 6, 7 };
    cout << "Sorted array is \n";
    insertionSort(arr, sizeof(arr) / sizeof(arr[0]));
    for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++)
        cout << arr[i] << " ";
    return 0;
}
