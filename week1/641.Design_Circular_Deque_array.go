type MyCircularDeque struct {
    Q []int
    capacity int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        Q: []int{},
        capacity: k,
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.Q = append([]int{value}, this.Q...)
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.Q = append(this.Q, value)
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.Q = this.Q[1:]
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    last := len(this.Q) - 1
    this.Q = this.Q[:last]
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Q[0]
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    last := len(this.Q) - 1
    return this.Q[last]
}


func (this *MyCircularDeque) IsEmpty() bool {
    return len(this.Q) == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return len(this.Q) == this.capacity
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
