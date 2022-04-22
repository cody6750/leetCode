package datastructures

import "log"

/* ARRAY
Description:
-In Go, an array is a numbered sequence of elements of a specific length.
- Arrays in go are passed in as values not pointers. This means that when an array is passed to a function, the whole array is
copied to that function. This is very expensive, especially when you have a large array.
- Compared to Slices, they are not resizable.
*/

func initArray() {
	// 1. Creating Array using zero value
	var a [5]int
	a[0] = 0

	//2. Creating array using shorthand declaration
	b := [5]int{1, 2, 3, 4, 5}
	log.Print(b)

}
