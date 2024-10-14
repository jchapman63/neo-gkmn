package pkg

// a priority queue
// using a priority queue
// will allow me to dynamically
// decide which pokemon goes first
// in a battle should the trainer use
// an item on a turn or use a quicker move
// etc, etc, etc

// using the go example implementation, we have to manually
// push an item and them update the item to
// properly order the queue (heap) by priority

// This example demonstrates a priority queue built using the heap interface.
import (
	"container/heap"
	"fmt"

	"github.com/google/uuid"
)

var speed_limit int = 999

// An Item is something we manage in a priority queue.
type Item struct {
	monID    uuid.UUID // The value of the item; arbitrary.
	speed    int       // used to determine priority
	priority int       // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n

	*pq = append(*pq, item)

	pq.registerPriority(item)
	if len(*pq) > 1 {
		fmt.Println("fixing heap")
		heap.Fix(pq, item.index)
	}
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Determine the priority of the item
// based on the associated monster's speed (for now)
func (pq *PriorityQueue) registerPriority(item *Item) {
	// faster items have higher priority
	item.priority = speed_limit - item.speed
}
