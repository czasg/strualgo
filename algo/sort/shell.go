package sort

import "github.com/czasg/strualgo/basis"

func ShellSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for step := len(newNums) / 2; step > 0; step /= 2 {
		for i := step; i < len(newNums); i++ {
			index := i
			for index >= i {
				if newNums[index-step] > newNums[index] {
					newNums[index-step], newNums[index] = newNums[index], newNums[index-step]
					index -= step
					continue
				}
				break
			}
		}
	}
	return newNums
}

func ShellSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for step := len(newNums) / 2; step > 0; step /= 2 {
		for i := step; i < len(newNums); i++ {
			index := i
			for index >= i {
				if newNums[index-step] < newNums[index] {
					newNums[index-step], newNums[index] = newNums[index], newNums[index-step]
					index -= step
					continue
				}
				break
			}
		}
	}
	return newNums
}
