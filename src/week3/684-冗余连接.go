package week3

/*
   1. 每次加一条边，就通过从头开始的递归判断是否有环，如果有环就输出当前加的边
   2. 对于无向图的表示，通过出边数组，1-2  分别在2的出边数组里加1  在1的出边数组里加2
*/

var edges [][]int
var visit []bool
var hasCycle bool

func findRedundantConnection(input [][]int) []int {
	hasCycle = false
	//初始化出边数组和visit数组
	visit = make([]bool, len(input)+1)
	// edges = make([][]int,len())
	edges = edges[0:0]
	for idx := 0; idx < len(input)+1; idx++ {
		var temp []int
		edges = append(edges, temp)
		visit[idx] = false
	}

	//遍历数组
	for idx, edge := range input {
		//加边
		addEdge(edge[0], edge[1])
		// fmt.Printf("edge:%v\n,edges:%v\n",edge,edges)
		//注意每次找环的起点应该是当前加入的点，而非1
		dfs(edge[0], edge[1])
		if hasCycle {
			return input[idx]
		}
	}
	return []int{}
}

//定义深度遍历递归方法,递归结束条件
//1.找到环  2.找完所有的边
func dfs(x int, father int) {
	//将x置为已经访问
	visit[x] = true
	for _, edge := range edges[x] {
		//如果在x的出边中找到了父节点，不算有环
		if edge == father {
			continue
		}
		//表示访问到了除父亲节点外的 已访问节点
		if visit[edge] {
			hasCycle = true
			return
		}
		visit[edge] = true
		dfs(edge, x)
	}
	//递归完一定要还原现场
	visit[x] = false
	return
}

//定义一个加边方法
func addEdge(x, y int) {
	edges[x] = append(edges[x], y)
	edges[y] = append(edges[y], x)
}
