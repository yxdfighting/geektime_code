package week1

type NumMatrix struct {
	//定义为一个二维数组，sumMatrix[i][j]表示从给定二维矩阵左上角开始,到（i，j）的矩阵元素之和
	sumMatrix [][]int
}

//此处有一点注意，因为第一行和第一列递推式也成立，只要i或者j小于0时返回0，式子就成立
func Constructor(matrix [][]int) NumMatrix {
	//初始化sumMatrix
	var sumMatrix [][]int
	for i := 0; i < len(matrix); i++ {
		tmp := make([]int, len(matrix[0]))
		sumMatrix = append(sumMatrix, tmp)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			sumMatrix[i][j] = get_sum(sumMatrix, i-1, j) + get_sum(sumMatrix, i, j-1) + matrix[i][j] - get_sum(sumMatrix, i-1, j-1)
		}
	}
	numMatrix := NumMatrix{sumMatrix}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sumMatrix[row2][col2] - get_sum(this.sumMatrix, row2, col1-1) - get_sum(this.sumMatrix, row1-1, col2) + get_sum(this.sumMatrix, row1-1, col1-1)
}

func get_sum(matrix [][]int, i, j int) int {
	if i >= 0 && j >= 0 {
		return matrix[i][j]
	}
	return 0
}
