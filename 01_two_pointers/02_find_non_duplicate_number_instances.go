package two_pointers

// This function moves all unique elements to the beginning of a sorted array in-place using two pointers.
// It returns the length of the unique subarray, maintaining the original order with O(1) space complexity.
func find_non_duplicate_number_instances(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := 1

	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i += 1
			j += 1
		} else {
			j += 1
		}
	}

	return i + 1
}
