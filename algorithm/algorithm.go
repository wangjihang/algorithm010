package algorithm

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	length := len(s)
	for i := 0; i < length; i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
