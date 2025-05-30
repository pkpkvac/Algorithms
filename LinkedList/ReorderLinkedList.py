from typing import Optional
from definition import ListNode, create_linked_list, print_linked_list

# You are given the head of a singly linked-list.

# The positions of a linked list of length = 7 for example, can intially be represented as:

# [0, 1, 2, 3, 4, 5, 6]

# Reorder the nodes of the linked list to be in the following order:

# [0, 6, 1, 5, 2, 4, 3]

# Notice that in the general case for a list of length = n the nodes are reordered to be in the following order:

# [0, n-1, 1, n-2, 2, n-3, ...]

# You may not modify the values in the list's nodes, but instead you must reorder the nodes themselves.

# Example 1:

# Input: head = [2,4,6,8]

# Output: [2,8,4,6]
# Example 2:

# Input: head = [2,4,6,8,10]

# Output: [2,10,4,8,6]


class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        # initial thoughts:
        # count all the nodes, and use that
        nodes = []
        curr = head  # Save the head pointer
        while curr:
            nodes.append(curr)
            curr = curr.next

        l, r = 0, len(nodes) - 1

        while l < r:
            # Connect left to right
            nodes[l].next = nodes[r]
            l += 1
            
            if l >= r:  # Prevent connecting when pointers meet
                break
            
            # Connect right to next left
            nodes[r].next = nodes[l]
            r -= 1

        # Terminate the list
        nodes[l].next = None
        return nodes[0]  # Return the original head
    

head = [2,4,6,8]
#  head = [2,4,6,8,10]

headList = create_linked_list(head)

solution = Solution()

head = solution.reorderList(head=headList)

print_linked_list(head)

