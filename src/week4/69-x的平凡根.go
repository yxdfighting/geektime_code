package week4

func mySqrt(x int) int {
	l, r := 1, x

	//平方小于等于x的最大值
	for l < r {
		mid := (l + r + 1) / 2
		//条件
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return r
}
