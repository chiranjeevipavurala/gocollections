package main

import (
	"fmt"

	"github.com/chiranjeevipavurala/gocollections/lists"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	hs := sets.NewHashSet[string]()
	hs.Add("xyz")
	hs.Add("abc")
	hs.Add("def")
	hs.Add("ghi")
	hs.Add("jkl")

	fmt.Println("size for hashset", hs.Size())
	if hs.Contains("xyz1") {
		fmt.Println("element exists")
	} else {
		fmt.Println("Element doesnt exist")
	}

	iterator := hs.Iterator()
	for iterator.HasNext() {
		val, _ := iterator.Next()
		fmt.Println("From Iterator", val)
	}

	tempList := lists.NewArrayList[string]()
	tempList.Add("list1")
	tempList.Add("list2")
	tempList.Add("list3")
	_ = tempList.AddAtIndex(0, "list0")
	_ = tempList.AddAtIndex(4, "list4")
	_ = tempList.AddAtIndex(2, "listPre2")
	fmt.Println("length of list is", tempList.Size())

	_, _ = tempList.RemoveAtIndex(0)
	_, _ = tempList.RemoveAtIndex(4)
	_, _ = tempList.RemoveAtIndex(1)
	_, _ = tempList.Set(2, "updated1")
	listIterator := tempList.Iterator()
	for listIterator.HasNext() {
		val, _ := listIterator.Next()
		fmt.Println("From List Iterator", val)
	}

	arrayList := lists.NewArrayListWithInitialCollection[string]([]string{"c", "d", "e", "a", "b"})
	for _, val := range arrayList.ToArray() {
		fmt.Println("From Array List", val)
	}
	arrayList.Sort(&StringComparator{})

	for _, val := range arrayList.ToArray() {
		fmt.Println("From Array List", val)
	}

	stack := lists.NewStack[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	fmt.Println("Size of Stack is", stack.Size())

	top, _ := stack.Pop()
	top, _ = stack.Pop()
	fmt.Println("Top of stack", *top)

	queue := lists.NewLinkedList[int]()
	queue.Add(1)
	queue.Add(3)
	queue.Add(2)

	fmt.Println("Queue", queue.Size())
	queue.RemoveFirst()
	fmt.Println("Queue", queue.Size())
	val, _ := queue.Get(0)
	fmt.Println("Queue at top of queue", *val)

}

type StringComparator struct {
}

func (c *StringComparator) Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
