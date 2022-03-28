package sort

import "github.com/czasg/strualgo/basis"

func HeapSortASC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	for i := len(newNums) / 2; i >= 0; i-- {
		root := i
		left := 2*i + 1
		right := 2*i + 2
		if left < len(newNums) && newNums[left] > newNums[root] {
			root = left
		}
		if right < len(newNums) && newNums[right] > newNums[root] {
			root = right
		}
		if root != i {
			newNums[root], newNums[i] = newNums[i], newNums[root]
		}
	}
	for l := len(newNums) - 1; l >= 0; l-- {
		newNums[0], newNums[l] = newNums[l], newNums[0]
		for i := l / 2; i >= 0; i-- {
			root := i
			left := 2*i + 1
			right := 2*i + 2
			if left < l && newNums[left] > newNums[root] {
				root = left
			}
			if right < l && newNums[right] > newNums[root] {
				root = right
			}
			if root != i {
				newNums[root], newNums[i] = newNums[i], newNums[root]
			}
		}
	}
	return newNums
}

func HeapSortDESC[Number basis.Number](nums []Number) []Number {
	newNums := make([]Number, len(nums))
	for index, num := range nums {
		newNums[index] = num
	}
	for i := len(newNums) / 2; i >= 0; i-- {
		root := i
		left := 2*i + 1
		right := 2*i + 2
		if left < len(newNums) && newNums[left] < newNums[root] {
			root = left
		}
		if right < len(newNums) && newNums[right] < newNums[root] {
			root = right
		}
		if root != i {
			newNums[root], newNums[i] = newNums[i], newNums[root]
		}
	}
	for l := len(newNums) - 1; l >= 0; l-- {
		newNums[0], newNums[l] = newNums[l], newNums[0]
		for i := l / 2; i >= 0; i-- {
			root := i
			left := 2*i + 1
			right := 2*i + 2
			if left < l && newNums[left] < newNums[root] {
				root = left
			}
			if right < l && newNums[right] < newNums[root] {
				root = right
			}
			if root != i {
				newNums[root], newNums[i] = newNums[i], newNums[root]
			}
		}
	}
	return newNums
}
