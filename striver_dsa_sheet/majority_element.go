package striver_dsa_sheet

// For LeetCode 169 — Majority Element, the optimal solution is Boyer–Moore Voting Algorithm.

// Problem

// Given an array nums, return the element that appears more than n / 2 times.

// How it works

// Think of every different element as cancelling out one occurrence of the candidate.
// The majority element cannot be completely cancelled because it occurs more than all other elements combined.

// Complexity
// Time: O(n)
// Space: O(1) ✅ optimal
func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, val := range nums {
		if count == 0 {
			candidate = val

		}

		if candidate == val {
			count++
		} else {
			count--
		}

	}
	return candidate
}
