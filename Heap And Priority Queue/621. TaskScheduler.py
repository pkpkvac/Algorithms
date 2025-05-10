from typing import List
import heapq
# You are given an array of CPU tasks tasks, where tasks[i] is an uppercase english character from A to Z. You are also given an integer n.
# Each CPU cycle allows the completion of a single task, and tasks may be completed in any order.
# The only constraint is that identical tasks must be separated by at least n CPU cycles, to cooldown the CPU.
# Return the minimum number of CPU cycles required to complete all tasks.

# Example 1:
# Input: tasks = ["X","X","Y","Y"], n = 2

# Output: 5
# Explanation: A possible sequence is: X -> Y -> idle -> X -> Y.

# Example 2:
# Input: tasks = ["A","A","A","B","C"], n = 3

# Output: 9

class Solution:
    # def leastInterval(self, tasks: List[str], n: int) -> int:
    #     frequency_map = [0]*26
    #     for letter in tasks:
    #         frequency_map[ord(letter) - ord('A')] += 1
    #     # print(frequency_map)

    #     heap = []
    #     for f in frequency_map:
    #         if f > 0:
    #             heapq.heappush(heap, -f)

    #     print(heapq)


        
# solution = Solution()


# tasks = ["A","A","A","B","C"]
# n = 3

# solution.leastInterval(tasks, n)