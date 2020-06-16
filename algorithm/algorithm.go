package algorithm

type MyCircularDeque struct {
	head     *DeNode
	tail     *DeNode
	length   int
	capacity int
}

type DeNode struct {
	pre   *DeNode
	next  *DeNode
	value int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head:     &DeNode{},
		tail:     &DeNode{},
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.length++
	node := &DeNode{
		pre:   this.head,
		next:  this.head.next,
		value: value,
	}
	next := this.head.next
	next.pre = node
	this.head.next = node
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.length++
	node := &DeNode{
		pre:   this.tail.pre,
		next:  this.tail,
		value: value,
	}
	pre := this.tail.pre
	pre.next = node
	this.tail.pre = node
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.length--
	next := this.head.next
	this.head.next = next.next
	next.next.pre = this.head
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.length--
	tail := this.tail.pre
	this.tail.pre = tail.pre
	tail.pre.next = this.tail
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.value
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.pre.value
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
