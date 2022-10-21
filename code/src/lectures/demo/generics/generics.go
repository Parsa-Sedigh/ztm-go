package main

import "fmt"

/* we're gonna create a generic priority queue. A priority queue will allow us to add items to the queue at different priorties and
when we get items out of the queue, it will be the highest priority first.*/

const (
	Low = iota
	Medium
	High
)

// usually we use T for generic type nam,but in this case, we're gonna use P for priority. The type for the value of actual items is V(value)
type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{items: make(map[P][]V), priorities: priorities}
}

/* Since we created a receiver function here, we don't need to specify the generic constraints again because they're already specified on our structure. */
func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

// this function is gonna pull out the next item out of priority queue.
// In this function, the second return value will indicate whether or notsth was actually returned. So true means we have an item and false means there's nothing in the queue
func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		// make a copy of priority
		priority := pq.priorities[i]

		// pull out all of the items in that priority
		items := pq.items[priority]

		// check if there's anything in the queue
		if len(items) > 0 {
			next := items[0]

			// re-slice the priority queue:
			// this line effectively deletes the item that's at position 0. Because we already took it off the queue
			pq.items[priority] = items[1:]

			return next, true
		}
	}

	/* This empty variable of type V, is gonna be initialized the default value of type V. So if V is int, the empty variable will be 0 and ... . */
	var empty V
	return empty, false
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})

	// even though we added a low priority item first, it should be the last thing that is printed out when we do queue.Next()
	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
}
