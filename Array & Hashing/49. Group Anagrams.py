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

    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        # Dictionary to store groups of anagrams
        anagram_groups = {}
        
        for s in strs:
            # Create a count array for 26 lowercase letters
            char_count = [0] * 26
            
            # Count frequency of each character in the string
            for char in s:
                # Convert character to index (a->0, b->1, etc.)
                char_index = ord(char) - ord('a')
                char_count[char_index] += 1
            
            # Convert the count array to a tuple to use as dictionary key
            # (lists can't be keys in dictionaries because they're mutable)
            count_tuple = tuple(char_count)
            
            # Add the string to its anagram group
            if count_tuple in anagram_groups:
                anagram_groups[count_tuple].append(s)
            else:
                anagram_groups[count_tuple] = [s]

        # Return all the anagram groups
        return list(anagram_groups.values())


strs = ["eat","tea","tan","ate","nat","bat"]    
    
solution = Solution()
# res = solution.group_anagrams_brute_force(strs)
res = solution.group_anagrams(strs)
print(res)