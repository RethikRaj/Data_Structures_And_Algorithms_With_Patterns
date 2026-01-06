package two_pointers

// two_sum_sorted finds a pair in arr with a sum equal to targetSum.
// Returns the indices of the pair if found, otherwise returns [-1, -1].
func two_sum_sorted(nums []int, target int) []int {
	leftPtr := 0
	rightPtr := len(nums) - 1

	ans := make([]int, 2)

	for leftPtr < rightPtr {
		sum := nums[leftPtr] + nums[rightPtr]

		if sum < target {
			leftPtr += 1
		} else if sum > target {
			rightPtr -= 1
		} else {
			ans[0] = leftPtr
			ans[1] = rightPtr
		}
	}

	ans[0] = -1
	ans[1] = -1
	return ans
}
