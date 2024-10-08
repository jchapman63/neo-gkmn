package pkg

import (
	"container/heap"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMonsterPriority(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
	}{
		{
			name: "fastest of two monsters",
			items: []Item{
				{monID: uuid.New(), speed: 80, priority: 0},
				{monID: uuid.New(), speed: 55, priority: 0},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			pq := make(PriorityQueue, len(tt.items))
			for i, item := range tt.items {
				pq[i] = &item
			}
			heap.Init(&pq)

			first := pq.Pop()
			firstSpeed := first.(*Item).speed
			second := pq.Pop()
			secondSpeed := second.(*Item).speed
			assert.Greater(t, firstSpeed, secondSpeed)
		})
	}
}
