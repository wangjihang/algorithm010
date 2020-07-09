// https://leetcode-cn.com/problems/re-space-lcci/
// 恢复空格

package main

// 动态规划
// 时间复杂度：O(N^2)
func respace(dictionary []string, sentence string) int {
	m := make(map[string]bool, len(dictionary))
	n := len(sentence)
	dp := make([]int, n+1)

	for _, v := range dictionary {
		m[v] = true
	}

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if m[sentence[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, v := range word {
		pos := v - 'a'
		if curr.next[pos] == nil {
			curr.next[pos] = &Trie{next: [26]*Trie{}}
		}
		curr = curr.next[pos]
	}
	curr.isEnd = true
}

func main() {
	respace([]string{"vprkj", "sqvuzjz", "ptkrqrkussszzprkqrjrtzzvrkrrrskkrrursqdqpp", "spqzqtrqs", "rkktkruzsjkrzqq", "rk", "k", "zkvdzqrzpkrukdqrqjzkrqrzzkkrr", "pzpstvqzrzprqkkkd", "jvutvjtktqvvdkzujkq", "r", "pspkr", "tdkkktdsrkzpzpuzvszzzzdjj", "zk", "pqkjkzpvdpktzskdkvzjkkj", "sr", "zqjkzksvkvvrsjrjkkjkztrpuzrqrqvvpkutqkrrqpzu"},
		"rkktkruzsjkrzqqzkvdzqrzpkrukdqrqjzkrqrzzkkrr")
}
