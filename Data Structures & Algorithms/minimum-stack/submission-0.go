type MinStack struct {
    stack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack{
        stack : make([]int,0),
        minStack : make([]int,0),
    }
}


func (this *MinStack) Push(value int)  {
    this.stack = append(this.stack,value)
    if len(this.minStack) == 0{
        this.minStack = append(this.minStack,value)
    } else{
        currVal := this.minStack[len(this.minStack)-1]
        if value < currVal{
            this.minStack = append(this.minStack,value)
        }else{
            this.minStack = append(this.minStack,currVal)
        }
    }
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0{
        return
    }
    this.stack = this.stack[:len(this.stack) - 1]
    this.minStack = this.minStack[:len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}