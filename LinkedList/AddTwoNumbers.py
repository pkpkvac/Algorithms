from typing import *
from definition import ListNode, create_linked_list, print_linked_list

# You are given two non-empty linked lists, l1 and l2, where each represents a non-negative integer.
# The digits are stored in reverse order, e.g. the number 123 is represented as 3 -> 2 -> 1 -> in the linked list.
# Each of the nodes contains a single digit. You may assume the two numbers do not contain any leading zero, except the number 0 itself.
# Return the sum of the two numbers as a linked list.

# Input: l1 = [1,2,3], l2 = [4,5,6]

# Output: [5,7,9]

# Explanation: 321 + 654 = 975.


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if not l1:
            return l2
        if not l2:
            return l1

        sumHead = ListNode()
        sumIterator = sumHead;        

        carry = 0 
        sumVal = l1.val + l2.val

        if sumVal >= 10:
            carry = 1
            sumVal = sumVal % 10

        l1 = l1.next
        l2 = l2.next

        sumHead.val = sumVal

        while l1 or l2:
            newSum = ListNode()
            sumIterator.next = newSum

            if l1 and l2:
                sumVal = l1.val + l2.val + carry
                l1 = l1.next
                l2 = l2.next
            
            elif l1:
                sumVal = l1.val + carry
                l1 = l1.next

            else: # l2 must exists
                sumVal = l2.val + carry
                l2 = l2.next

            if sumVal >= 10:
                carry = 1
                sumVal = sumVal % 10
            else:
                carry = 0

            newSum.val = sumVal
            sumIterator = sumIterator.next 

        # handle remaining carries
        if carry > 0:
            sumIterator.next = ListNode(carry)
        
        return sumHead



solution = Solution()

list1 =  create_linked_list([1,2,3])
list2 =  create_linked_list([4,5,6])

print(print_linked_list(solution.addTwoNumbers(list1=list1, list2=list2)))
