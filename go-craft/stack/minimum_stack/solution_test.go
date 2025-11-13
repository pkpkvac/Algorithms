package minimumstack

import (
	"testing"
)

func TestMinimumStack(t *testing.T) {

	tests := []struct {
		name       string
		operations []string
		expected   []interface{}
	}{

		{
			name:       "example 1",
			operations: []string{"MinStack", "push", "1", "push", "2", "push", "0", "getMin", "pop", "top", "getMin"},
			expected:   []interface{}{nil, nil, nil, nil, nil, 0, nil, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minStack := Constructor()
			minStack.Push(1)
			minStack.Push(2)
			minStack.Push(0)
			minStack.GetMin() // return 0
			minStack.Pop()
			minStack.Top()    // return 2
			minStack.GetMin() // return 1
			if minStack.GetMin() != 1 {
				t.Errorf("MinimumStack() = %v, want %v", minStack.GetMin(), 1)
			}
		})
	}

}
