import java.sql.Array;
import java.util.*;

public class TopKFrequentElements {

	public int[] topKFrequent(int[] nums, int k) {
		// this is an O(nlogn) time

		HashMap<Integer, Integer> map = new HashMap<>();

		// get the frequencies in a map
		for (int num : nums) {
			map.put(num, map.getOrDefault(num, 0) + 1);
		}

		// get the top k most frequent elements in the map
		// do this by converting the map keys into a list, then sort
		// by their frequency
		// use an array list because it's dynamic, and allows use of objects,
		// not just primitives

		List<Integer> list = new ArrayList<>(map.keySet());

		// sort this list based on the maps frequencies,
		// in descending order
		list.sort((a, b) -> map.get(b) - map.get(a));

		System.out.println(list);

		// now you can pick the top k elements from the list
		int[] result = new int[k];
		for (int i = 0; i < k; i++) {
			result[i] = list.get(i);
		}

		return result;
	}

	// 🧪 Test cases go here
	public static void main(String[] args) {

		int nums[] = { 1, 1, 1, 2, 2, 3 };
		int k = 2;

		TopKFrequentElements solution = new TopKFrequentElements();
		int[] result = solution.topKFrequent(nums, k);
		System.out.println(Arrays.toString(result));
	}
}
