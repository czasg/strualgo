package sort

import (
	"github.com/czasg/strualgo/basis"
)

func QuickSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	var quickSortASC func(nums []Number)
	quickSortASC = func(nums []Number) {
		if len(nums) < 2 {
			return
		}
		left := 0
		right := len(nums) - 1
		x := nums[left]
		for {
			for left < right && nums[right] >= x {
				right--
			}
			if left >= right {
				break
			}
			nums[left] = nums[right]
			left++
			for left < right && nums[left] <= x {
				left++
			}
			if left >= right {
				break
			}
			nums[right] = nums[left]
			right--
		}
		nums[left] = x
		quickSortASC(nums[:left])
		quickSortASC(nums[right+1:])
	}
	quickSortASC(newNums)
	return newNums
}

func QuickSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	var quickSortDESC func(nums []Number)
	quickSortDESC = func(nums []Number) {
		if len(nums) < 2 {
			return
		}
		left := 0
		right := len(nums) - 1
		x := nums[left]
		for {
			for left < right && nums[right] <= x {
				right--
			}
			if left >= right {
				break
			}
			nums[left] = nums[right]
			left++
			for left < right && nums[left] >= x {
				left++
			}
			if left >= right {
				break
			}
			nums[right] = nums[left]
			right--
		}
		nums[left] = x
		quickSortDESC(nums[:left])
		quickSortDESC(nums[right+1:])
	}
	quickSortDESC(newNums)
	return newNums
}
