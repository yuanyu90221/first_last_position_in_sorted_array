package first_last_position

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}

	return []int{searchFirstMatchElement(nums, target), searchLastMatchElement(nums, target)}
}

func searchFirstMatchElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else { // target == nums[target]
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchLastMatchElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else { // target == nums[target]
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
