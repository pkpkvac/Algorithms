import heapq
from typing import List

# There is a car with capacity empty seats. The vehicle only drives
#  east (i.e., it cannot turn around and drive west).

# You are given the integer capacity and a integer
#  array trips where trips[i] = [numPassengers[i], from[i], to[i]] indicates 
# that the ith trip has numPassengers[i] passengers and the
#  locations to pick them up and drop them off are from[i] and to[i] 
# respectively. The locations are given as the number of kilometers
#  due east from the car's initial location.

# Return true if it is possible to pick up and drop off
#  all passengers for all the given trips, or false otherwise.

# Example 1:

# Input: trips = [[4,1,2],[3,2,4]], capacity = 4

# Output: true
# Example 2:

# Input: trips = [[2,1,3],[3,2,4]], capacity = 4

# Output: false


class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        # 1. sort by pick up location
        trips.sort(key=lambda trip: trip[1])

        # 2. Min heap stores drop off locations (to_loc)
        dropoff_heap = []
        current_passengers = 0

        # 3. process each pickup
        for num_passengers, pickup_loc, dropoff_loc in trips:

            # 4. Remove passengers who have been dropped off
            # If the earliest dropoff location is at or behind where we are now,
            #  those passengers have already gotten off, so remove them from our count
            while dropoff_heap and dropoff_heap[0][0] <= pickup_loc:
                _, passengers_dropped = heapq.heappop(dropoff_heap)
                current_passengers -= passengers_dropped

            # 5. Pick up new passengers
            current_passengers += num_passengers

            # 6. Check capacity
            if current_passengers > capacity:
                return False

            # 7. Add their dropoff time to heap
            heapq.heappush(dropoff_heap, (dropoff_loc, num_passengers))

        return True



# trips = [[4,1,2],[3,2,4]]
trips = [[2,1,3],[3,2,4]]
capacity = 4

solution = Solution()

print(solution.carPooling(trips, capacity))
        
