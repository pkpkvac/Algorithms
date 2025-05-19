from typing import List

# You are given an array of strings strs. Return the longest common prefix of all the strings.

# If there is no longest common prefix, return an empty string "".

# Example 1:

# Input: strs = ["bat","bag","bank","band"]

# Output: "ba"
# Example 2:

# Input: strs = ["dance","dag","danger","damage"]

# Output: "da"
# Example 3:

# Input: strs = ["neet","feet"]

# Output: ""

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 0: return ""

        res = ""
        shortestWord = len(strs[0])
        for word in strs:
            shortestWord = min(shortestWord, len(word))

        for i in range(shortestWord):
            letter = strs[0][i]
            for word in strs:
                if word[i] != letter:
                    return res
            res += letter

        return res
    

solution = Solution()
# strs = ["bat","bag","bank","band"]
# strs = ["dance","dag","danger","damage"]
strs = ["neet","feet"]
print(solution.longestCommonPrefix(strs))