package questions

func addBinary(a string, b string) string {
	binary := map[string]string{
		"000": "00",
		"001": "10",
		"010": "10",
		"011": "01",
		"100": "10",
		"101": "01",
		"110": "01",
		"111": "11",
	}
	var max, low, result, finalString, carryOver string
	carryOver = "0"
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if len(a) > len(b) {
		max = a
		low = b
	} else {
		max = b
		low = a
	}

	for index := range max {
		if len(low)-index <= 0 {
			entry := string(max[len(max)-1-index]) + "0" + carryOver
			result = binary[entry]
		} else {
			entry := string(max[len(max)-1-index]) + string(low[len(low)-1-index]) + carryOver
			result = binary[entry]
		}
		carryOver = string(result[1])
		result = string(result[0])
		finalString = result + finalString
	}
	if carryOver == "1" {
		finalString = carryOver + finalString
	}
	return finalString
}
