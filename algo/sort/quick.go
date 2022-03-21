package sort

import (
	"github.com/czasg/strualgo/basis"
)

func QuickSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	var quickSortASC func(nums []Number)
	quickSortASC = func(nums []Number) {
		if len(nums) < 2 {
			return
		}
		left := 0
		right := len(nums) - 1
		tmpNum := nums[left]
		for {
			for left < right && nums[right] >= tmpNum {
				right--
			}
			if left >= right {
				break
			}
			nums[left] = nums[right]
			left++
			for left < right && nums[left] <= tmpNum {
				left++
			}
			if left >= right {
				break
			}
			nums[right] = nums[left]
			right--
		}
		nums[left] = tmpNum
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
		tmpNum := nums[left]
		for {
			for left < right && nums[right] <= tmpNum {
				right--
			}
			if left >= right {
				break
			}
			nums[left] = nums[right]
			left++
			for left < right && nums[left] >= tmpNum {
				left++
			}
			if left >= right {
				break
			}
			nums[right] = nums[left]
			right--
		}
		nums[left] = tmpNum
		quickSortDESC(nums[:left])
		quickSortDESC(nums[right+1:])
	}
	quickSortDESC(newNums)
	return newNums
}
