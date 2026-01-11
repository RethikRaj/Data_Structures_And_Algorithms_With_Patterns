package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

// TC: O(1), SC: O(1)
func (s *SinglyLinkedList) InsertAtHead(val int) {
	newNode := &Node{val: val, next: nil}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

// TC: O(1) because we have tail pointer , SC: O(1)
func (s *SinglyLinkedList) InsertAtTail(val int) {
	newNode := &Node{val: val, next: nil}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
}

// TC: O(1), SC: O(1)
func (s *SinglyLinkedList) DeleteAtHead() (int, error) {
	// if s.head == nil { // if LL is empty
	// 	return 0, fmt.Errorf("LL is empty")
	// } else if s.head == s.tail { // LL has one node
	// 	val := s.head.val
	// 	s.head = nil
	// 	s.tail = nil
	// 	return val, nil
	// } else {
	// 	// Go by default analyzes that the old head node isn't in use(unreachable) and garbage collects it.
	// 	val := s.head.val
	// 	s.head = s.head.next
	// 	return val, nil
	// }

	// 2nd way
	if s.head == nil {
		return 0, fmt.Errorf("LL is empty")
	}

	// Capture the value to return
	val := s.head.val

	// Move the head forward
	s.head = s.head.next

	// If the list is now empty, then that means LL was having single node => sync tail
	if s.head == nil {
		s.tail = nil
	}

	return val, nil
}

// TC: O(n), SC: O(1)
func (s *SinglyLinkedList) DeleteAtTail() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("LL is empty")
	} else if s.head.next == nil { // LL has 1 node
		val := s.head.val
		s.head = nil
		s.tail = nil
		return val, nil
	}

	ptr := s.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	// now ptr points to 2nd last node
	val := ptr.next.val
	ptr.next = nil
	s.tail = ptr
	return val, nil
}

// TC: O(n), SC: O(1)
func (s *SinglyLinkedList) SearchLL(key int) *Node {
	ptr := s.head

	for ptr != nil {
		if ptr.val == key {
			return ptr
		}
		ptr = ptr.next
	}
	return nil
}

// TC: O(n), SC: O(1)
func (s *SinglyLinkedList) printLL() {
	if s.head == nil {
		fmt.Println("List is empty")
		return
	}

	curr := s.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}

func main() {
	sll := &SinglyLinkedList{}

	// Testing InsertAtHead and InsertAtTail
	fmt.Println("--- Inserting Nodes ---")
	sll.InsertAtHead(10) // [10]
	sll.InsertAtHead(20) // [20, 10]
	sll.InsertAtTail(30) // [20, 10, 30]
	sll.printLL()

	// Testing SearchLL
	fmt.Println("\n--- Searching Nodes ---")
	found := sll.SearchLL(10)
	if found != nil {
		fmt.Printf("Found value: %d at address %p\n", found.val, found)
	} else {
		fmt.Println("Value not found")
	}

	// Testing DeleteAtHead
	fmt.Println("\n--- Deleting Head ---")
	valH, errH := sll.DeleteAtHead()
	if errH == nil {
		fmt.Printf("Deleted Head: %d\n", valH)
	}
	sll.printLL()

	// Testing DeleteAtTail
	fmt.Println("\n--- Deleting Tail ---")
	valT, errT := sll.DeleteAtTail()
	if errT == nil {
		fmt.Printf("Deleted Tail: %d\n", valT)
	}
	sll.printLL()

	// Final state
	fmt.Print("\nFinal List: ")
	sll.printLL()
}
