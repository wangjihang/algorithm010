package group_anagrams

// Map去重
// O(N*M)
// 空间: O(N)
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var arr [26]int
		for _, ch := range str {
			arr[ch-'a']++
		}
		m[arr] = append(m[arr], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// O(n)
func anagrams(s, t string) bool {
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
