package week3

/*
   状态空间
   每个基因序列可以看作图的一个点 每个基因序列有8个字符，所以每个基因序列有3的8次方条出边

   采用bfs，从start出发，在bank中找到符合条件的出边
*/

type StringLevel struct {
	gean  string
	level int
}

func minMutation(start string, end string, bank []string) int {
	//key表示基因序列 value表示层数
	var queue []StringLevel
	//遍历start的出边
	strLevel := StringLevel{start, 0}
	queue = append(queue, strLevel)

	if !existIn(end, bank) {
		return -1
	}

	for len(queue) != 0 {
		// fmt.Printf("queue:%v\n",queue)
		//取对头
		head := queue[0]
		//出队
		queue = queue[1:]
		//遍历队首的出边
		for i := 0; i < 8; i++ {
			for _, str := range string("ACGT") {
				// fmt.Printf("head.gean[i]:%v,str:%v\n",string(head.gean[i]),string(str))
				if string(str) != string(head.gean[i]) {
					byteSlice := []rune(head.gean)
					byteSlice[i] = str

					if !existIn(string(byteSlice), bank) {
						continue
					}

					// fmt.Printf("byteSlice:%v\n",string(byteSlice))
					if string(byteSlice) == end {
						return head.level + 1
					}

					strLev := StringLevel{string(byteSlice), head.level + 1}
					queue = append(queue, strLev)
				}
			}
		}
	}
	return -1
}

func existIn(str string, strSlice []string) bool {
	for _, strIn := range strSlice {
		if str == strIn {
			return true
		}
	}
	return false
}
