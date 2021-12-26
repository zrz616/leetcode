func climbStairs(n int) int {
    steps := []int{1, 2}
    fn := make([]int, n+1)
    fn[0], fn[1] = 1, 1
    for i := 2; i <= n; i++ {
        for _, step := range steps {
            fn[i] += fn[i-step]
        }
    }
    return fn[n]
}
