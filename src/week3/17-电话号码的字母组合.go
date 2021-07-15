package week3

/*
	如何将这个问题，归纳为状态空间
	数字个数：n     字符串：str
					0,""
       1,a  		1b   		1c
2,ad ae af	 	2,bd be bf   3,cd ce cf

	所以问题本质其实是通过搜索找到长度为n的字符串，这样就可以通过dfs或者bfs进行搜索
*/

var result []string
var res string

func letterCombinations(digits string) []string {
	result = result[0:0]
	res = ""
	if len(digits) == 0 {
		return result
	}
	//定义一个map
	letterMap := make(map[string][]string, len(digits))
	letterMap["2"] = []string{"a", "b", "c"}
	letterMap["3"] = []string{"d", "e", "f"}
	letterMap["4"] = []string{"g", "h", "i"}
	letterMap["5"] = []string{"j", "k", "l"}
	letterMap["6"] = []string{"m", "n", "o"}
	letterMap["7"] = []string{"p", "q", "r", "s"}
	letterMap["8"] = []string{"t", "u", "v"}
	letterMap["9"] = []string{"w", "x", "y", "z"}

	//进行dfs,第idx个位置的选择
	dfs(letterMap, digits, 0, len(digits))
	return result
}

func dfs(letterMap map[string][]string, digits string, idx int, n int) {
	if len(res) == n {
		result = append(result, res)
		return
	}
	num := string(digits[idx])
	for _, letter := range letterMap[num] {
		res = res + letter
		dfs(letterMap, digits, idx+1, n)
		//回溯还原
		res = res[:len(res)-len(letter)]
	}
}
