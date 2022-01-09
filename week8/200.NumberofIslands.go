func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    fa := make([]int, m*n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            a := i*n + j
            if grid[i][j] == '0' {
                fa[a] = -1
            } else {
                fa[a] = a
            }
        }
    }

    dx := []int{-1, 0}
    dy := []int{0, 1}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            a := i*n + j
            if grid[i][j] == '0' {
                continue
            }
            for k, _ := range dx {
                ni := i + dx[k]
                nj := j + dy[k]
                na := ni*n + nj
                if ni < 0 || nj < 0 || ni >= m || nj >= n {
                    continue
                }
                if grid[ni][nj] == '0' {
                    continue
                }
                unionSet(fa, a, na)
            }
        }
    }

    ans := 0
    for i, v := range fa {
        if i == v {
            ans++
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
        fa[j] = i
    }
}
