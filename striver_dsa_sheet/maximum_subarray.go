package striver_dsa_sheet

// For LeetCode 53 — Maximum Subarray, the optimal solution is Kadane's Algorithm.
// Core idea

// At every element, decide:

// Should I extend the existing subarray, or start a new subarray from here?
func maxSubArray(nums []int) int {
	current := nums[0]
	maxi := nums[0]

	for i := 1; i < len(nums); i++ {
		current = max(nums[i], current+nums[i])
		if current > maxi {
			maxi = current
		}

	}
	return maxi
}
