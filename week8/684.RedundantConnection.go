func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    fa := make([]int, n+1)
    for i := 0; i <= n; i++ {
        fa[i] = i
    }
    ans := []int{0, 0}
    for _, v := range edges {
        if find(fa, v[0]) != find(fa, v[1]) {
            unionSet(fa, v[0], v[1])
        } else {
            ans = v
        }
    }
    return ans
}

func find(fa []int, i int) int {
    if fa[i] == i {
        return i
    }
    fa[i] = find(fa, fa[i])
    return fa[i]
}

func unionSet(fa []int, i, j int) {
    i, j = find(fa, i), find(fa, j)
    if i != j {
        fa[i] = j
    }
}
