func countRangeSum(nums []int, lower int, upper int) int {
    var mergeCount func([]int) int
    mergeCount = func(arr []int) int {
        n := len(arr)
        if n <= 1 {
            return 0
        }

        n1 := append([]int{}, arr[:n/2]...)
        n2 := append([]int{}, arr[n/2:]...)
        cnt := mergeCount(n1) + mergeCount(n2)

        l, r := 0, 0
        for _, v := range n1 {
            for l < len(n2) && n2[l]-v < lower {
                l++
            }
            for r < len(n2) && n2[r]-v <= upper {
                r++
            }
            cnt += r - l
        }

        p1, p2 := 0, 0
        for i := range arr {
            if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
                arr[i] = n1[p1]
                p1++
            } else {
                arr[i] = n2[p2]
                p2++
            }
        }
        return cnt
    }
    prefixSum := make([]int, len(nums)+1)
    for i, v := range nums {
        prefixSum[i+1] = prefixSum[i] + v
    }
    return mergeCount(prefixSum)
}
