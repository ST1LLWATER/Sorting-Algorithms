#include <iostream>
#include <vector>

using std::cout;
using std::cin;
using std::vector;


template <class T>
void merge(vector<T>& list, int start, int mid, int end, int &nComparisons) {
    vector<T> left;
    vector<T> right;
    int leftSize = mid - start + 1;
    int rightSize = end - mid;

    for (int i = 0; i < leftSize; i++) {
        left.push_back(list[start+i]);
    }

    for (int i = 0; i < rightSize; i++) {
        right.push_back(list[i+mid+1]);
    }

    int pos = start;
    int leftPos = 0;
    int rightPos = 0;

    while (leftPos < leftSize && rightPos < rightSize) {
        nComparisons++;    
        if (left[leftPos] < right[rightPos]) {
            list[pos] = left[leftPos];
            leftPos++;
        } else {
            list[pos] = right[rightPos];
            rightPos++;
        }
        pos++;
    }

    while (leftPos < leftSize) {
        list[pos] = left[leftPos];
        leftPos++;
        pos++;
    }

    while (rightPos < rightSize) {
        list[pos] = right[rightPos];
        rightPos++;
        pos++;
    }
}

template <class T>
void mergeSort(vector<T>& list, int start, int end, int &nComparisons) {

    if(end > start) {
        int mid = (start + end) / 2;    
        mergeSort(list, start, mid, nComparisons);    
        mergeSort(list, mid+1, end, nComparisons);    
        merge(list, start, mid, end, nComparisons);
    }
}


int main()
{
    int nComparisons = 0;
    vector<int> list;
    int n;
    cout << "Enter the number of elements in the list: ";
    cin >> n;
    for (int i = 0; i < n; i++) {
        int x;
        cout << "Enter element " << i+1 << ": ";
        cin >> x;
        list.push_back(x);
    }
    cout << "The list is: ";
    for (int i = 0; i < n; i++) {
        cout << list[i] << " ";
    }
    cout << "\n";
    mergeSort(list, 0, n-1, nComparisons);
    cout << "The sorted list is: ";
    for (int i = 0; i < n; i++) {
        cout << list[i] << " ";
    }
    cout << "\n";
    cout << "The number of comparisons is: " << nComparisons << "\n";
    return 0;
}
