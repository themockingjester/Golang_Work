package striver_dsa_sheet

func removeDuplicates(nums []int) int {
	uniqueCnt := 1
	currInd := 1
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			//ignore
			continue
		} else {
			last = nums[i]
			uniqueCnt++
			nums[currInd] = nums[i]
			currInd++
		}
	}
	return uniqueCnt
}
