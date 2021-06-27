package week2

import (
	"strconv"
	"strings"
)

/*
   首先遍历每一个数组，在一个数组中，对第二部分从后往前，
遇到 . 表示遇到了一个层级的域名，放入map 同时count 第一部分的数字
*/

func subdomainVisits(cpdomains []string) []string {
	ansMap := make(map[string]int)
	var ans []string
	for _, domain := range cpdomains {
		//先按照空格拆分
		cpDomain := strings.Split(domain, " ")

		domainList := strings.Split(cpDomain[1], ".")
		temp := ""
		for i := len(domainList) - 1; i >= 0; i-- {
			temp = domainList[i] + "." + temp
			temp = strings.TrimSuffix(temp, ".")
			intRes, _ := strconv.Atoi(cpDomain[0])
			ansMap[temp] += intRes
		}
	}

	// fmt.Printf("map:%v",ansMap)

	for k, v := range ansMap {
		str := strconv.Itoa(v) + " " + k
		ans = append(ans, str)
	}

	return ans
}
