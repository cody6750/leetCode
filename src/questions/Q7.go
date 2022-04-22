package questions

import (
	"math"
)

func ReverseInt(x int) int {
	var reverse int
	s := Stack{}
	for x != 0 {
		s.Push(x % 10)
		x = x / 10
	}
	lengthList := len(s.List)
	for i := 0; i < lengthList; i++ {
		reverse += s.Pop() * int(math.Pow(10, float64(i)))
		if float64(reverse) < -math.Pow(2, 31) || float64(reverse) > math.Pow(2, 31) {
			return 0
		}
	}

	return reverse
}
