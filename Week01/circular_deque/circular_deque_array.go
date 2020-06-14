package circular_deque

//
//// 数组实现
//// 插入、删除: O(1)
//// 查询: O(n)
//// 空间复杂度: O(n)
//// 头插入 front 向前走
//// 尾插入 rear 向后走
//
//type MyCircularDeque struct {
//	arr      []int
//	capacity int
//	front    int
//	rear     int
//}
//
///** Initialize your data structure here. Set the size of the deque to be k. */
//func Constructor(k int) MyCircularDeque {
//	capacity := k + 1
//	return MyCircularDeque{
//		arr:      make([]int, capacity),
//		capacity: capacity,
//	}
//}
//
///** Adds an item at the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertFront(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//	this.front = (this.front - 1 + this.capacity) % this.capacity
//	this.arr[this.front] = value
//	return true
//}
//
///** Adds an item at the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertLast(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//	this.arr[this.rear] = value
//	this.rear = (this.rear + 1) % this.capacity
//	return true
//}
//
///** Deletes an item from the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteFront() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	this.front = (this.front + 1) % this.capacity
//	return true
//}
//
///** Deletes an item from the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteLast() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	this.rear = (this.rear - 1 + this.capacity) % this.capacity
//	return true
//}
//
///** Get the front item from the deque. */
//func (this *MyCircularDeque) GetFront() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.arr[this.front]
//}
//
///** Get the last item from the deque. */
//func (this *MyCircularDeque) GetRear() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.arr[(this.rear-1+this.capacity)%this.capacity]
//}
//
///** Checks whether the circular deque is empty or not. */
//func (this *MyCircularDeque) IsEmpty() bool {
//	return this.front == this.rear
//}
//
///** Checks whether the circular deque is full or not. */
//func (this *MyCircularDeque) IsFull() bool {
//	return (this.rear+1)%this.capacity == this.front
//}
//
///**
// * Your MyCircularDeque object will be instantiated and called as such:
// * obj := Constructor(k);
// * param_1 := obj.InsertFront(value);
// * param_2 := obj.InsertLast(value);
// * param_3 := obj.DeleteFront();
// * param_4 := obj.DeleteLast();
// * param_5 := obj.GetFront();
// * param_6 := obj.GetRear();
// * param_7 := obj.IsEmpty();
// * param_8 := obj.IsFull();
// */
