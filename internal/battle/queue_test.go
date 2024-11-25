package battle

import (
	"container/heap"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMonsterPriority(t *testing.T) {
	tests := []struct {
		name  string
		items []*Item
	}{
		{
			name: "fastest of two monsters",
			items: []*Item{
				{monID: uuid.New().String(), speed: 80, priority: 0},
				{monID: uuid.New().String(), speed: 55, priority: 0},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			pq := make(PriorityQueue, 0)
			heap.Init(&pq)
			for _, item := range tt.items {
				pq.Push(item)
			}
			first := pq.Pop()
			firstSpeed := first.(*Item).speed

			second := pq.Pop()
			secondSpeed := second.(*Item).speed

			assert.Greater(t, firstSpeed, secondSpeed)
		})
	}
}
