package main

import "testing"

func Test_initHeap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initHeap()
		})
	}
}
