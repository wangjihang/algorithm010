package ladder_length

// bfs
// 时间复杂度：O(M*N) M单词长度 N列表中单词数量
// 空间复杂度：O(M*N) dict 最坏情况下 每个预处理的单词 要存储N个单词
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	var (
		isContain bool
		depth     int
		dict      = make(map[string][]int)
		visited   = make(map[int]bool)
		queue     = make([]string, 0, len(wordList))
	)
	// 预处理，将所有的通用状态存储到dict
	for idx, word := range wordList { // O(N*M)
		for i := range word {
			key := word[:i] + "*" + word[i+1:]
			dict[key] = append(dict[key], idx)
		}
		if endWord == word {
			isContain = true
		}
	}

	if !isContain {
		return 0
	}

	queue = append(queue, beginWord)

	for len(queue) > 0 {
		l := len(queue)
		depth++

		for i := 0; i < l; i++ {
			word := queue[i]
			for i := range word {
				for _, idx := range dict[word[:i]+"*"+word[i+1:]] {
					if endWord == wordList[idx] {
						return depth + 1
					}
					if !visited[idx] {
						visited[idx] = true
						queue = append(queue, wordList[idx])
					}
				}
			}
		}

		queue = queue[l:]
	}
	return 0
}
