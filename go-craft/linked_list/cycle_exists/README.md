Reverse Linked List
Given the beginning of a singly linked list head, reverse the list, and return the new beginning of the list.

Example 1:

Input: head = [0,1,2,3]

Output: [3,2,1,0]
Example 2:

Input: head = []

Output: []
Constraints:

0 <= The length of the list <= 1000.
-1000 <= Node.val <= 1000

Hint: Think Iteratively
Instead of trying to move pointers towards each other, think about reversing links one by one as you traverse forward.
Key Insight:
You need to reverse the direction of each link while keeping track of:
The current node you're processing
The previous node (so you can point current back to it)
The next node (so you don't lose the rest of the list)
