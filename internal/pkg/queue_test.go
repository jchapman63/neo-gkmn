package pkg

import (
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
				{monID: uuid.New(), speed: 50, priority: 0},
				{monID: uuid.New(), speed: 55, priority: 0},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// TODO : finish priority queue
			// pq := make(PriorityQueue, len(tt.items))
			actual := tt.name
			assert.Equal(t, actual, tt.name)
		})
	}
}
