from typing import List

# Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

# In other words, return true if one of s1's permutations is the substring of s2.

# Example 1:
# Input: s1 = "ab", s2 = "eidbaooo"
# Output: true
# Explanation: s2 contains one permutation of s1 ("ba").


# Example 2:
# Input: s1 = "ab", s2 = "eidboaoo"
# Output: false


# initial thoughts:
# for s2 to contain a permutation of s1,
# s1 lenght will be equal to or shorter s2 length
# the number of letters in s1 requires
# that the number of adjacent letters in s2 be the same length
# this means we could use a sliding window for this problem
# keep the character frequency of s1 in a map
# build a sliding window of length s1 over s2
# if you're on a character of s2 and it exists in s1, subtract the
# frequency of that character in the map, if it exists
# if the sliding window is of length s1, and all the values in the
# map are 0, then s2 contains a permutation of s1

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        
        # Count characters in s1
        s1_count = {}
        for char in s1:
            s1_count[char] = s1_count.get(char, 0) + 1
        
        # Use sliding window of size len(s1)
        window_size = len(s1)
        left = 0
        window_count = {}
        
        for right in range(len(s2)):
            # Add right character to window
            char = s2[right]
            window_count[char] = window_count.get(char, 0) + 1
            
            # If window size exceeds s1 length, shrink from left
            if right - left + 1 > window_size:
                left_char = s2[left]
                window_count[left_char] -= 1
                if window_count[left_char] == 0:
                    del window_count[left_char]
                left += 1
            
            # Check if current window matches s1
            if right - left + 1 == window_size and window_count == s1_count:
                return True
        
        return False


solution = Solution()

s1 = "ab"
s2 = "eidbaooo"

print(solution.checkInclusion(s1, s2))