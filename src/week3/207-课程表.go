package week3

/*
	构建一个有向图，每一个元素后者指向前者
	1、思考一个问题？这个题并没有定义visit数组
*/

var edges [][]int

//表示每个点的入度
var dig []int

func canFinish(numCourses int, prerequisites [][]int) bool {
	dig = make([]int, numCourses)
	edges = edges[0:0]
	for idx := 0; idx < numCourses; idx++ {
		var temp []int
		edges = append(edges, temp)
	}

	//根据prerequisites
	for _, prerequisite := range prerequisites {
		//1指向0
		addEdge(prerequisite[0], prerequisite[1])
	}

	if bfs() == numCourses {
		return true
	}

	return false
}

func bfs() int {
	var queue []int
	//表示零度点的数量
	length := 0
	//先从入度为0的节点开始
	for x, d := range dig {
		if d == 0 {
			queue = append(queue, x)
		}
	}

	for len(queue) != 0 {
		//队首出队
		head := queue[0]
		queue = queue[1:]
		length++
		for _, child := range edges[head] {
			dig[child] = dig[child] - 1
			if dig[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	return length
}

func addEdge(x, y int) {
	edges[y] = append(edges[y], x)
	dig[x] = dig[x] + 1
}
