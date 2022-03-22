package sort

import "github.com/czasg/strualgo/basis"

func SelectionSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for i := 0; i < len(newNums)-1; i++ {
		minValue := newNums[i]
		selectIndex := i
		for j := selectIndex + 1; j < len(newNums); j++ {
			if newNums[j] < minValue {
				minValue = newNums[j]
				selectIndex = j
			}
		}
		if selectIndex == i {
			continue
		}
		newNums[selectIndex] = newNums[i]
		newNums[i] = minValue
	}
	return newNums
}

func SelectionSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for i := 0; i < len(newNums)-1; i++ {
		maxValue := newNums[i]
		selectIndex := i
		for j := selectIndex + 1; j < len(newNums); j++ {
			if newNums[j] > maxValue {
				maxValue = newNums[j]
				selectIndex = j
			}
		}
		if selectIndex == i {
			continue
		}
		newNums[selectIndex] = newNums[i]
		newNums[i] = maxValue
	}
	return newNums
}
