package main

import (
	"fmt"
)

type elem struct {
	name string
	next *elem
}

type singleList struct {
	len  int
	head *elem
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddBack(name string) *elem{
	elem := &elem{
		name: name,
	}
	if s.head == nil {
		s.head = elem
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = elem
	}
	s.len++
	return elem
}

func (s *singleList) showList() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func main() {
	singleList := initList()
	s := Act(singleList)

	//Вывести весь список.
	/*err := singleList.showList()
	if err != nil {
		fmt.Println(err.Error())
	}*/
	fmt.Println(s)
}

func Act(singleList *singleList) string{

	m := make(map[*elem]interface{})
	words := fill()
	for i := range words {
		singleList.AddBack(words[i])
	}

	current := singleList.head

	for current.next != nil {
		current = current.next
	}
	current.next = singleList.head.next.next


	for current.next != nil {
		for range m {
			if _, ok := m[current]; ok {
				loopPoint := current.next
				fmt.Println("Loop found", current.next.name)
				return loopPoint.name
			}
		}
		m[current] = nil
		current = current.next
	}

	// fmt.Println("Loop starting point is", fast.name)
	return "Loop not found."
}

func fill() [6]string{
	words := [6]string{"A", "B", "C", "D", "E"}
	return words
}
