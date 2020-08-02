学习笔记

## 位运算
### 概念
例如 5：`0000 0101`
- 反码：将二进制数按位取反。`1111 1010`
- 补码：反码+1。`1111 1011`
- 负数：在计算机中，负数以其正值的补码形式表达。`1111 1011`

注意：正数的源码、反码、补码都相同。

### 指定位置的位运算符
1. 将 x 最右边的 n 位清零：x& (~0 << n)
2. 获取 x 的第 n 位值（0 或者 1）： (x >> n) & 1
3. 获取 x 的第 n 位的幂值：x& (1 <<n)
4. 仅将第 n 位置为 1：x | (1 << n)
5. 仅将第 n 位置为 0：x & (~ (1 << n))
6. 将 x 最高位至第 n 位（含）清零：x& ((1 << n) -1)
7. 将第 n 位至第 0 位（含）清零：x& (~ ((1 << (n + 1)) -1))

### 实战位运算要点
- 判断奇偶：
> x % 2 == 1 —> (x & 1) == 1
> x % 2 == 0 —> (x & 1) == 0
- x >> 1 —> x / 2. 
> 即： x = x / 2; —> x = x >> 1;
> mid = (left + right) / 2; —> mid = (left + right) >> 1;
- X = X & (X-1) 清零最低位的 1 
- X & -X => 得到最低位的 1 
- X & ~X => 0

## 排序
### 快速排序(Quick Sort)
数组取标杆pivot，将小元素放pivot左边，大元素放右侧，然后依次对左边和右边的子数组继续快排，以达到整个序列有序。

code:
```go
func quickSort(arr []int, begin, end int) {
	if end <= begin {
		return
	}
	pivot := partition(arr, begin, end)
	quickSort(arr, begin, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, begin, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}
```

### 归并排序(Merge Sort) -- 分治
1. 把长度为n的输入序列分成两个长度为n/2的子序列；
1. 把这两个子序列分别采用归并排序；
1. 将两个排好序的子序列合并成一个最终的排序序列。

code:
```go
func mergeSort(arr []int, left, right int) {
	if right <= left {
		return
	}
	mid := (left + right) >> 1

	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for ; i <= mid && j <= right; k++ {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		temp[k] = arr[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		temp[k] = arr[j]
	}
	copy(arr[left:], temp)
}
```

### 堆排序(Heap Sort) -- 堆插入O(logN)，取最大/小值O(1)
1. 数组元素依次建立小顶堆
1. 依次取堆顶元素，并删除
