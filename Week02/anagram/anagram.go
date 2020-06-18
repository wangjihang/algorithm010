package anagram

// 数组法
// a-z字符串
func isAnagra(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']++
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
