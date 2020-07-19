package ladder_length

// 双向bfs
// 时间复杂度：O(N*M) N为数组长度，M为每个单词的长度
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		depth     int
		length    = len(wordList)
		dict      = make(map[string]int)
		startQ    = make([]string, 0, length)
		startUsed = make([]bool, length)
		endQ      = make([]string, 0, length)
		endUsed   = make([]bool, length)
	)
	for idx, word := range wordList {
		dict[word] = idx
	}
	if idx, ok := dict[endWord]; !ok {
		return 0
	} else {
		endUsed[idx] = true
	}

	startQ = append(startQ, beginWord)
	endQ = append(endQ, endWord)

	for len(startQ) > 0 {
		l := len(startQ)
		depth++

		for i := 0; i < l; i++ {
			chars := []byte(startQ[i])
			for i := range chars {
				o := chars[i]
				for c := 'a'; c <= 'z'; c++ {
					chars[i] = byte(c)
					idx, ok := dict[string(chars)]
					if !ok || startUsed[idx] {
						continue
					}
					if endUsed[idx] {
						return depth + 1
					}
					startQ = append(startQ, wordList[idx])
					startUsed[idx] = true
				}
				chars[i] = o
			}
		}

		startQ = startQ[l:]
		if len(startQ) > len(endQ) {
			startQ, endQ = endQ, startQ
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}
