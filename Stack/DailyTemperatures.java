import java.util.*;

public class DailyTemperatures {

	public int[] dailyTemperatures(int[] temperatures) {

		// easiest is to use brute force solution
		int[] result = new int[temperatures.length];
		Stack<Integer> stack = new Stack<>();

		for (int i = 0; i < temperatures.length; i++) {

			// In the array [2, 1, 1, 3], we don't perform any pop
			// operations while processing [2, 1, 1] because these elements
			// are already in decreasing order. However, when we reach 3, we
			// pop elements from the stack until the top element of the stack
			// is no longer less than the current element. For each popped element,
			// we compute the difference between the indices and store it in the
			// position corresponding to the popped element.

			while (!stack.isEmpty() && temperatures[i] > temperatures[stack.peek()]) {
				int prevIndex = stack.pop();

				// find days to wait
				result[prevIndex] = i - prevIndex;

			}

			stack.push(i);
		}

		return result;

	}

	public static void main(String[] args) {

		DailyTemperatures solution = new DailyTemperatures();

		int[] temperatures = { 30, 38, 30, 36, 35, 40, 28 };

		int[] result = solution.dailyTemperatures(temperatures);

		System.out.println(Arrays.toString(result));

	}

}
