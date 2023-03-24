package two_pointer

import (
	"testing"
)

func TestRemoveNthNode(t *testing.T) {

	head := &LinkedList{data: 1}
	two := &LinkedList{data: 2}
	head.next = two
	three := &LinkedList{data: 3}
	two.next = three
	four := &LinkedList{data: 4}
	three.next = four
	five := &LinkedList{data: 5}
	four.next = five

	expected := []int{1, 2, 3, 5}
	head = removeNthNode(head, 2)

	idx := 0
	isPassed := true
	for head != nil {
		if expected[idx] != head.data {
			isPassed = false
			break
		}
		idx++
		head = head.next
	}

	if !isPassed {
		t.Errorf(" test case is not passed")
	}
}

func TestRemoveNthNode_ZeroNode(t *testing.T) {
	head := &LinkedList{data: 1}
	two := &LinkedList{data: 2}
	head.next = two

	head = removeNthNode(head, 0)

	if head != nil {
		t.Errorf(" test case is not passed")
	}
}

func TestRemoveNthNode_NegativeNode(t *testing.T) {
	head := &LinkedList{data: 1}
	two := &LinkedList{data: 2}
	head.next = two

	head = removeNthNode(head, -1)

	if head != nil {
		t.Errorf(" test case is not passed")
	}
}

func TestRemoveNthNode_NilNode(t *testing.T) {

	head := removeNthNode(nil, 0)

	if head != nil {
		t.Errorf(" test case is not passed")
	}
}
