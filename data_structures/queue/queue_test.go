package datastructures

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run()
		})
	}
}
