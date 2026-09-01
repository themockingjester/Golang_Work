package main

func quickSort(arr []int, low, high int) []int {

	if low >= high {
		return arr
	}
	pivotVal := arr[high]
	i := low - 1   // it tracks smaller values than pivot values
	j := low       // it tracks larger values than the pivot values
	for j < high { // iterating through every element till pivot value
		if arr[j] < pivotVal {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)

	return arr
}
