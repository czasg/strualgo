package sort

func InsertionSortASC(nums []int) []int {
    newNums := make([]int, len(nums))
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
            }
            index--
        }
    }
    return newNums
}

func InsertionSortDESC(nums []int) []int {
    newNums := make([]int, len(nums))
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
            }
            index--
        }
    }
    return newNums
}
