from typing import List

# There are n cars at given miles away from the starting mile 0, traveling to reach the mile target.
# You are given two integer array position and speed, both of length n, where position[i] is the starting mile of the ith car and speed[i] is the speed of the ith car in miles per hour.
# A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.
# A car fleet is a car or cars driving next to each other. The speed of the car fleet is the minimum speed of any car in the fleet.
# If a car catches up to a car fleet at the mile target, it will still be considered as part of the car fleet.
# Return the number of car fleets that will arrive at the destination.

# Input: 
target = 12
position = [10,8,0,5,3]
speed = [2,4,1,1,3]
# Output: 3

# Input: target = 100, position = [0,2,4], speed = [4,2,1]
# Output: 1

class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        pairs = []
        fleets = 1
        for i in range(len(position)):
            pairs.append([position[i], speed[i]])

        sorted_pairs = sorted(pairs, key=lambda x: x[0], reverse=True)

        time_lead = (target - sorted_pairs[0][0])/sorted_pairs[0][1]
        for i in range(1, len(sorted_pairs)):
            t_i = (target - sorted_pairs[i][0])/sorted_pairs[i][1]
            if t_i > time_lead:
                time_lead = max(t_i, time_lead)
                fleets += 1
        return fleets

solution = Solution()

solution.carFleet(target, position, speed)