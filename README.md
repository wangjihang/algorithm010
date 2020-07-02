# 极客大学「算法训练营-第10期」作业提交仓库

### 五毒神掌

| 题目 | week | 第一二遍 | 第三遍 | 第四遍 | 第五遍 |
|---|---|---|---|---|---|
| [两数之和](https://leetcode-cn.com/problems/two-sum/) | week01 | 6.14 done | 6.15 done | 6.21 done | 待定 |
| [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/) | week01 | 6.15 done | 6.17 done | 6.22 done | 待定 |
| [删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) | week01 | 6.15 done | 6.17 done | 6.22 done | 待定 |
| [设计循环双端队列](https://leetcode.com/problems/design-circular-deque/) | week01 | 6.15 done | 6.17 done | 6.22 done | 待定 |
| [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/description/) | week02 | 6.18 done | 6.19 done | 6.25 done | 待定 |
| [丑数](https://leetcode-cn.com/problems/chou-shu-lcof/) | week02 | 6.18 done | 6.19 done | 6.25 done | 待定 |
| [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/) | week02 | 6.18 done | 6.19 done | 6.25 done | 待定 |
| [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/) | week02 | 6.19 done | 6.20 done | 6.26 done | 待定 |
| [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water//) | week02 | 6.19 done | 6.21 done | 6.26 done | 待定 |
| [N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/) | week02 | 6.20 done | 6.21 done | 6.27 done | 待定 |
| [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/) | week02 | 6.20 done | 6.21 done | 6.27 done | 待定 |
| [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | week03 | 6.28 done | 6.29 done | 7.05 | 待定 |
| [全排列](https://leetcode-cn.com/problems/permutations/) | week03 | 6.28 done | 6.29 done | 7.05 | 待定 |
| [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/) | week03 | 6.29 done | 6.30 done | 7.06 | 待定 |
| [组合](https://leetcode-cn.com/problems/combinations/) | week03 | 6.29 done | 6.30 done | 7.06 | 待定 |
| [全排列 II ](https://leetcode-cn.com/problems/permutations-ii/) | week03 | 6.29 done | 6.30 done | 7.06 | 待定 |

# qms
## 关于场景
- 适用场景
    - API服务
    - 需要对HTTP|GRPC协议支持的服务
- 暂不适用场景
    - Job任务
    - HTTP页面渲染

## 关于监控
- qms监控会[自动收集](https://git.qutoutiao.net/gopher/qms/blob/master/docs/development.md#43-%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0metrics%E7%9B%91%E6%8E%A7) [advance.yaml中autometrics.enabled:true]，**不需要提工单**。
- 容器监控是由容器进行收集，如果没有数据，请核对容器工单中监控端口是否正确。