import heapq

# A string s is called happy if it satisfies the following conditions:

# s only contains the letters 'a', 'b', and 'c'.
# s does not contain any of "aaa", "bbb", or "ccc" as a substring.
# s contains at most a occurrences of the letter 'a'.
# s contains at most b occurrences of the letter 'b'.
# s contains at most c occurrences of the letter 'c'.
# You are given three integers a, b, and c, return the longest possible happy string. If there are multiple longest happy strings, return any of them. If there is no such string, return the empty string "".

# A substring is a contiguous sequence of characters within a string.

# Example 1:

# Input: a = 3, b = 4, c = 2

# Output: "bababcabc"
# Example 2:

# Input: a = 0, b = 1, c = 5

# Output: "ccbcc"

class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        # always use the character with the most remaining count (greedy approach)
        heap = []
        res = ""

        if a > 0:
            heapq.heappush(heap, (-a, 'a'))
        if b > 0:
            heapq.heappush(heap, (-b, 'b'))
        if c > 0:
            heapq.heappush(heap, (-c, 'c'))

        while heap:

            first_count, first_char = heapq.heappop(heap)

            if len(res) >= 2 and res[-1] == res[-2] == first_char:
                # can't use first, use second, put back first
                if heap:
                    second_count, second_char = heapq.heappop(heap)
                    res += second_char
                    # push them both back now
                    heapq.heappush(heap, (first_count, first_char))
                    if second_count < -1:
                        heapq.heappush(heap, (second_count + 1, second_char))

            else:
                res += first_char
                if first_count < -1:
                    # still have remaining
                    heapq.heappush(heap, (first_count + 1, first_char))

        print(res)
        return res

solution = Solution()

# a = 0
# b = 1
# c = 5

a = 3
b = 4
c = 2

solution.longestDiverseString(a=a, b=b,c=c)




