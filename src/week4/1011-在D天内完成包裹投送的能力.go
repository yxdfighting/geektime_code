package week4

/*
	对于weight，能够在d天内完成的最小载重量
	本质上是个求解问题 我们可以先假设一个载重量t 看能否满足
	然后在一个范围内通过单调性依次判断每个t
*/
func shipWithinDays(weights []int, days int) int {
	//先明确要找的元素上下边界,假定最少天数为1天，那么载重量最大为所有元素之和
	//假定需要len(weights)天，则需要的载重量为元素里面的最大值
	var max, min int
	for i := 0; i < len(weights); i++ {
		max += weights[i]
		if weights[i] > min {
			min = weights[i]
		}
	}
	//for i := min;i< max;i++{
	//	fmt.Printf("载重量为%v时，满足结果：%v\n",i,canShip(weights,days,i))
	//}

	//二分查找，在小于需要的最小载重量时，canShip为false，大于等于时为true  满足单调性
	l, r := min, max
	for l < r {
		mid := (l + r) / 2
		//条件
		if canShip(weights, days, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

//判断载重为t时能否满足条件
func canShip(weights []int, days, t int) bool {
	sum := 0
	dayTimes := 1
	//先贪心判断weight
	for i := 0; i < len(weights); i++ {
		//fmt.Printf("i:%v,sum:%v,dayTimes:%v\n",i,sum,dayTimes)
		if sum+weights[i] > t {
			sum = weights[i]
			dayTimes++
		} else {
			sum = sum + weights[i]
		}
	}
	//fmt.Printf("t:%v,dayTimes:%v\n",t,dayTimes)
	return dayTimes <= days
}
