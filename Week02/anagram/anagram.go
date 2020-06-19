package anagram

func groupAnagrams(strs []string) [][]string {
	arrM := make(map[[26]int][]string)
	for _, str := range strs {
		arr := [26]int{}
		for _, v := range str {
			arr[v-'a']++
		}
		arrM[arr] = append(arrM[arr], str)
	}
	result := make([][]string, 0, len(arrM))
	for _, strs := range arrM {
		result = append(result, strs)
	}
	return result
}
