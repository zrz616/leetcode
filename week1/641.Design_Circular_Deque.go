type Node struct {
    Val int
    Pre *Node
    Next *Node
}

type MyCircularDeque struct {
    capacity int
    size int
    head *Node
    rear *Node
}


func Constructor(k int) MyCircularDeque {
    deque := MyCircularDeque{}
    deque.capacity = k
    deque.size = 0
    deque.head = &Node{}
    deque.rear = &Node{}
    deque.head.Next = deque.rear
    deque.rear.Pre = deque.head
    return deque
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.size ++
    newNode := &Node{}
    newNode.Val = value
    newNode.Next = this.head.Next
    this.head.Next.Pre = newNode
    this.head.Next = newNode
    newNode.Pre = this.head
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.size ++
    newNode := &Node{}
    newNode.Val = value
    newNode.Pre = this.rear.Pre
    this.rear.Pre.Next = newNode
    this.rear.Pre = newNode
    newNode.Next = this.rear
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    } 
    this.size--
    deleteNode := this.head.Next
    this.head.Next = deleteNode.Next
    deleteNode.Next.Pre = this.head
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.size--
    deleteNode := this.rear.Pre
    this.rear.Pre = deleteNode.Pre
    deleteNode.Pre.Next = this.rear
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.head.Next.Val
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.rear.Pre.Val   
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.capacity
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
