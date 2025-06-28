package main

import (
	"testing"

	"github.com/chiranjeevipavurala/gocollections/lists"
)

// ArrayList Benchmarks

func BenchmarkArrayListAdd(b *testing.B) {
	list := lists.NewArrayList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
}

func BenchmarkArrayListGet(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(i % MediumSize)
	}
}

func BenchmarkArrayListInsertAtIndex(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with some data
	for i := 0; i < SmallSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddAtIndex(i%SmallSize, i)
	}
}

func BenchmarkArrayListRemove(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Remove(i % MediumSize)
	}
}

func BenchmarkArrayListContains(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Contains(i % MediumSize)
	}
}

// Additional ArrayList benchmarks for missing methods

func BenchmarkArrayListAddFirst(b *testing.B) {
	list := lists.NewArrayList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AddFirst(i)
	}
}

func BenchmarkArrayListAddLast(b *testing.B) {
	list := lists.NewArrayList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AddLast(i)
	}
}

func BenchmarkArrayListRemoveFirst(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.RemoveFirst()
		}
	}
}

func BenchmarkArrayListRemoveLast(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.RemoveLast()
		}
	}
}

func BenchmarkArrayListGetFirst(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.GetFirst()
	}
}

func BenchmarkArrayListGetLast(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.GetLast()
	}
}

func BenchmarkArrayListIndexOf(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.IndexOf(i % MediumSize)
	}
}

func BenchmarkArrayListLastIndexOf(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.LastIndexOf(i % MediumSize)
	}
}

func BenchmarkArrayListSet(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Set(i%MediumSize, i)
	}
}

func BenchmarkArrayListAddAll(b *testing.B) {
	list := lists.NewArrayList[int]()
	otherList := lists.NewArrayList[int]()
	// Pre-populate other list
	for i := 0; i < SmallSize; i++ {
		otherList.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddAll(otherList)
	}
}

func BenchmarkArrayListRemoveAll(b *testing.B) {
	list := lists.NewArrayList[int]()
	otherList := lists.NewArrayList[int]()
	// Pre-populate both lists
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
		if i%2 == 0 {
			otherList.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RemoveAll(otherList)
	}
}

func BenchmarkArrayListRetainAll(b *testing.B) {
	list := lists.NewArrayList[int]()
	otherList := lists.NewArrayList[int]()
	// Pre-populate both lists
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
		if i%2 == 0 {
			otherList.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RetainAll(otherList)
	}
}

func BenchmarkArrayListClear(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Clear()
		// Re-populate for next iteration
		for j := 0; j < MediumSize; j++ {
			list.Add(j)
		}
	}
}

func BenchmarkArrayListToArray(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.ToArray()
	}
}

func BenchmarkArrayListIterator(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := list.Iterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkArrayListSort(b *testing.B) {
	list := lists.NewArrayList[int]()
	comparator := &IntComparator{}
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(MediumSize - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Sort(comparator)
	}
}

func BenchmarkArrayListClone(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Clone()
	}
}

func BenchmarkArrayListShuffle(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Shuffle()
	}
}

func BenchmarkArrayListForEach(b *testing.B) {
	list := lists.NewArrayList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.ForEach(func(element int) {
			// Do nothing, just iterate
		})
	}
}

// LinkedList Benchmarks

func BenchmarkLinkedListAdd(b *testing.B) {
	list := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
}

func BenchmarkLinkedListGet(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(i % MediumSize)
	}
}

func BenchmarkLinkedListInsertAtIndex(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with some data
	for i := 0; i < SmallSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddAtIndex(i%SmallSize, i)
	}
}

func BenchmarkLinkedListRemove(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Remove(i % MediumSize)
	}
}

func BenchmarkLinkedListContains(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Contains(i % MediumSize)
	}
}

// LinkedList-specific operations

func BenchmarkLinkedListAddFirst(b *testing.B) {
	list := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AddFirst(i)
	}
}

func BenchmarkLinkedListAddLast(b *testing.B) {
	list := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AddLast(i)
	}
}

func BenchmarkLinkedListRemoveFirst(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.RemoveFirst()
		}
	}
}

func BenchmarkLinkedListRemoveLast(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.RemoveLast()
		}
	}
}

// Additional LinkedList benchmarks for missing methods

func BenchmarkLinkedListGetFirst(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.GetFirst()
	}
}

func BenchmarkLinkedListGetLast(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.GetLast()
	}
}

func BenchmarkLinkedListIndexOf(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.IndexOf(i % MediumSize)
	}
}

func BenchmarkLinkedListLastIndexOf(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.LastIndexOf(i % MediumSize)
	}
}

func BenchmarkLinkedListSet(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Set(i%MediumSize, i)
	}
}

func BenchmarkLinkedListAddAll(b *testing.B) {
	list := lists.NewLinkedList[int]()
	otherList := lists.NewLinkedList[int]()
	// Pre-populate other list
	for i := 0; i < SmallSize; i++ {
		otherList.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddAll(otherList)
	}
}

func BenchmarkLinkedListRemoveAll(b *testing.B) {
	list := lists.NewLinkedList[int]()
	otherList := lists.NewLinkedList[int]()
	// Pre-populate both lists
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
		if i%2 == 0 {
			otherList.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RemoveAll(otherList)
	}
}

func BenchmarkLinkedListClear(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Clear()
		// Re-populate for next iteration
		for j := 0; j < MediumSize; j++ {
			list.Add(j)
		}
	}
}

func BenchmarkLinkedListToArray(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.ToArray()
	}
}

func BenchmarkLinkedListIterator(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := list.Iterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkLinkedListDescendingIterator(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := list.DescendingIterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkLinkedListSort(b *testing.B) {
	list := lists.NewLinkedList[int]()
	comparator := &IntComparator{}
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Add(MediumSize - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Sort(comparator)
	}
}

// Queue operations for LinkedList

func BenchmarkLinkedListOffer(b *testing.B) {
	list := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Offer(i)
	}
}

func BenchmarkLinkedListPoll(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Offer(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.Poll()
		}
	}
}

func BenchmarkLinkedListPeek(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Offer(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Peek()
	}
}

func BenchmarkLinkedListElement(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Offer(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Element()
	}
}

// Stack operations for LinkedList

func BenchmarkLinkedListPush(b *testing.B) {
	list := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Push(i)
	}
}

func BenchmarkLinkedListPop(b *testing.B) {
	list := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		list.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if list.Size() > 0 {
			list.Pop()
		}
	}
}
