package week2

import "fmt"

/*
   双链表+map,链表元素的val是对应的key和value，map存每个key对应的链表节点
   R ---> recently
*/

//先定义一个双链表的节点，一开始没有把链表的struct放入key，导致后续删除尾节点时找不到map中对应的key
type NodeList struct {
	key  int
	val  int
	pre  *NodeList
	next *NodeList
}

type LRUCache struct {
	//包含链表头尾相互指向的保护节点
	head     *NodeList
	tail     *NodeList
	mapInfo  map[int]*NodeList
	capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     &NodeList{},
		tail:     &NodeList{},
		mapInfo:  make(map[int]*NodeList, capacity),
	}

	lru.tail.next = lru.head
	lru.head.next = lru.tail

	return lru
}

func (this *LRUCache) Get(key int) int {
	//如果匹配成功的话
	if res, ok := this.mapInfo[key]; ok {
		//将对应的链表节点插入到表头，同时删除源节点
		// fmt.Printf("get,key:%v,delete：%v\n",key,res.val)
		this.deleteFromList(res)
		//这也要判断有没有超capa
		this.insertToHead(res)
		//返回查找的结果
		return res.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// this.printList()
	newNode := &NodeList{key, value, nil, nil}
	//如果该key已经存在
	if res, ok := this.mapInfo[key]; ok {
		//将对应的链表节点插入到表头，同时删除源节点
		this.deleteFromList(res)
		this.insertToHead(newNode)
		//更新map,golang不能对map取值的value取地址
		// this.mapInfo[key] = newNode
	} else {
		//判断是否超出capacity
		if len(this.mapInfo) >= this.capacity {
			//删除链表尾部节点和map数组
			// fmt.Printf("beyond capacity,key:%v,delete：%v\n",key,this.tail.pre.val)
			this.deleteFromList(this.tail.pre)
			// delete(this.mapInfo,this.tail.pre.val)
		}
		// this.mapInfo[key] = newNode
		this.insertToHead(newNode)
	}
}

func (this *LRUCache) insertToHead(res *NodeList) {
	// this.printList()
	res.next = this.head.next
	this.head.next.pre = res
	//完成head和res的指向
	res.pre = this.head
	this.head.next = res

	this.mapInfo[res.key] = res
	// this.printList()
	// for k,_ := range this.mapInfo{
	//     fmt.Printf("k:%v",k)
	// }

	// fmt.Print("\n")
}

func (this *LRUCache) deleteFromList(res *NodeList) {
	res.pre.next = res.next
	res.next.pre = res.pre
	delete(this.mapInfo, res.key)
}

func (this *LRUCache) printList() {
	fmt.Print("start\n")
	tmp := this.head
	for tmp.next != this.tail {
		fmt.Printf("%v-", tmp.next.val)
		tmp = tmp.next
	}
	fmt.Print("\n")
	fmt.Print("end\n")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
