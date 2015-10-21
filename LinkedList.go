package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}
type LinkedList struct {
	Root   *Node
	Length int
}

// Returns the length of the list
func (l *LinkedList) Len() int { return l.Length }

// Insert inserts node n to the back of the list
// Returns the node that was added
func (l *LinkedList) Insert(n *Node) *Node {

	if l.Length == 0 {
		l.Root = n
		l.Length += 1
		return n
	}

	currentNode := l.Root

	for {
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}

	l.Length += 1
	currentNode.Next = n

	return n
}

// Remove takes out the node from the list where d matches the data in the node
// Returns the int of the node that was removed
func (l *LinkedList) Remove(d int) error {

	if l.Length == 0 {
		return errors.New("List is empty, can't remove value")
	}

	if l.Root.Data == d {
		l.Root = l.Root.Next
		l.Length -= 1
		return nil
	}

	currentNode := l.Root

	for {

		if currentNode.Next == nil {
			return errors.New("Data is not in list")
		}

		if currentNode.Next.Data == d {
			break
		}

		currentNode = currentNode.Next
	}

	currentNode.Next = currentNode.Next.Next
	l.Length -= 1

	return nil
}

func (l *LinkedList) Print() error {

	if l.Length == 0 {
		fmt.Println("List is empty")
		return errors.New("Can't print, list is empty")
	}

	currentNode := l.Root

	for {
		fmt.Print(currentNode.Data, ", ")
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next

	}

	fmt.Println()
	return nil
}

func main() {

	list := LinkedList{}

	firstNode := Node{1, nil}
	secondNode := Node{2, nil}
	thirdNode := Node{3, nil}
	forthNode := Node{4, nil}

	list2 := LinkedList{Root: &firstNode, Length: 4}
	fmt.Println(list2)

	var x int 
	x = 0 

	if x == 0 {
		fmt.Println("hello there")
	}

	list.Insert(&firstNode)
	list.Print()

	list.Insert(&secondNode)
	list.Print()

	list.Insert(&thirdNode)
	list.Print()

	list.Insert(&forthNode)
	list.Print()
	fmt.Println("Length after 4 adds: ", list.Length)

	err := list.Remove(1)
	if err != nil {
		fmt.Println(err)
	}

	list.Print()
	err = list.Remove(10)
	if err != nil {
		fmt.Println(err)
	}

	list.Print()
	err = list.Remove(2)
	if err != nil {
		fmt.Println(err)
	}
	list.Print()

	err = list.Remove(3)
	if err != nil {
		fmt.Println(err)
	}

	list.Print()
	err = list.Remove(4)
	if err != nil {
		fmt.Println(err)
	}
	list.Print()
	fmt.Println(list.Length)

}
