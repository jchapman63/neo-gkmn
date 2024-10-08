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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, monID uuid.UUID, priority int) {
	item.monID = monID
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[uuid.UUID]int{
		uuid.New(): 3, uuid.New(): 2, uuid.New(): 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for id, priority := range items {
		pq[i] = &Item{
			monID:    id,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		monID:    uuid.New(),
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.monID, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.monID)
	}
}
