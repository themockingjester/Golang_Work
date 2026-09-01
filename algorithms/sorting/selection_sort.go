package main

func selectionSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		smaller := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smaller] {
				smaller = j
			}
		}
		temp := arr[i]
		arr[i] = arr[smaller]
		arr[smaller] = temp
	}
	return arr
}
