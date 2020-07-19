package algorithm

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		depth     int
		length    = len(wordList)
		dict      = make(map[string]int, length)
		startQ    = make([]string, 0, length)
		endQ      = make([]string, 0, length)
		startUsed = make([]bool, length)
		endUsed   = make([]bool, length)
	)
	for i, w := range wordList {
		dict[w] = i
	}
	if i, ok := dict[endWord]; !ok {
		return 0
	} else {
		endUsed[i] = true
	}

	startQ = append(startQ, beginWord)
	endQ = append(endQ, endWord)

	for len(startQ) > 0 {
		depth++
		l := len(startQ)

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
