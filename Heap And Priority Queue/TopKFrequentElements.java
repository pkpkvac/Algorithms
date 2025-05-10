import java.sql.Array;
import java.util.*;

public class TopKFrequentElements {

	public int[] topKFrequentMinHeap(int[] nums, int k) {
		// this implementation uses a priority queue (a min heap)
		// first obtain all the frequencies

		// The min heap approach is more efficient when K is small compared
		// to the number of unique elements, as you only need to maintain K elements in
		// the heap.

		HashMap<Integer, Integer> map = new HashMap<>();

		// get the frequencies in a map
		for (int num : nums) {
			map.put(num, map.getOrDefault(num, 0) + 1);
		}

		// now use a priority queue,
		// order by ascending
		PriorityQueue<Integer> minHeap = new PriorityQueue<>((o1, o2) -> map.get(o1) - map.get(o2));

		System.out.print(minHeap);
		// Add elements to heap
		for (int num : map.keySet()) {
			minHeap.add(num);
			if (minHeap.size() > k) {
				minHeap.poll(); // remove least frequent element
			}
		}

		System.out.print(minHeap);

		// build result array from heap
		int[] result = new int[k];
		for (int i = k - 1; i >= 0; i--) {
			result[i] = minHeap.poll();
		}

		// System.out.print(Arrays.toString(result));

		return result;

	}

	public int[] topKFrequentMaxHeap(int[] nums, int k) {
		// this implementation uses a priority queue (a min heap)
		// first obtain all the frequencies

		// The min heap approach is more efficient when K is small compared
		// to the number of unique elements, as you only need to maintain K elements in
		// the heap.

		HashMap<Integer, Integer> map = new HashMap<>();

		// get the frequencies in a map
		for (int num : nums) {
			map.put(num, map.getOrDefault(num, 0) + 1);
		}

		// now use a priority queue,
		// order by ascending
		PriorityQueue<Integer> maxHeap = new PriorityQueue<>((o1, o2) -> map.get(o2) - map.get(o1));

		System.out.print(maxHeap);
		// Add elements to heap
		for (int num : map.keySet()) {
			maxHeap.add(num);

		}

		int count = 0;
		int[] result = new int[k];
		while (count < k) {
			result[count] = maxHeap.poll();
			count++;
		}

		// System.out.print(Arrays.toString(result));

		return result;

	}

	// 🧪 Test cases go here
	public static void main(String[] args) {

		int nums[] = { 1, 1, 1, 2, 2, 3 };
		int k = 2;

		TopKFrequentElements solution = new TopKFrequentElements();
		// int[] result = solution.topKFrequentMinHeap(nums, k);
		int[] result = solution.topKFrequentMaxHeap(nums, k);
		System.out.println(result);
	}
}
