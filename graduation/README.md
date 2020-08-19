## 1. 如何高效学习
### 1.1 摒弃旧习惯
- 不要死磕
- 五毒神掌(敢于放手、敢于死记硬背代码)
- 不懒于看高票答案

最佳方式：5-15分钟想不出来，直接看题解或者高票代码，用五毒神掌变成自己的东西。

最差方式：不借助外部帮助，死磕2-3小时或一晚上，导致精疲力尽，丧失信心.

### 1.2 学习方法上
- 改变自己的学习习惯（不要死磕）
- 五毒神掌（重要的是五！过遍数，而不是每次花伸长时间）
- 不要死磕AC，要看高票答案和高质量题解

最大误区：LeetCode题目只做一遍！

## 2. 总览
### 2.1 精通一个领域
- Chunk it up 切碎知识点
- Deliberate Practicing 刻意练习
- Feedback 反馈

### 2.2 数据结构
- 一维：
    - 基础：数组array，链表linked list
    - 高级：栈stack、队列queue、双端队列deque、集合set、映射map(hash or map)、etc
- 二维：
    - 基础：树tree、图graph
    - 高级：二叉搜索树binary search tree（red-black tree、AVL）、堆heap、并查集disjoint set、字典树Trie、etc
- 特殊：
    - 位运算bitwise、布隆过滤器bloomFilter
    - LRU Cache
   
### 2.3 算法
- if-else，switch--->branch
- for,while loop--->iteration
- 递归Recursion(Divide & Conquer，Backtrace)
- 搜索Search：深度优先搜索 Depth first search，广度优先搜索Breadth first search，A*，etc
- 动态规划Dynamic Programming
- 二分查找Binary Search
- 贪心Greedy
- 数学Math，几何Geometry

### 2.4 切题四件套
- Clarification 沟通
- Possible solutions 所有可能的解法 比较优劣
- Coding
- Test case

### 2.5 五毒神掌
- 第一遍
    - 5-15分钟：读题+思考
    - 直接看题解：注意多解法，比较解法优劣
    - 背诵、默写好的解法
- 第二遍
    - 马上自己写-->LeetCode提交
    - 多种解法比较、体会-->优化
- 第三遍
    - 过了一天，再重复做题
    - 不同解法的熟练程度-->专项练习
- 第四遍
    - 过了一周：反复回来练习相同的题目
- 第五遍
    - 前一周恢复性训练

## 3. 时间空间复杂度
### 3.1 Big O notation
- O(1): Constant Complexity 常数复杂度
- O(log n): Logarithmic Complexity 对数复杂度
- O(n): Linear Complexity 线性复杂度
- O(n^2): N square Complexity 平方复杂度
- O(n^3): N cubic Complexity 立方复杂度
- O(2^n): Exponential Complexity 指数复杂度
- O(n^2): Factorial 阶乘复杂度

### 3.2 时间复杂度曲线
![大O表示法](https://raw.githubusercontent.com/wangjihang/img/master/20200819190353.png?token=AFCBVMTIFZMRWHWS2YXQQW27HUDV2)

## 4 数据结构
### 4.1 数组Array
### 4.2 链表Linked List
### 4.3 双向链表Double Linked List
### 4.4 跳表Skip List
#### 4.4.1 跳表的特点
- **只能用于元素有序的情况**
- 跳表对标的是平衡树(AVL)和二分查找，是一种插入/删除/搜索都是O(log n)的数据结构
- 优势：原理简单、容易实现、方便扩展、效率更高

升维思想+空间换时间

#### 4.4.2 examples
给有序链表加速
![给有序的链表加速](https://raw.githubusercontent.com/wangjihang/img/master/20200819192217.png?token=AFCBVMXGDKUVLSWX4N7YSVC7HUF2S)
多级索引
![多级索引](https://raw.githubusercontent.com/wangjihang/img/master/20200819192253.png?token=AFCBVMWCJ5SS4EZL47RMNNK7HUF44)
现实中跳表的形态
![现实中跳表的形态](https://raw.githubusercontent.com/wangjihang/img/master/20200819192534.png?token=AFCBVMWU7ELK3K3555AAV7S7HUGG2)

### 4.5 栈Stack
![Stack](https://raw.githubusercontent.com/wangjihang/img/master/20200819192822.png?token=AFCBVMRTUP45NUBWPF3QQNK7HUGRO)

Stack：先入后出；添加、删除皆为O(1)

### 4.6 队列Queue
![queue](https://raw.githubusercontent.com/wangjihang/img/master/20200819192856.png?token=AFCBVMWKWAD2RFR22H7EJUC7HUGTQ)

Queue：先入先出；添加、删除皆为O(1)

### 4.6 双端队列Deque
![Deque](https://raw.githubusercontent.com/wangjihang/img/master/20200819193052.png?token=AFCBVMRMLAGDTRSBFZ236A27HUG2W)

- 简单理解：双端可以进出的Queue
- 插入和删除都是O(1)操作

### 4.7 优先队列
- 插入操作：O(1)
- 取出操作：O(logN)-按照元素的优先级取出
- 底层具体实现的数据结构较为多样和复杂：heap、bst、treap

TODO: Go heap实现代码

### 4.8 哈希表HashTable
- 哈希表也叫散列表，是根据关键码值(Key value)而直接进行访问的数据结构；
- 通过把关键码值映射到表中的一个位置来访问记录，以加快查找的速度；
- 这个映射函数叫做散列函数。

注意：哈希冲突(拉链式)、因子、链表

Map：key-value对，key不重复
Set：不重复元素的集合

### 4.9 树Tree
Linked List是特殊化的Tree

Tree是特殊化的Graph
1. 树Tree
![Tree](https://raw.githubusercontent.com/wangjihang/img/master/20200819195141.png?token=AFCBVMRA5PSFDJFPPNE6HSK7HUJI4)

2. 二叉树
- 遍历：
    - 前序(pre-order): 根-左-右
    - 中序(in-order): 左-根-右
    - 后序(post-order): 左-右-根
- 示例代码：
![示例代码](https://raw.githubusercontent.com/wangjihang/img/master/20200819195450.png?token=AFCBVMRO7Z5ORPP36CV4GPK7HUJUU)

3. 二叉搜索树Binary Search Tree
- 定义：二叉搜索树也称有序二叉树、排序二叉树，是指一颗空树或者具有下列性质的二叉树：
    - 左子树上所有结点的值均小于它的根结点的值；
    - 右子树上所有结点的值均大于它的根节点的值；
    - 以此类推：左、右子树也分别为二叉查找树(重复性)。

### 4.18 复杂度分析
![复杂度分析](https://raw.githubusercontent.com/wangjihang/img/master/20200819193456.png?token=AFCBVMWELLDLZSOEA6WJMPS7HUHKI)

### 工具
- https://visualgo.net/zh/bst?slide=1