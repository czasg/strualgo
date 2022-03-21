package sort

func BubbleSortASC(nums []int) []int {
    newNums := make([]int, len(nums))
    for index, num := range nums {
        newNums[index] = num
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

func BubbleSortDESC(nums []int) []int {
    newNums := make([]int, len(nums))
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
