package main

import (
	"reflect"
	"testing"
)

func TestDelSliceIndex(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		index  int
		output []int
	}{
		{
			name:   "Delete from middle",
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			index:  2,
			output: []int{1, 2, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "Delete from beginning",
			input:  []int{1, 2, 3, 4, 5},
			index:  0,
			output: []int{2, 3, 4, 5},
		},
		{
			name:   "Delete from end",
			input:  []int{1, 2, 3, 4, 5},
			index:  4,
			output: []int{1, 2, 3, 4},
		},
		{
			name:   "Delete out of range",
			input:  []int{1, 2, 3, 4, 5},
			index:  5,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Delete from empty slice",
			input:  []int{},
			index:  0,
			output: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DelSliceIndex(tt.input, tt.index)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("Expected %v, but got %v", tt.output, res)
			}
		})
	}
}
