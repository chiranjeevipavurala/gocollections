package main

import (
	"testing"

	"github.com/chiranjeevipavurala/gocollections/lists"
	"github.com/chiranjeevipavurala/gocollections/queues"
)

// PriorityQueue Benchmarks

func BenchmarkPriorityQueueAdd(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		priorityQueue.Add(i)
	}
}

func BenchmarkPriorityQueuePoll(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if priorityQueue.Size() > 0 {
			priorityQueue.Poll()
		}
	}
}

func BenchmarkPriorityQueuePeek(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Peek()
	}
}

func BenchmarkPriorityQueueRemove(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Remove(i % MediumSize)
	}
}

func BenchmarkPriorityQueueContains(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Contains(i % MediumSize)
	}
}

// Stack Benchmarks

func BenchmarkStackPush(b *testing.B) {
	stack := lists.NewStack[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := lists.NewStack[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if stack.Size() > 0 {
			stack.Pop()
		}
	}
}

func BenchmarkStackPeek(b *testing.B) {
	stack := lists.NewStack[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

// LinkedList as Queue Benchmarks

func BenchmarkLinkedListAsQueueOffer(b *testing.B) {
	queue := lists.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		queue.Offer(i)
	}
}

func BenchmarkLinkedListAsQueuePoll(b *testing.B) {
	queue := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		queue.Offer(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if queue.Size() > 0 {
			queue.Poll()
		}
	}
}

func BenchmarkLinkedListAsQueuePeek(b *testing.B) {
	queue := lists.NewLinkedList[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		queue.Offer(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Peek()
	}
}

// IntComparator for PriorityQueue
type IntComparator struct{}

func (c *IntComparator) Compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Additional PriorityQueue benchmarks for missing methods

func BenchmarkPriorityQueueSize(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Size()
	}
}

func BenchmarkPriorityQueueIsEmpty(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.IsEmpty()
	}
}

func BenchmarkPriorityQueueClear(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Clear()
		// Re-populate for next iteration
		for j := 0; j < MediumSize; j++ {
			priorityQueue.Add(j)
		}
	}
}

func BenchmarkPriorityQueueToArray(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.ToArray()
	}
}

func BenchmarkPriorityQueueIterator(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := priorityQueue.Iterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkPriorityQueueAddAll(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	otherQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate other queue
	for i := 0; i < SmallSize; i++ {
		otherQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.AddAll(otherQueue)
	}
}

func BenchmarkPriorityQueueRemoveAll(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	otherQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate both queues
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
		if i%2 == 0 {
			otherQueue.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.RemoveAll(otherQueue)
	}
}

func BenchmarkPriorityQueueContainsAll(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	otherQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate both queues
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
		if i%2 == 0 {
			otherQueue.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.ContainsAll(otherQueue)
	}
}

func BenchmarkPriorityQueueEquals(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	otherQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate both queues identically
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
		otherQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Equals(otherQueue)
	}
}

func BenchmarkPriorityQueueRemoveIf(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	// Type assert to access RemoveIf
	if pq, ok := priorityQueue.(*queues.PriorityQueue[int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pq.RemoveIf(func(element int) bool {
				return element%2 == 0
			})
		}
	}
}

func BenchmarkPriorityQueueForEach(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	// Type assert to access ForEach
	if pq, ok := priorityQueue.(*queues.PriorityQueue[int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pq.ForEach(func(element int) {
				// Do nothing, just iterate
			})
		}
	}
}

func BenchmarkPriorityQueueGetComparator(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)

	// Type assert to access GetComparator
	if pq, ok := priorityQueue.(*queues.PriorityQueue[int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			pq.GetComparator()
		}
	}
}

func BenchmarkPriorityQueueElement(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priorityQueue.Element()
	}
}

func BenchmarkPriorityQueueOffer(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		priorityQueue.Offer(i)
	}
}

func BenchmarkPriorityQueueRemoveHead(b *testing.B) {
	comparator := &IntComparator{}
	priorityQueue := queues.NewPriorityQueue[int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		priorityQueue.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if priorityQueue.Size() > 0 {
			priorityQueue.RemoveHead()
		}
	}
}
