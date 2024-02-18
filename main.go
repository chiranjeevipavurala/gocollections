package main

import (
	"fmt"
	"gocollections/sets"
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

}
