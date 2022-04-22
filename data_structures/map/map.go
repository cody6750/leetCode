package datastructures

import "log"

/* Mao
Description:
- Datastructure that contains key value pairs
- Lookup is constant time because inserting the key and value into the map using a hashing function takes that long
*/
func initMap() {

	// 1. Create map using make
	a := make(map[string]string)
	a[""] = ""

	// 2. Create map using short hand decleration
	b := map[string]string{
		"": "",
	}
	b[""] = ""

	// 3. Create map using var. Var allocates memory. Make initializes the map. Without make, the function errors out because the map isn't intialized.
	var c map[string]string
	c = make(map[string]string)
	c[""] = ""

	// Trick to check if value exist within map
	if value, exist := c[""]; exist {
		log.Print(value)
	}

	// Used to delete key
	delete(c, "")
}
