func shipWithinDays(weights []int, days int) int {
    left, right := 0, 0
    for _, weight := range weights {
        left = max(left, weight)
        right += weight
    }

    for left < right {
        mid := left + (right-left)/2
        if validate(&weights, days, mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return right
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func validate(weights *[]int, days int, k int) bool {
    count := 1
    box := 0
    for _, weight := range *weights {
        if weight > k {
            return false
        }
        if box+weight <= k {
            box += weight
        } else {
            box = weight
            count++
        }
    }
    return count <= days
}
