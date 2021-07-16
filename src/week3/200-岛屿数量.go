package week3

/*
   从每一个陆地开始，进行dfs遍历，标记visit数组
   每个dfs返回本次遍历的岛屿个数

	本题有个疑问，visit数组为啥没有还原？？？？？？
*/
var visited [][]bool

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func numIslands(grid [][]byte) int {
	result := 0
	visited = visited[0:0]
	for i := 0; i < len(grid); i++ {
		temp := make([]bool, len(grid[0]))
		visited = append(visited, temp)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j]-'0' == 1 && !visited[i][j] {
				dfs(i, j, grid)
				result++
			}
		}
	}
	return result
}

//从ij位置开始dfs
func dfs(i, j int, grid [][]byte) {
	if visited[i][j] || grid[i][j]-'0' == 0 {
		return
	}

	visited[i][j] = true

	//向上下左右延伸
	for m := 0; m < 4; m++ {
		if (i+dx[m] >= 0 && i+dx[m] < len(grid)) &&
			(j+dy[m] >= 0 && j+dy[m] < len(grid[0])) {
			dfs(i+dx[m], j+dy[m], grid)
		}
	}

	// visited[i][j] = false
}
