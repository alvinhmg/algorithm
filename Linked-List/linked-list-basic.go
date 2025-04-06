package ALGORITHM

/**
@author: alvin
@create: 2025-03-5 09:00
@title: 链表的常用操作
@leetcode:
*/

// NewListNode 创建一个新的链表节点
func NewListNode(Val int) *ListNode {
	return &ListNode{
		Val:  Val,
		Next: nil,
	}
}

// 插入节点
// 头插法
func InsertNode(head *ListNode, Val int) *ListNode {
	if head == nil {
		return NewListNode(Val)
	}
	newNode := NewListNode(Val)
	newNode.Next = head
	head = newNode
	return head
}

// 尾插法
func InsertNodeTail(head *ListNode, Val int) *ListNode {
	if head == nil {
		return NewListNode(Val)
	}
	newNode := NewListNode(Val)
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

// 删除节点（单节点）
func DeleteNode(head *ListNode, Val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == Val { // 处理头结点就是要删除的节点的情况
		return head.Next
	}

	current := head
	for current.Next != nil { // 遍历链表查找要删除的节点
		if current.Next.Val == Val {
			current.Next = current.Next.Next
			return head // 这里是删除一个节点就返回，如果要删除所有节点，这里就不能返回
		}
		current = current.Next
	}
	return head
}

// 删除所有指定的节点
func DeleteAllNodes(head *ListNode, Val int) *ListNode {
	dummy := ListNode{Next: head} // 创建一个虚拟头结点，方便处理头结点就是要删除的节点的情况
	current := &dummy
	for current.Next != nil {
		if current.Next.Val == Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}

// 反转链表操作

// 检测环形链表
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 获取环形链表的入口节点
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 找到环后，将慢指针重置到头节点
			slow = head
			// 两个指针以相同速度移动，相遇点即为环的入口
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	// 没有环的情况
	return nil
}

// 获取链表的中间节点
func GetMiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 合并两个有序链表 -- 递归法
func MergeTowList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTowList(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTowList(l1, l2.Next)
		return l2
	}
}

// 合并两个有序链表 -- 迭代法
func MergeTowList2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil { // l1 和 l2 都不为空
		if l1.Val < l2.Val { // 将较小的节点添加到新链表中
			curr.Next = l1
			l1 = l1.Next
		} else { // 将较小的节点添加到新链表中
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next // 更新新链表的尾节点
	}
	// 将剩余的节点添加到新链表中
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// 删除链表倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}


// 链表相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	return dummy.Next
}
