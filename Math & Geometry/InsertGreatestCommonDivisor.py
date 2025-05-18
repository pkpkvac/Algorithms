import sys
import os

# Add the parent directory (Algorithms) to the Python path
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

# Now you can import from LinkedList
from LinkedList.definition import ListNode, create_linked_list, print_linked_list

# Given the head of a linked list head, in which each node contains an integer value.
# Between every pair of adjacent nodes, insert a new node with a value equal to the greatest common divisor of them.
# Return the linked list after insertion.
# The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.


# Input: head = [18,6,10,3]
# Output: [18,6,6,2,10,1,3]
# Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram 
# denotes the linked list after inserting the new nodes (nodes in blue are the inserted nodes).
# - We insert the greatest common divisor of 18 and 6 = 6 between the 1st and the 2nd nodes.
# - We insert the greatest common divisor of 6 and 10 = 2 between the 2nd and the 3rd nodes.
# - We insert the greatest common divisor of 10 and 3 = 1 between the 3rd and the 4th nodes.
# There are no more adjacent nodes, so we return the linked list.
# Example 2:


# Input: head = [7]
# Output: [7]
# Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes.
# There are no pairs of adjacent nodes, so we return the initial linked list.


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
from typing import Optional

class Solution:
    def insertGreatestCommonDivisors(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next: return head
        
        current = head

        while current.next != None:
            gcd = self.findGCD(current.val, current.next.val)

            gcdNode = ListNode(gcd)

            next_node = current.next

            current.next = gcdNode
            gcdNode.next = next_node

            current = next_node
        
        return head 

    def findGCD(self, first: int, second: int) -> int:
        while second:
            first, second = second, first % second
        return first


head = [18,6,10,3]
headNode = create_linked_list(head)

solution = Solution()

# Don't forget to pass headNode to the function
print_linked_list(solution.insertGreatestCommonDivisors(headNode))