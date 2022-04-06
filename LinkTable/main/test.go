package main

import (
	"fmt"
)

const (
	SUCCESS = 0
	FAILURE = -1
)

type LinkTableNode interface {
	Next() *LinkTableNode
}

type LinkTable interface {
	GetHead() *LinkTableNode
	GetTail() *LinkTableNode
	GetLength() int
}
type tLinkTableNode struct {
	pNext *tLinkTableNode
	value int
}

type tLinkTable struct {
	pHead     *tLinkTableNode
	pTail     *tLinkTableNode
	SumOfNode int
}

func (pLinkTableNode *tLinkTableNode) Next() *tLinkTableNode {
	return pLinkTableNode.pNext
}

func (pLinkTable *tLinkTable) GetHead() *tLinkTableNode {
	return pLinkTable.pHead
}

func (pLinkTable *tLinkTable) GetTail() *tLinkTableNode {
	return pLinkTable.pTail
}

func (pLinkTable *tLinkTable) GetLength() int {
	return pLinkTable.SumOfNode
}

//创建链表
func CreateLinkTable() *tLinkTable {
	var pLinkTable *tLinkTable = new(tLinkTable)
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return pLinkTable
}

//删除链表
func DeleteLinkTable(pLinkTable *tLinkTable) int {
	if pLinkTable == nil {
		return FAILURE
	}
	pLinkTable = nil
	return SUCCESS
}

//获取节点值
func GetValue(pLinkTableNode *tLinkTableNode) int {
	return pLinkTableNode.value
}

//添加节点
func AddLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) int {
	if (pLinkTable == nil) || (pNode == nil) {
		return FAILURE
	}
	pNode.pNext = nil
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
	}
	if pLinkTable.pTail == nil {
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.SumOfNode += 1
	return SUCCESS
}

//删除目标节点
func DeleteLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) int {
	if (pLinkTable == nil) || (pNode == nil) || pLinkTable.SumOfNode == 0 {
		return FAILURE
	}

	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.pNext
		pLinkTable.SumOfNode -= 1
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return SUCCESS
	}
	pTempNode := pLinkTable.pHead
	// length := pLinkTable.GetLength()
	for nil != pTempNode {
		if pTempNode.pNext == pNode {
			pTempNode.pNext = pTempNode.pNext.pNext
			pLinkTable.SumOfNode -= 1
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			return SUCCESS
		} else {
			pTempNode = pTempNode.pNext
		}
	}
	return FAILURE
}

//获取目标节点的下一个节点
func GetNextLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) *tLinkTableNode {
	if (pLinkTable == nil) || (pNode == nil) {
		return nil
	}
	pTempNode := pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.pNext
		}
		pTempNode = pTempNode.pNext
	}
	return nil
}

//输出链表
func PrintLinkTable(pLinkTable *tLinkTable) int {
	if pLinkTable == nil {
		fmt.Println("链表为空")
		return FAILURE
	}
	p := pLinkTable.pHead
	// length := pLinkTable.GetLength()
	for nil != p {
		fmt.Printf("%d ", p.value)
		p = p.Next()
	}
	fmt.Println()
	return SUCCESS
}

func main() {

	pLinkTable := new(tLinkTable)
	//添加节点
	for i := 0; i < 10; i++ {
		var pNode *tLinkTableNode = new(tLinkTableNode)
		pNode.value = i
		AddLinkTableNode(pLinkTable, pNode)
	}
	//使用继承自接口的GetLength和GetHead等方法
	fmt.Println("链表长度为：", pLinkTable.GetLength())
	//输出链表
	PrintLinkTable(pLinkTable)
	//删除偶数节点
	pTempNode := pLinkTable.GetHead()
	for nil != pTempNode {
		if pTempNode.value%2 == 0 {
			q := pTempNode
			pTempNode = pTempNode.Next()
			DeleteLinkTableNode(pLinkTable, q)
		} else {
			pTempNode = pTempNode.Next()
		}

	}
	//输出删除后的链表
	PrintLinkTable(pLinkTable)
	//输出第三个节点的值
	pTempNode = pLinkTable.pHead.Next().Next()
	fmt.Println("该节点的值为:", GetValue(pTempNode))
	fmt.Println("下一个节点的值为:", GetValue(GetNextLinkTableNode(pLinkTable, pTempNode)))
}
