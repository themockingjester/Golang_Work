package striver_dsa_sheet

func check(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	//Holding the first element as we know the array originally was sorted in asc
	start := 0

	//Denotes whether array was disturbed or not
	isDefect := false
	for i := 1; i < len(nums); i++ {
		// the place where we found the defect
		if !isDefect && nums[i] < nums[i-1] {
			isDefect = true

			// checking defect creating elment lies between start and just one elment before range (as that was the sorted array) if yes then we can return false
			if nums[i] > nums[start] && nums[i] < nums[i-1] {
				return false
			}
			// move on
			continue
		} else if isDefect {
			// from here onwards we are iterating on the defective array

			// If curent elment lies outside the range of last array (whcih is sorted) then things are fine
			if nums[i] <= nums[start] && nums[i] >= nums[i-1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
