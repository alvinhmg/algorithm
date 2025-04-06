package ALGORITHM

/**
@author: alvin
@create: 2025-02-10 22:30
@title: 反转链表
@leetcode: https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
@题解：
	1、使用双指针
	2、使用递归
	3、使用头插法
	4、使用栈
*/

// ListNode 定义链表节点结构
type ListNode struct {
    Val  int
    Next *ListNode
}

type DoublyLinkedList struct {
    Val int
    Prev *ListNode
    Next *ListNode
}


// ReverseList1 迭代法（双指针法） --推荐
func ReverseList1(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next // 保存下一个节点
        curr.Next = prev // 反转指针
        prev = curr // 更新指针
        curr = next // 更新指针
    }
    return prev
}

// ReverseList2 递归法
func ReverseList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := ReverseList2(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

// ReverseList3 头插法
func ReverseList3(head *ListNode) *ListNode {
    dummy := &ListNode{} // 虚拟头节点
    curr := head // 当前节点
    for curr != nil { 
        next := curr.Next // 保存下一个节点
        curr.Next = dummy.Next // 头插法
        dummy.Next = curr // 头插法
        curr = next // 更新指针
    }
    return dummy.Next // 返回头节点
}

// ReverseList4 使用栈
func ReverseList4(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    // 将节点入栈
    stack := []*ListNode{}
    curr := head
    for curr != nil {
        stack = append(stack, curr)
        curr = curr.Next
    }
    
    // 重建链表
    dummy := &ListNode{}
    curr = dummy
    for i := len(stack) - 1; i >= 0; i-- {
        curr.Next = stack[i]
        curr = curr.Next
    }
    curr.Next = nil
    return dummy.Next
}
