package striver_dsa_sheet

// Maximum Product Subarray (LeetCode 152).

// The issue is that, unlike maximum sum, a negative number can turn the smallest negative product into the largest positive product.
func maxProduct(nums []int) int {
	currMaxProduct := nums[0]
	currMinProduct := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		currVal := nums[i]
		if currVal < 0 {
			// swapping both single negative value can turn max to min and vice versa
			currMaxProduct, currMinProduct = currMinProduct, currMaxProduct
		}

		currMaxProduct = max(currVal, currVal*currMaxProduct)
		currMinProduct = min(currVal, currVal*currMinProduct)

		result = max(currMaxProduct, result)
	}
	return result
}
