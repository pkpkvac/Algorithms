# Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

# Implement the MyStack class:

# void push(int x) Pushes element x to the top of the stack.
# int pop() Removes the element on the top of the stack and returns it.
# int top() Returns the element on the top of the stack.
# boolean empty() Returns true if the stack is empty, false otherwise.
# Notes:

# You must use only standard operations of a queue, which means that only push to back,
# peek/pop from front, size and is empty operations are valid.
# Depending on your language, the queue may not be supported natively. 
# You may simulate a queue using a list or deque (double-ended queue)
#  as long as you use only a queue's standard operations.

# Example 1:

# Input
# ["MyStack", "push", "push", "top", "pop", "empty"]
# [[], [1], [2], [], [], []]
# Output
# [null, null, null, 2, 2, false]

# Explanation
# MyStack myStack = new MyStack();
# myStack.push(1);
# myStack.push(2);
# myStack.top(); // return 2
# myStack.pop(); // return 2
# myStack.empty(); // return False

class MyStackPopExpensive:

    def __init__(self):
        self.q1 = []
        self.q2 = []

    def push(self, x: int) -> None:
        self.q1.append(x)

    def pop(self) -> int:
        # Move all but the last element to q2
        while len(self.q1) > 1:
            self.q2.append(self.q1.pop(0))
        
        # The last element in q1 is our answer
        result = self.q1.pop(0)
        
        # Swap queues: q1 becomes q2, q2 becomes empty
        self.q1 = self.q2
        self.q2 = []
        
        return result

    def top(self) -> int:
        return self.q1[len(self.q1) - 1]
        

    def empty(self) -> bool:
        return len(self.q1) == 0


class MyStackPushExpensive:

    def __init__(self):
        self.q1 = []
        self.q2 = []

    def push(self, x: int) -> None:
        # Put new element in q2
        self.q2.append(x)
        
        # Move all elements from q1 to q2 (so new element is at front)
        while len(self.q1) > 0:
            self.q2.append(self.q1.pop(0))
        
        # Swap queues: q1 becomes q2, q2 becomes empty
        self.q1 = self.q2
        self.q2 = []

    def pop(self) -> int:
        return self.q1.pop(0)

    def top(self) -> int:
        return self.q1[0]
        

    def empty(self) -> bool:
        return len(self.q1) == 0


obj = MyStackPopExpensive()

obj.push(20)
obj.push(30)

print(obj.top())

print(obj.pop())

print(obj.pop())


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()