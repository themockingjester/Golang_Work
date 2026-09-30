package striver_dsa_sheet

type DoublyLinkedListNode struct {
	left  *DoublyLinkedListNode
	right *DoublyLinkedListNode
	val   int
	key   int
}

type LRUCache struct {
	memory     map[int]*DoublyLinkedListNode
	head       *DoublyLinkedListNode
	last       *DoublyLinkedListNode
	listLength int
	capacity   int
}

func (this *LRUCache) CheckEviction(key int) {
	_, ok := this.memory[key]
	if ok {
		// Since elment  present so no eviction
		return
	}
	if this.capacity == this.listLength {
		if this.capacity == 0 {
			return
		} else if this.capacity == 1 {
			nodeKey := this.head.key
			delete(this.memory, nodeKey)
			this.head = nil
			this.listLength--
			this.last = nil
			return
		} else {

			nodeKey := this.head.key
			delete(this.memory, nodeKey)
			this.head = this.head.right
			this.head.left = nil
			this.listLength--
		}

	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, listLength: 0, memory: make(map[int]*DoublyLinkedListNode)}
}

func (this *LRUCache) rearrangeMemory(key, value int, isPut bool) {
	val, ok := this.memory[key]
	if ok {
		// Key exists
		if this.capacity <= 1 {
			if isPut {
				this.memory[key].val = value
			}

			return
		} else if val.right == nil {
			// already last
			if isPut {
				val.val = value
			}

			return
		} else if val.left == nil {
			// first
			currNode := this.memory[key]
			if isPut {
				currNode.val = value
			}
			this.head = currNode.right
			currNode.right.left = nil
			currNode.right = nil
			this.last.right = currNode
			currNode.left = this.last
			this.last = currNode
		} else {
			currNode := this.memory[key]
			if isPut {
				currNode.val = value

			}
			currNode.left.right = currNode.right
			currNode.right.left = currNode.left
			this.last.right = currNode
			currNode.left = this.last
			currNode.right = nil
			this.last = currNode
		}
	} else {
		if this.capacity == 0 {
			return
		}
		// Key dont exists
		currNode := DoublyLinkedListNode{
			val: value,
			key: key,
		}
		if this.last == nil {
			this.head = &currNode
			this.last = &currNode
			this.listLength++
		} else {
			this.last.right = &currNode
			currNode.left = this.last
			this.last = &currNode
			this.listLength++
		}
		this.memory[key] = this.last
	}

}

func (this *LRUCache) Get(key int) int {
	val, ok := this.memory[key]
	if ok {
		this.rearrangeMemory(key, val.val, false)

		return val.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.listLength > 0 {
		this.CheckEviction(key)
	}

	this.rearrangeMemory(key, value, true)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
