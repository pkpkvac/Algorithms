from typing import *
from definition import ListNode, create_linked_list, print_linked_list

# Input: list1 = [1,2,4], list2 = [1,3,5]

# Output: [1,1,2,3,4,5]
# Example 2:

# Input: list1 = [], list2 = [1,2]

# Output: [1,2]

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
                
        l1 = list1
        l2 = list2

        if not list1:
            return l2
        if not list2:
            return l1    
        
        head = ListNode()

        if l1.val <= l2.val:
            head = l1
            l1 = l1.next
        else:
            head = l2
            l2 = l2.next

        iterator = head

        while l1 or l2:
            if l1 and l2:
                if l1.val <= l2.val:
                    iterator.next = l1
                    iterator = iterator.next
                    l1 = l1.next
                else:
                    iterator.next = l2
                    iterator = iterator.next
                    l2 = l2.next
            else:
                if l1 :
                    iterator.next = l1
                    break
                else:
                    iterator.next = l2
                    break

        return head

                


solution = Solution()

list1 =  create_linked_list([1,2,4])
list2 =  create_linked_list([1,3,5])

print(print_linked_list(solution.mergeTwoLists(list1=list1, list2=list2)))
