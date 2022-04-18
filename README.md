# Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

## 初步想法

一開始想說要一次找到兩個 range 

所以一直在於怎麼去做到把找到的 target 給延伸

結果就卡住了

後來看到 leetcode cookbook 的說明後 明白到其實可以分開找尋

拆成 1 找到第一個符合的整數位置 2 找到第二個最後一個符合的整數位置

由於都是用 Sorted Array 去找

用二元搜尋法加上不同的結束條件就可以個別去找到了

所以時間複雜度 = Max(O(logN) , O(logN)) = O(logN)

### 找到第一個符合的整數位置


Step 1: 初始化 high = len(nums) - 1, low = 0 

Step 2: 檢查 low <= high , 不成立則, 跳到 Step 8 

Step 3: 計算 mid = (high+low) / 2

Step 4: 每次找到去比對 nums[mid] 與 target

Step 5: 如果 nums[mid] > target , 則更新 high = mid - 1 , 跳到 Step 2 

Step 6: 如果 nums[mid] < target , 則更新 low = mid + 1 , 跳到 Step 2

Step 7: 如果 nums[mid] == target , 檢查 mid == 0 或是 nums[mid - 1] != target

        如果符合, 則 mid 則是最小使得 nums[mid] 等於 target 的 index, 回傳 mid

        如果不符合, 則需要把 high 更新為 mid - 1 , 跳到 Step 2
Step 8: 回傳 -1

### 找到最後一個符合的整數位置


Step 1: 初始化 high = len(nums) - 1, low = 0 

Step 2: 檢查 low <= high , 不成立則, 跳到 Step 8 

Step 3: 計算 mid = (high+low) / 2

Step 4: 每次找到去比對 nums[mid] 與 target

Step 5: 如果 nums[mid] > target , 則更新 high = mid - 1 , 跳到 Step 2 

Step 6: 如果 nums[mid] < target , 則更新 low = mid + 1 , 跳到 Step 2

Step 7: 如果 nums[mid] == target , 檢查 mid == len(nums) - 1 或是 nums[mid + 1] != target

        如果符合, 則 mid 則是最大使得 nums[mid] 等於 target 的 index, 回傳 mid

        如果不符合, 則需要把 low 更新為 mid + 1 , 跳到 Step 2
Step 8: 回傳 -1

## 實作

```golang
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
```