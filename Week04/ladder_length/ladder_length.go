package ladder_length

// bfs
// 时间复杂度：O(M*N) M单词长度 N列表中单词数量
// 空间复杂度：O(M*N) dict 最坏情况下 每个预处理的单词 要存储N个单词
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		dict    = make(map[string][]string)
		visited = make(map[string]bool)
		queue   = &Queue{}
	)
	// 预处理，将所有的通用状态存储到dict
	for _, word := range wordList {
		for i := range word {
			key := word[:i] + "*" + word[i+1:]
			dict[key] = append(dict[key], word)
		}
	}

	queue.Push(&Pair{beginWord, 1})

	for !queue.IsEmpty() {
		pair := queue.Pop()
		visited[pair.Word] = true

		// process node/ generate sub node/ queue add
		for i := range pair.Word {
			for _, word := range dict[pair.Word[:i]+"*"+pair.Word[i+1:]] {
				if word == endWord { // 匹配
					return pair.Depth + 1
				}
				if !visited[word] { // 防止环路
					queue.Push(&Pair{word, pair.Depth + 1})
					visited[word] = true
				}
			}
		}
	}
	return 0
}

type Pair struct {
	Word  string
	Depth int
}

type Queue []*Pair

func (q *Queue) Push(pair *Pair) {
	*q = append(*q, pair)
}

func (q *Queue) Pop() *Pair {
	if q.IsEmpty() {
		return nil
	}

	pair := (*q)[0]
	*q = (*q)[1:]
	return pair
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
