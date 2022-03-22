package sort

import "github.com/czasg/strualgo/basis"

func InsertionSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for i := 1; i < len(newNums); i++ {
		index := i
		for index > 0 {
			if newNums[index-1] > newNums[index] {
				newNums[index-1], newNums[index] = newNums[index], newNums[index-1]
			} else {
				break
			}
			index--
		}
	}
	return newNums
}

func InsertionSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for i := 1; i < len(newNums); i++ {
		index := i
		for index > 0 {
			if newNums[index-1] < newNums[index] {
				newNums[index-1], newNums[index] = newNums[index], newNums[index-1]
			} else {
				break
			}
			index--
		}
	}
	return newNums
}
