package two_pointers

// squaring_a_sorted_array takes a sorted array and returns a new array
// containing the squares of all numbers, also in sorted order.
func squaring_a_sorted_array_solution_1(nums []int) []int {
	// 1. Find the index of the first non-negative number (the "pivot")
	i := len(nums)
	for index, num := range nums {
		if num >= 0 {
			i = index
			break
		}
	}

	ans := make([]int, 0, len(nums)) // Pre-allocate capacity for efficiency
	j := i - 1                       // Pointer for negative numbers (moving left)
	// j and to its left has negative elements
	// i and to its right has positive elements

	// 2. Use two pointers to compare absolute values and merge squares
	for j >= 0 && i < len(nums) {
		if nums[i] < -nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i++
		} else if nums[i] > -nums[j] {
			ans = append(ans, nums[j]*nums[j])
			j--
		} else {
			ans = append(ans, nums[i]*nums[i], nums[j]*nums[j])
			i++
			j--
		}
	}

	// 3. Add remaining squares from the negative side (if any)
	for j >= 0 {
		ans = append(ans, nums[j]*nums[j])
		j--
	}

	// 4. Add remaining squares from the positive side (if any)
	for i < len(nums) {
		ans = append(ans, nums[i]*nums[i])
		i++
	}

	return ans
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func squaring_a_sorted_array_solution_2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1

	// Fill the result array from back to front
	// since we can only find the largest squares by comparing nums[left] and nums[right] so we insert from back.
	for i := n - 1; i >= 0; i-- {
		if absInt(nums[left]) > absInt(nums[right]) {
			ans[i] = nums[left] * nums[left]
			left++
		} else {
			ans[i] = nums[right] * nums[right]
			right--
		}
	}
	return ans
}
