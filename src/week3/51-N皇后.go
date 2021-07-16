package week3

/*
   同一行同一列不能有皇后
   考虑如何转化为状态空间的查找？
   以行的维度,目标是把每一行都摆上皇后，同时需要满足在竖直和斜线方向不能有其他皇后
*/

var result [][]string
var oneRes []string

func solveNQueens(n int) [][]string {
	result = result[0:0]
	oneRes = oneRes[0:0]
	putQueen(0, n)
	return result
}

func putQueen(col, n int) {
	//递归终止，当第n-1行的皇后也放置完成
	if len(oneRes) == n {
		temp := make([]string, len(oneRes))
		copy(temp, oneRes)
		result = append(result, temp)
		return
	}

	var tempStr string
	for i := 0; i < n; i++ {
		tempStr += "."
	}

	//在第row行可以放在每一列
	for idx := 0; idx < n; idx++ {
		judgeRes := judgeValid(oneRes, col, idx)
		if judgeRes {
			//判断在idx这个位置能不能放
			tempSlice := []byte(tempStr)
			tempSlice[idx] = 'Q'
			oneRes = append(oneRes, string(tempSlice))
			putQueen(col+1, n)
			oneRes = oneRes[:len(oneRes)-1]
		}
	}
}

//判断在某个点能否放置皇后
func judgeValid(now []string, col, row int) bool {
	var resultJudge [][]int
	//将now中的Q位置坐标提取出来，放入一个二维数组
	for idx := 0; idx < len(now); idx++ {
		for j := 0; j < len(now[0]); j++ {
			if string(now[idx][j]) == "Q" {
				var temp []int
				temp = append(temp, idx)
				temp = append(temp, j)
				resultJudge = append(resultJudge, temp)
			}
		}
	}

	for _, res := range resultJudge {
		if res[0] == col || res[1] == row {
			return false
		}
		if res[0]+res[1] == col+row || res[0]-res[1] == col-row {
			return false
		}
	}

	return true

}
