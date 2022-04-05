package main

import "fmt"

const (
	SUCCESS = 0
	FAILURE = -1
)

type LinkTableNode struct {
	next *LinkTableNode
}

type LinkTable struct {
	pHead     *LinkTableNode
	pTail     *LinkTableNode
	SumOfNode int
}

//创建链表
func CreateLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return pLinkTable
}

//删除链表
func DeleteLinkTable(pLinkTable *LinkTable) int {
	if pLinkTable == nil {
		return FAILURE
	}
	pLinkTable = nil
	return SUCCESS

}

//添加节点
func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if (pLinkTable == nil) || (pNode == nil) {
		return FAILURE
	}
	pNode.next = nil
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
	}
	if pLinkTable.pTail == nil {
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.next = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.SumOfNode += 1
	return SUCCESS
}

//删除目标节点
func DeleteLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if (pLinkTable == nil) || (pNode == nil) {
		return FAILURE
	}

	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.next
		pLinkTable.SumOfNode -= 1
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return SUCCESS
	}
	pTempNode := pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode.next == pNode {
			pTempNode.next = pTempNode.next.next
			pLinkTable.SumOfNode -= 1
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			return SUCCESS
		}
	}
	return FAILURE
}

//获取头节点
func GetLinkTableHead(pLinkTable *LinkTable) *LinkTableNode {
	if pLinkTable == nil {
		return nil
	}
	return pLinkTable.pHead
}

//获取目标节点的下一个节点
func GetNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode {
	if (pLinkTable == nil) || (pNode == nil) {
		return nil
	}
	pTempNode := pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.next
		}
		pTempNode = pTempNode.next
	}
	return nil
}
func main() {
	var pLinkTable *LinkTable = CreateLinkTable()
	for i := 0; i < 10; i++ {
		var pNode *LinkTableNode = new(LinkTableNode)
		AddLinkTableNode(pLinkTable, pNode)
	}

	fmt.Println("链表长度为：", pLinkTable.SumOfNode)
	var p2 *LinkTableNode = pLinkTable.pHead
	for i := 5; i > 0; i-- {
		q := p2
		p2 = p2.next
		DeleteLinkTableNode(pLinkTable, q)
	}
	fmt.Println("链表长度为：", pLinkTable.SumOfNode)
}
