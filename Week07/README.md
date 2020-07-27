学习笔记

## 字典树
### 定义
字典树，即 Trie 树，又称单词查找树或键树，是一种树形结构。典型应用是用于统计和排序大量的字符串（但不仅限于
字符串），所以经常被搜索引擎系统用于文本词频统计。
### 优点
空间换时间
### 基本性质
- 节点本身不存完整单词
- 从根结点到某一结点，路径上经过的字符连接起来，为该结点对应的字符串
- 每个节点的所有子节点路径代表的字符都不相同
### 模板
```go
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return true
}
```

## 并查集
### 基本特征
- new(n)：建立一个新的并查集，其中包含n个元素
- union(x,y): 把元素x和元素y所在集合合并，要求x和y所在的集合不相交，相交则不合并
- find(x): 找到元素x所在集合的代表，该操作也可以用于判断两个元素是否位于同一个集合，只要将他们各自的代表比较一下即可。
### 优化
路径压缩：将find的空间复杂度优化到O(1)
### 场景
朋友圈、组团、配对问题
### 代码模板：
```go
type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
```

## 搜索
### 初级搜索
- 朴素搜索
- 优化方式：不重复(fibonacci[记忆化])、剪枝(生成括号问题)
- 搜索方向：
    - DFS：depth first search 深度优先搜索
    - BFS：breadth first search 广度优先搜索
    
### 高级搜索
- 双向BFS
- 启发式搜索

### 启发式搜索
- 估价函数：h(n) 用来评价哪些结点最有希望是一个我们要找的节点，h(n)会返回一个非负实数，也可以认为是从结点n的目标结点路径的估计成本。

## 高级树
### 二叉搜索树
二叉搜索树，也称有序二叉树（Ordered Binary Tree）、排
序二叉树（Sorted Binary Tree），是指一棵空树或者具有下列性质的二叉
树：
1. 左子树上所有结点的值均小于它的根结点的值；
2. 右子树上所有结点的值均大于它的根结点的值；
3. 以此类推：左、右子树也分别为二叉查找树。 （这就是 重复性！）

中序遍历：升序排列
#### 保证性能的关键
- 保证二维维度：左右子树结点平衡(避免退化成链表)
- balanced

### AVL树
1. 发明者 G. M. Adelson-Velsky和 Evgenii Landis
2. **Balance Factor（平衡因子）：是它的左子树的高度减去它的右子树的高度（有时相反）。balance factor = {-1, 0, 1}**
3. **通过旋转操作来进行平衡（四种）**
4. https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree
#### 旋转操作进行平衡
- 右右子树-->左旋
- 左左子树-->右旋
- 左右子树-->左右旋
- 右左子树-->右左旋
#### AVL总结
1. 平衡二叉搜索树
2. 每个结点存`balance factor = {-1, 0, 1}`
3. 四种旋转操作

不足：结点需要存储额外信息、且调整次数频繁

### 红黑树
红黑树是一种近似平衡的二叉搜索树（Binary Search Tree），它能够确保任何一个结点的左右子树的高度差小于两倍。具体来说，红黑树是满足如下条件的二叉搜索树： 
- 每个结点要么是红色，要么是黑色
- 根结点是黑色
- 每个叶结点（NIL结点，空结点）是黑色的 
- `不能有相邻接的两个红色结点` 
- `从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点。`

#### 关键性质
从根到叶子的最长的可能路径不多于最短的可能路径的两倍长。

### AVL和 Red-Black Tree 比较
- AVL trees provide `faster lookups` than Red Black Trees because they are `more strictly balanced`. 
- Red Black Trees provide `faster insertion and removal` operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.
- AVL trees store balance factors or heightswith each node, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of information per node.
- Red Black Trees are used in most of the `language libraries likemap, multimap, multisetin C++` whereas AVL trees are used in `databases` where faster retrievals are required.