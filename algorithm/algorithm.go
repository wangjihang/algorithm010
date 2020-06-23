package algorithm

type MyCircularDeque struct {
	head     *DeNode
	tail     *DeNode
	length   int
	capacity int
}

type DeNode struct {
	value int
	pre   *DeNode
	next  *DeNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := &DeNode{}
	tail := &DeNode{}
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
	node := &DeNode{
		value: value,
		pre:   this.head,
		next:  this.head.next,
	}
	this.head.next.pre = node
	this.head.next = node
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &DeNode{
		value: value,
		pre:   this.tail.pre,
		next:  this.tail,
	}
	this.tail.pre.next = node
	this.tail.pre = node
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	node := this.head.next
	node.next.pre = this.head
	this.head.next = node.next
	node.pre, node.next = nil, nil
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	node := this.tail.pre
	node.pre.next = this.tail
	this.tail.pre = node.pre
	node.pre, node.next = nil, nil
	this.length--
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
