package datastructures

import "fmt"

/*

New:
- Allocates memory and sets it to the zero value for all tpyes including referenced data types. Thus slics, maps and channels return a pointer to nothing.
- Creates a pointer reference to all data type. Composite(slice, map, chan) and non composite (int, string, struct,etc)
- Behaves like var (when the variable is not initialized)

Make:
- Allocates memory and intiailizes the data type.
- Behaves like :=
- Allocates memory for referenced data types (slice, map, chan), plus initializes their underlying data structures


*/
func main() {
	fmt.Println("-- MAKE --")
	a := make([]int, 0)
	aPtr := &a
	fmt.Println("pointer == nil :", *aPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *aPtr)

	fmt.Println("-- COMPOSITE LITERAL --")
	b := []int{}
	bPtr := &b
	fmt.Println("pointer == nil :", *bPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *bPtr)

	fmt.Println("-- NEW --")
	cPtr := new([]int)
	fmt.Println("pointer == nil :", *cPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *cPtr)

	fmt.Println("-- VAR (not initialized) --")
	var d []int
	dPtr := &d
	fmt.Println("pointer == nil :", *dPtr == nil)
	fmt.Printf("pointer value: %p\n", *dPtr)
}
