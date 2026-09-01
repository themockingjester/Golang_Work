package main

func findAndShiftValueInArr(arr []int, value int, backwardInd int) []int {
	i := 0
	for i = backwardInd; i >= 0; i-- {
		if arr[i] > value {

			// Here we are doing elements  shifting
			arr[i+1] = arr[i]
		} else {
			break
		}

	}
	arr[i+1] = value // Correctly placing elment
	return arr

}
func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		arr = findAndShiftValueInArr(arr, arr[i], i-1)
	}
	return arr
}
