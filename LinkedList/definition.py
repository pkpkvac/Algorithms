# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def create_linked_list(input_arr):
    """
    Creates a linked list from an array of values.
    
    Args:
        input_arr (list): List of values to create nodes from
        
    Returns:
        ListNode or None: Head of the created linked list, or None if input is empty
    """
    if not input_arr:
        return None
    
    # Create head node from first element
    head = ListNode(input_arr[0])
    
    # Current node to build the list
    current = head
    
    # Add remaining elements
    for i in range(1, len(input_arr)):
        new_node = ListNode(input_arr[i])
        current.next = new_node
        current = new_node
    
    return head

def print_linked_list(node):
    """
    Converts a linked list to an array of values for printing.
    
    Args:
        node (ListNode or None): Head of the linked list
        
    Returns:
        list: Array containing all values in the linked list
    """
    result = []
    
    # Traverse the list and collect values
    while node is not None:
        result.append(node.val)
        node = node.next
    
    return result 