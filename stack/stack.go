package stack

type (
    Stack struct {      // 栈类,
        top *node       // 当前栈顶所对应的元素信息,
        length int      // 当前栈的长度
    }
    node struct {       // 栈中节点类的设计
        value interface{}
        prev *node
    }
)


// Create a new stack
func New() *Stack {
    return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
    return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
    if this.length == 0 {
        return nil
    }
    return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
    if this.length == 0 {
        return nil
    }

    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
    n := &node{value,this.top}
    this.top = n
    this.length++
}