func findNumberOfLIS(nums []int) int {
    fn := make([]int, len(nums))
    count := make([]int, len(nums))
    total := 0
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        fn[i] = 1
        count[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if fn[i] < fn[j]+1 {
                    // fmt.Printf("i: %d, j: %d\n", i, j)
                    fn[i] = fn[j] + 1
                    count[i] = count[j]
                } else if fn[i] == fn[j]+1 {
                    // fmt.Printf("-i: %d, j: %d\n", i, j)
                    count[i] += count[j]
                }
            }
        }
        // fmt.Println(i, count)
        if maxLen == fn[i] {
            total += count[i]
        } else if maxLen < fn[i] {
            maxLen = fn[i]
            total = count[i]
        }
    }
    fmt.Println(fn, count)
    return total
}
