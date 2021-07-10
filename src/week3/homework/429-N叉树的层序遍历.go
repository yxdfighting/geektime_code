package homework

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

/*
   每一层按照从左到右的顺序进入 按照从左到右的顺序提出 符合队列的特点 FIFO
   对于每一个节点，每一层依次入队 出队的时候将子节点入队
*/

type LevelNode struct {
	node  *Node
	level int
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	//初始化queue
	queue := make([]LevelNode, 0)
	//初始化空间
	result := make([][]int, 1000)

	idx, max := 0, 0
	//节点入队
	rootLevel := LevelNode{root, idx}
	queue = append(queue, rootLevel)

	for len(queue) != 0 {
		tempRoot := queue[0]
		max = tempRoot.level
		//每次拿到队首元素append到result对应的idx数组
		result[tempRoot.level] = append(result[tempRoot.level], tempRoot.node.Val)
		//节点出队
		queue = queue[1:len(queue)]
		//子节点入队
		for _, child := range tempRoot.node.Children {
			levelNode := LevelNode{child, tempRoot.level + 1}
			queue = append(queue, levelNode)
		}
	}

	return result[0 : max+1]
}
