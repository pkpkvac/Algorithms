import java.util.*;

public class ContainerWithMostWater {
	public int maxArea(int[] heights) {

		int left = 0;
		int right = heights.length - 1;

		int maxArea = 0;

		while (left < right) {

			int height = Math.min(heights[left], heights[right]);
			int length = right - left;

			maxArea = Math.max(maxArea, height * length);

			if (heights[right] > heights[left]) {
				left++;
			} else {
				right--;
			}

		}

		return maxArea;

	}

	public static void main(String[] args) {

		int[] height = { 1, 7, 2, 5, 4, 7, 3, 6 };

		ContainerWithMostWater solution = new ContainerWithMostWater();

		int result = solution.maxArea(height);

		System.out.println(result);

	}
}
