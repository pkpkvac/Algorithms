import java.sql.Array;
import java.util.*;

public class MajorityElement {

	public int majorityElement(int[] nums) {
		// this solution utilizes a trick called hte Boyer-Moore voting algorithm
		// you need to track only a count, and the element, and follow these properties:
		// 1. Start with a candidate and count:
		// Initialize candidate to any value (first element is fine)
		// Initialize count to 0
		// Process each element:
		// If count is 0, set the current element as the new candidate
		// If the current element equals the candidate, increment count
		// If the current element differs from the candidate, decrement count
		// Return the final candidate:
		// After processing all elements, the candidate will be the majority element

		// For array [2,2,1,1,1,2,2]:
		// candidate = 2, count = 1
		// See 2: candidate = 2, count = 2
		// See 1: candidate = 2, count = 1
		// See 1: candidate = 2, count = 0
		// See 1: candidate = 1, count = 1 (reset candidate when count was 0)
		// 6. See 2: candidate = 1, count = 0
		// See 2: candidate = 2, count = 1 (reset candidate when count was 0)
		// Final candidate = 2, which is the majority element.
		// No sorting or hash map required - just two variables!

		if (nums.length == 0)
			return 0;

		int candidate = nums[0];
		int count = 1;

		for (int i = 1; i < nums.length; i++) {

			int newCandidate = nums[i];

			if (candidate == newCandidate) {
				count++;
			} else if (candidate != newCandidate) {

				count--;

				if (count == 0) {
					candidate = newCandidate;
					count++;
				}

			}

		}

		return candidate;

	}

	public int majorityElementHashMap(int[] nums) {

		HashMap<Integer, Integer> map = new HashMap<>();

		for (int number : nums) {
			map.put(number, map.getOrDefault(number, 0) + 1);
		}

		for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
			System.out.println(entry);
			int key = entry.getKey();
			int count = entry.getValue();

			if (count > nums.length / 2) {
				return key;
			}

		}

		return 0;

	}

	// 🧪 Test cases go here
	public static void main(String[] args) {
		MajorityElement solution = new MajorityElement();

		int[] nums = { 2, 2, 1, 1, 1, 2, 2 };
		// int result = solution.majorityElementHashMap(nums);
		int result = solution.majorityElement(nums);
		System.out.println(result);
	}
}
