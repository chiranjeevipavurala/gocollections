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
		val := iterator.Next()
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
		fmt.Println("From List Iterator", listIterator.Next())
	}

}
