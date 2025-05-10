import java.util.*;

// Given an integer array nums and an integer k,
//  return true if there are two distinct indices 
//  i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

public class ContainsDuplicate2 {

	public boolean containsNearbyDuplicateFirstAttempt(int[] nums, int k) {

		if (nums.length <= 1)
			return false;

		int left = 0;
		int right = left + 1;

		while (right < nums.length) {
			if (right - left <= k) {
				if (nums[left] == nums[right]) {
					return true;
				}
				right++;
			} else {
				// r has gone too farm, boost l, set r = left + 1
				left++;
				right = left + 1;
			}
		}
		return false;

	}

	public boolean containsNearbyDuplicateHashSet(int[] nums, int k) {

		HashSet<Integer> window = new HashSet<>();

		for (int i = 0; i < nums.length; i++) {
			// if window size exceeds k, remove leftmost element
			if (i > k) {
				window.remove(nums[i - k - 1]);
			}

			// if current element is already in window, found a duplicate
			if (!window.add(nums[i])) {
				return true;
			}

		}
		return false;
	}

	public boolean containsNearbyDuplicateHashMap(int[] nums, int k) {

		HashMap<Integer, Integer> map = new HashMap<>();

		for (int i = 0; i < nums.length; i++) {
			if (map.containsKey(nums[i])) {
				if (i - map.get(nums[i]) <= k) {
					return true;
				}
			}
			map.put(nums[i], i);
		}

		return false;

	}

	public static void main(String[] args) {

		int[] nums = { 1, 2, 3, 1 };
		// int[] nums = { 1, 2, 3, 4 };
		int k = 3;

		ContainsDuplicate2 solution = new ContainsDuplicate2();

		boolean result = solution.containsNearbyDuplicateFirstAttempt(nums, k);

		System.out.println(result);

	}
}
