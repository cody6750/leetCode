package questions

func StrStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}

	diff := len(haystack) - len(needle)
	for i := 0; i <= diff; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// BRUTE FORCE
// func StrStr(haystack string, needle string) int {
// 	if len(needle) == 0 {
// 		return 0
// 	}
// 	if len(needle) > len(haystack) {
// 		return -1
// 	}
// 	for i := 0; i < len(haystack); i++ {
// 		if string(haystack[i]) == string(needle[0]) {
// 			if len(needle) == 1 {
// 				return i
// 			}
// 			for j := 1; j < len(needle); j++ {
// 				if i+j == len(haystack) {
// 					return -1
// 				}
// 				if string(haystack[i+j]) != string(needle[j]) {
// 					break
// 				} else if j == len(needle)-1 {
// 					return i
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }
