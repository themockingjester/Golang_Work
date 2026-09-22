package striver_dsa_sheet

// For LeetCode 229 — Majority Element II, the key difference from Majority Element I is:

// Majority Element I → appears > n/2
// Majority Element II → appears > n/3
// There can be at most 2 such elements.

// The optimal approach is an extension of Boyer–Moore Voting.
func majorityElement2(nums []int) []int {

	// Filtering out the candidates
	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0

	for _, val := range nums {
		if candidate1 == val {
			count1++
		} else if candidate2 == val {
			count2++
		} else if count1 == 0 {
			count1++
			candidate1 = val
		} else if count2 == 0 {
			count2++
			candidate2 = val
		} else {
			count1--
			count2--
		}
	}

	// now need to do verification whether they actually appeared more than n/3 times

	count1, count2 = 0, 0
	for _, val := range nums {
		if candidate1 == val {
			count1++
		} else if candidate2 == val {
			count2++
		}
	}
	result := []int{}
	if count1 > len(nums)/3 {
		result = append(result, candidate1)
	}
	if count2 > len(nums)/3 {
		result = append(result, candidate2)
	}

	return result
}
