package circular_deque

//
//// 链表设计法
//// 插入和删除都是O(1)
//// 查询: O(n)

type MyCircularDeque struct {
	head     *CircularDeque
	tail     *CircularDeque
	length   int
	capacity int
}

type CircularDeque struct {
	data int
	pre  *CircularDeque
	next *CircularDeque
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	// circular
	head := &CircularDeque{}
	tail := &CircularDeque{}
	head.next = tail
	tail.pre = head

	return MyCircularDeque{
		head:     head,
		tail:     tail,
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.length++
	node := &CircularDeque{
		data: value,
	}
	nextnode := this.head.next
	nextnode.pre = node
	node.next = nextnode
	node.pre = this.head
	this.head.next = node
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.length++
	node := &CircularDeque{
		data: value,
	}
	prenode := this.tail.pre
	prenode.next = node
	node.pre = prenode
	node.next = this.tail
	this.tail.pre = node
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.length--
	nnextnode := this.head.next.next
	nnextnode.pre = this.head
	this.head.next = nnextnode
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.length--
	pprenode := this.tail.pre.pre
	pprenode.next = this.tail
	this.tail.pre = pprenode
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.data
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.pre.data
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
