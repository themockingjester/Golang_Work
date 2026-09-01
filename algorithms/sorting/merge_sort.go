package main

func merge(left, right []int) []int {
	leftInd := 0
	rightInd := 0
	result := []int{}

	for true {
		if leftInd < len(left) && rightInd < len(right) {
			if left[leftInd] < right[rightInd] {
				result = append(result, left[leftInd])
				leftInd++
			} else {
				result = append(result, right[rightInd])
				rightInd++
			}
		} else {
			break
		}

	}
	result = append(result, left[leftInd:]...)
	result = append(result, right[rightInd:]...)
	return result

}
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}
