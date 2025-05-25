from typing import List
# You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.

# Return the integer that represents the evaluation of the expression.

# The operands may be integers or the results of other operations.
# The operators include '+', '-', '*', and '/'.
# Assume that division between integers always truncates toward zero.
# Example 1:

# Input: tokens = ["1","2","+","3","*","4","-"]

# Output: 5

# Explanation: ((1 + 2) * 3) - 4 = 5


# I am trying to understand this better. I have identified that if I have a character that is either a +, -, *, or "/", I want to pop two values off the stack and perform an operation, at the end the stack should have just one value - I think, I'm not sure. Then I return that value. What is the issue with this algorithm

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:

        stack = []

        for token in tokens:
            if token == "+":
                val1 = stack.pop()
                val2 = stack.pop()
                stack.append(val2 + val1)
            elif token == "*":
                val1 = stack.pop()
                val2 = stack.pop()
                stack.append(val2 * val1)
            elif token == "-":
                val1 = stack.pop()
                val2 = stack.pop()
                stack.append(val2 - val1)
            elif token == "/":
                val1 = stack.pop()
                val2 = stack.pop()
                # Use int() to truncate toward zero
                stack.append(int(val2 / val1))
            else:
                # Convert string to integer before pushing
                stack.append(int(token))

        return stack.pop()

tokens = ["1","2","+","3","*","4","-"]

solution = Solution() 
result = solution.evalRPN(tokens)
print(f"Result: {result}")