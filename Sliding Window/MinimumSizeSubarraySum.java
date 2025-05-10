import java.util.*;

public class MinimumSizeSubarraySum {

	public int minSubArrayLen(int target, int[] nums) {

		int left = 0;
		int right = 0;
		int minWindow = Integer.MAX_VALUE;
		int currSum = 0;

		while (right < nums.length) {
			// Add the current element to our running sum
			currSum += nums[right];

			// Shrink window from the left as long as sum >= target
			while (currSum >= target) {
				// Update minimum window size
				minWindow = Math.min(minWindow, right - left + 1);

				// Remove leftmost element and move left pointer
				currSum -= nums[left];
				left++;
			}

			// Expand window to the right
			right++;
		}

		// Return 0 if no valid subarray found
		return minWindow == Integer.MAX_VALUE ? 0 : minWindow;

	}

	public static void main(String[] args) {

		int[] nums = { 2, 3, 1, 2, 4, 3 };
		int target = 7;

		MinimumSizeSubarraySum solution = new MinimumSizeSubarraySum();

		int result = solution.minSubArrayLen(target, nums);

		System.out.println(result);

	}

}
