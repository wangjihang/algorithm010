package ladder_length

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if indexOf(endWord, wordList) == -1 {
		return 0
	}
	var level int
	usedM := make([]bool, len(wordList))
	queue := make([]string, 0, len(wordList))
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		level++
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				return level
			}
			for j, w := range wordList {
				if !usedM[j] && hasOneDiff(queue[i], w) {
					queue = append(queue, w)
					usedM[j] = true
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}

func indexOf(str string, arr []string) int {
	for i, s := range arr {
		if str == s {
			return i
		}
	}
	return -1
}

func hasOneDiff(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	if count == 1 {
		return true
	}
	return false
}
