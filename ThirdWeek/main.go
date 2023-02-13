package ThirdWeek

import (
	"GoPackages/FirstWeek"
	"fmt"
)

type StudentStack struct {
	top *Node
}

type Node struct {
	student *FirstWeek.Student
	next    *Node
}

func (s *StudentStack) Push(st *FirstWeek.Student) {
	node := &Node{student: st, next: s.top}
	s.top = node
}

func (s *StudentStack) Pop() *FirstWeek.Student {
	if s.top == nil {
		return nil
	}

	st := s.top.student
	s.top = s.top.next
	return st
}

func (s *StudentStack) Peek() *FirstWeek.Student {
	if s.top == nil {
		return nil
	}

	return s.top.student
}

func (s *StudentStack) Clear() {
	s.top = nil
}

func (s *StudentStack) Contains(st *FirstWeek.Student) bool {
	node := s.top
	for node != nil {
		if node.student == st {
			return true
		}
		node = node.next
	}
	return false
}

func (s *StudentStack) Increment() {
	node := s.top
	for node != nil {
		node.student.Age++
		node = node.next
	}
}

func (s *StudentStack) Print() {
	node := s.top
	for node != nil {
		fmt.Println(node.student)
		node = node.next
	}
}

func (s *StudentStack) PrintReverse() {
	stack := &StudentStack{}
	node := s.top
	for node != nil {
		stack.Push(node.student)
		node = node.next
	}

	node = stack.top
	for node != nil {
		fmt.Println(node.student)
		node = node.next
	}
}
