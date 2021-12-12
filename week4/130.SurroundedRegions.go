var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}
var visited [][]bool

func turnOver(board *[][]byte, innerBlock *[][]int) {
    for len(*innerBlock) > 0 {
        x := (*innerBlock)[0][0]
        y := (*innerBlock)[0][1]
        (*board)[x][y] = 'X'
        *innerBlock = (*innerBlock)[1:]
    }
}

func bfs(board *[][]byte, x, y int) {
    if (*board)[x][y] == 'O' && !visited[x][y] {
        // board m * n
        m, n := len(*board), len((*board)[0])
        isInnerBlock := true
        innerBlock := make([][]int, 0)
        visited[x][y] = true
        queue := make([][]int, 0)
        queue = append(queue, []int{x, y})
        innerBlock = append(innerBlock, []int{x, y})
        for len(queue) > 0 {
            x, y = queue[len(queue)-1][0], queue[len(queue)-1][1]
            queue = queue[:len(queue)-1]
            for i := 0; i < 4; i++ {
                nx := x + dx[i]
                ny := y + dy[i]
                if nx < 0 || nx >= m || ny < 0 || ny >= n {
                    isInnerBlock = false
                    continue
                }
                if visited[nx][ny] {
                    continue
                }
                if (*board)[nx][ny] == 'X' {
                    continue
                }
                queue = append(queue, []int{nx, ny})
                visited[nx][ny] = true
                innerBlock = append(innerBlock, []int{nx, ny})
            }
        }
        if isInnerBlock {
            turnOver(board, &innerBlock)
        }
    }
}

func solve(board [][]byte) {
    m := len(board)
    n := len(board[0])
    visited = make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            bfs(&board, x, y)
        }
    }
}
