package questions

type StackString struct {
	list []string
}

func (s *StackString) Push(input string) {
	s.list = append(s.list, input)
}

func (s *StackString) Pop() string {
	var poppedNumber string
	if len(s.list) > 0 {
		poppedNumber = s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
	}
	return poppedNumber
}

func IsValid(s string) bool {
	var isValid bool
	var stack StackString
	parentheses := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	if len(s)%2 != 0 {
		return false
	}
	isValid = true
	for index, character := range s {
		if string(character) == "(" || string(character) == "[" || string(character) == "{" {
			stack.Push(parentheses[string(character)])
		} else {
			if stack.Pop() != string(character) {
				return false
			}
		}
		if index == len(s)-1 && len(stack.list) > 0 {
			return false
		}
	}
	return isValid
}
