// https://leetcode-cn.com/problems/short-encoding-of-words/

package word_encoding

import "sort"

// Map
// 时间复杂度：O(N*M)
// 空间复杂度：O(N)	// map大小
func minimumLengthEncoding(words []string) int {
	wm := make(map[string]struct{}, len(words))
	for _, v := range words {
		wm[v] = struct{}{}
	}
	for w := range wm {
		lw := len(w)
		for i := 1; i < lw; i++ {
			delete(wm, w[i:])
		}
	}
	count := 0
	for w := range wm {
		count += len(w) + 1
	}
	return count
}

// Trie
// 时间复杂度：O(N*M)
func minimumLengthEncoding2(words []string) int {
	var (
		count int
		root  = &Trie{next: [26]*Trie{}}
	)

	sort.Sort(StringSlice(words))

	for _, word := range words {
		count += root.insert(word)
	}
	return count
}

type Trie struct {
	next [26]*Trie
}

func (t *Trie) insert(word string) (count int) {
	curr := t
	for i := len(word) - 1; i >= 0; i-- {
		idx := word[i] - 'a'
		if curr.next[idx] == nil { // 只要有一个match不到 就不能压缩
			curr.next[idx] = &Trie{next: [26]*Trie{}}
			count = len(word) + 1
		}
		curr = curr.next[idx]
	}
	return count
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) > len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
