package datastructures

/* SLICE
Description: An array that is flexible, convenient and cheap
- Flexible in capacity
- Cheap as in it is passed as a reference or pointer to functions
- Can set length and capacity
*/

func initArray() {
	// 1. Creating slice using zerio value
	var a []int
	a = append(a, 2)

	// 2. Creating slice using shorthand declaration
	b := []int{1, 2, 3, 4, 5}
	b = append(b, 2)

	// 3. Creating slize using make function
	// make([]T, len, capacity)
	c := make([]int, 4, 4)
	c = append(c, 4)
}
