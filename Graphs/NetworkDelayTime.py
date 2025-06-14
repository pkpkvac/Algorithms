import heapq
from typing import List

# You are given a network of n directed nodes, labeled from 1 to n. You are also given times, a list of directed edges where times[i] = (ui, vi, ti).

# ui is the source node (an integer from 1 to n)
# vi is the target node (an integer from 1 to n)
# ti is the time it takes for a signal to travel from the source to the target node (an integer greater than or equal to 0).
# You are also given an integer k, representing the node that we will send a signal from.

# Return the minimum time it takes for all of the n nodes to receive the signal. If it is impossible for all the nodes to receive the signal, return -1 instead.

# Example 1:

# Input: times = [[1,2,1],[2,3,1],[1,4,4],[3,4,1]], n = 4, k = 1

# Output: 3
# Example 2:

# Input: times = [[1,2,1],[2,3,1]], n = 3, k = 2

# Output: -1

# Why This is Perfect for Dijkstra:
# 1. Single-source shortest path: Find shortest time from node k to all other nodes
# 2. Weighted graph: Each edge has a travel time ti
# 3. Non-negative weights: All times ≥ 0 (required for Dijkstra)
# 4. The key insight:
# Signal travels simultaneously to all paths
# All nodes receive signal when the  FURTHEST node gets it
# So answer = max(shortest_time_to_each_node)


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # Step 1: Build adjacency list
        graph = {}
        for u, v, time in times:
            if u not in graph:
                graph[u] = []
            graph[u].append((v, time))  # (neighbor, travel_time)
        
        print(f"Graph: {graph}")
        
        # Step 2: Initialize distances and priority queue
        distances = {}
        for i in range(1, n + 1):  # Nodes are labeled 1 to n
            distances[i] = float('inf')
        distances[k] = 0  # Starting node has distance 0
        
        # Priority queue: (distance, node)
        heap = [(0, k)]  # Start with (distance=0, node=k)
        
        print(f"Initial distances: {distances}")
        print(f"Initial heap: {heap}")
        
        # Step 3: Dijkstra's main loop
        while heap:
            current_dist, current_node = heapq.heappop(heap)
            
            print(f"\nProcessing node {current_node} with distance {current_dist}")
            
            # Skip if we've already found a better path to this node
            if current_dist > distances[current_node]:
                continue
            
            # Check all neighbors of current node
            if current_node in graph:
                # {1: [(2, 1), (4, 4)], 2: [(3, 1)], 3: [(4, 1)]}
                print(f"\ngraph[current_node] {graph[current_node]}")
                for neighbor, travel_time in graph[current_node]:
                    new_distance = current_dist + travel_time
                    
                    # If we found a shorter path, update it
                    if new_distance < distances[neighbor]:
                        distances[neighbor] = new_distance
                        heapq.heappush(heap, (new_distance, neighbor))
                        print(f"  Updated {neighbor} distance to {new_distance}")
        
        print(f"\nFinal distances: {distances}")
        
        # Step 4: Check if all nodes are reachable and find the answer
        max_time = 0
        for node in range(1, n + 1):
            if distances[node] == float('inf'):
                print(f"Node {node} is unreachable!")
                return -1
            max_time = max(max_time, distances[node])
        
        print(f"Answer: {max_time}")
        return max_time

times = [[1,2,1],[2,3,1],[1,4,4],[3,4,1]]
n = 4
k = 1
solution = Solution()
solution.networkDelayTime(times, n, k)