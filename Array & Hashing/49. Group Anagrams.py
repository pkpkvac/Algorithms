from typing import List
# Input: strs = ["eat","tea","tan","ate","nat","bat"]

# Output: [["bat"],["nat","tan"],["ate","eat","tea"]]


class Solution:
    def group_anagrams_brute_force_totally_bad(self, strs: List[str]) -> List[List[str]]:
        # two loops totally unnecessary, also wrong.
        map = {}
        for i, outer in enumerate(strs):
            sorted_str_outer = ''.join(sorted(outer))
            if sorted_str_outer not in map:
                map[sorted_str_outer]=[outer]

            for j, inner in enumerate(strs, i + 1):
                sorted_str_inner = ''.join(sorted(inner))
                if sorted_str_inner in map:
                    if inner not in map.get(sorted_str_inner):
                        map[sorted_str_outer].append(inner)
        values = map.values()
        result = list(values)
        return result
        
        
    def group_anagrams_bad_soluton(self, strs: List[str]) -> List[List[str]]:
        # sorting makes O(n * k log k)
        map = {}
        for string in strs:
            sorted_str = ''.join(sorted(string))
            if sorted_str not in map:
                map[sorted_str] = [string]
            else:
                map[sorted_str].append(string)
        print(map)
        return list(map.values())
    
    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        # using just the letter frequency
        # O(n * k)
        letter_freq = {}
        for string in strs:
            key_str_int = [0] * 26
            print(key_str_int)
            for char in string:
                print(char)
                num = ord(char) - ord('a')
                key_str_int[num] += 1
                key_str = (''.join(map(str,key_str_int)))
            if key_str not in letter_freq:
                letter_freq[key_str] = [string]
            else:
                letter_freq[key_str].append(string)
        print(letter_freq)
        return list(letter_freq.values())


strs = ["eat","tea","tan","ate","nat","bat"]    
    
solution = Solution()
# res = solution.group_anagrams_brute_force(strs)
res = solution.group_anagrams(strs)
print(res)