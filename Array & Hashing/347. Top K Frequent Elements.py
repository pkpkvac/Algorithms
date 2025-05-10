from typing import List
class Solution:
    def topKFrequentNaive(self, nums: List[int], k: int) -> List[int]:
        map = {}
        for num in nums:
            map[num] = map.get(num, 0) + 1

        print(map)
        sorted_items = sorted(map.items(), key=lambda x: x[1], reverse=True)

        print(sorted_items[:k])

        return [num for num, freq in sorted_items[:k]]

    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq_map = {}
        res = []
        for num in nums:
           freq_map[num] = freq_map.get(num, 0) + 1 

        buckets = [[] for _ in range(len(nums) + 1)]

        for num, freq in freq_map.items():
            buckets[freq].append(num)


        for freq in range(len(buckets) - 1, -1, -1):
            if len(res) >= k:
                break

            for num in buckets[freq]:
                if len(res) >= k:
                    break
                res.append(num)

        print(res)
        return res



        

nums = [1,2,2, 2,3,3,3, 4, 4]
k = 3
solution = Solution()
# print(solution.topKFrequentNaive(nums, k))
print(solution.topKFrequent(nums, k))