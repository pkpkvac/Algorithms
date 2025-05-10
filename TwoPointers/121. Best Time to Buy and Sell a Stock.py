from typing import List
class Solution:
    def max_profit_brute_force(self, prices: List[int]) -> int:
        max_val = 0
        for i, day_1 in enumerate(prices):
            for j, day_2 in enumerate(prices[i + 1:]):
                print(day_2, day_1)
                max_val = max(max_val, day_2 - day_1)
        print(max_val)
        return max_val

    def max_profit_two_pointers_NAIVE(self, prices: List[int]) -> int:
        # just because it's two pointers, doesn't imply 
        # one pointer is at the start, and the other is at the end
        # this simplistic algorithm fails for 
        # [2, 1, 4]
        # max_val = 0
        # b = 0
        # s = len(prices) - 1
        # while (b < s):
        #     sell_val = prices[s] - prices[b]
        #     max_val = max(max_val, sell_val)
        #     if sell_val < 0:
        #         b += 1
        #     else:
        #         s -= 1
        # print(max_val)
        # return max_val
        return None
    def max_profit_two_pointers(self, prices: List[int]) -> int:
        max_val = 0
        b = 0
        s = 1
        while s < len(prices):
            val = prices[s] - prices[b]
            if val >= 0:
                max_val = max(max_val, val)
            else:
                b = s
            s += 1
        print(max_val)
        return max_val


prices = [7,1,5,3,6,4]
# prices = [1,2,3,4,5]
# prices = [7,6,4,3,1]
solution = Solution()
# solution.max_profit_brute_force(prices)
solution.max_profit_two_pointers(prices)


        