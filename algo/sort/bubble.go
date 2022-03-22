package sort

import "github.com/czasg/strualgo/basis"

func BubbleSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	if len(newNums) < 2 {
		return newNums
	}
	for i := 0; i < len(newNums)-1; i++ {
		for j := 0; j < len(newNums)-i-1; j++ {
			if newNums[j] > newNums[j+1] {
				newNums[j], newNums[j+1] = newNums[j+1], newNums[j]
			}
		}
	}
	return newNums
}

func BubbleSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	for i := 0; i < len(newNums)-1; i++ {
		for j := 0; j < len(newNums)-i-1; j++ {
			if newNums[j] < newNums[j+1] {
				newNums[j], newNums[j+1] = newNums[j+1], newNums[j]
			}
		}
	}
	return newNums
}
