package two_pointer

type LinkedList struct {
	data int
	next *LinkedList
}

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/
func removeNthNode(head *LinkedList, n int) *LinkedList {

	if head == nil {
		return nil
	}

	if n <= 0 {
		return nil
	}

	dummy := head
	fast := head

	for fast != nil && n > 0 {
		fast = fast.next
		n--
	}
	slow := head
	for fast != nil && fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next

	return dummy
}
