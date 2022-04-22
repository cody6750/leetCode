package datastructures

import "testing"

/* ARRAY
Description:
- Stores key and value
- Also known as hashes and dictionaries in other programming languages
- O(1) lookup time
- Zero value of map is nil
*/
func Test_initMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMap()
		})
	}
}
