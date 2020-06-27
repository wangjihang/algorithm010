package anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	for _, value := range arr {
		if value != 0 {
			return false
		}
	}
	return true
}
