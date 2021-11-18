package easy

func FirstUniqChar(s string) int {

	m := [26]int{}
	for _, b := range s {
		m[b-'a']++
	}

	for i, b := range s {
		if m[b-'a'] == 1 {
			return i
		}
	}

	return -1

}
