package sort

import "github.com/czasg/strualgo/basis"

func MergeSortASC[Number basis.Number](nums []Number) []Number {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSortASC(nums[:mid])
	right := MergeSortASC(nums[mid:])
	newNums := make([]Number, len(left)+len(right))
	l, r := 0, 0
	for i := 0; i < len(left)+len(right); i++ {
		if l < len(left) && r < len(right) {
			if left[l] < right[r] {
				newNums[i] = left[l]
				l++
			} else {
				newNums[i] = right[r]
				r++
			}
		} else if l < len(left) {
			newNums[i] = left[l]
			l++
		} else {
			newNums[i] = right[r]
			r++
		}
	}
	return newNums
}

func MergeSortDESC[Number basis.Number](nums []Number) []Number {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSortDESC(nums[:mid])
	right := MergeSortDESC(nums[mid:])
	newNums := make([]Number, len(left)+len(right))
	l, r := 0, 0
	for i := 0; i < len(left)+len(right); i++ {
		if l < len(left) && r < len(right) {
			if left[l] < right[r] {
				newNums[i] = right[r]
				r++
			} else {
				newNums[i] = left[l]
				l++
			}
		} else if l < len(left) {
			newNums[i] = left[l]
			l++
		} else {
			newNums[i] = right[r]
			r++
		}
	}
	return newNums
}
