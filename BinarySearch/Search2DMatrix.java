import java.util.*;

public class Search2DMatrix {

	// You are given an m x n integer matrix matrix with the following two
	// properties:

	// Each row is sorted in non-decreasing order.
	// The first integer of each row is greater than the last integer of the
	// previous row.
	// Given an integer target, return true if target is in matrix or false
	// otherwise.

	// You must write a solution in O(log(m * n)) time complexity.

	public boolean searchMatrix(int[][] matrix, int target) {

		int left = 0;
		int right = matrix.length - 1;
		while (left <= right) {
			int pivot = left + (right - left) / 2;
			// isolate to solution array
			int first = matrix[pivot][0];
			int last = matrix[pivot][matrix[pivot].length - 1];
			if (target >= first && target <= last) {
				return binarySearch(matrix[pivot], target);
			} else if (target >= last) {
				left = pivot + 1;
			} else if (target <= first) {
				right = pivot - 1;
			}

		}

		return false;
	}

	public boolean binarySearch(int[] array, int target) {

		int left = 0;
		int right = array.length - 1;

		if (array.length == 0)
			return false;

		while (left <= right) {
			int pivot = left + (right - left) / 2;
			int currValue = array[pivot];

			if (currValue == target)
				return true;

			if (currValue > target) {
				right = pivot - 1;
			} else {
				left = pivot + 1;
			}

		}

		return false;
	}

	public static void main(String[] args) {

		int[][] matrix = { { 1, 3, 5, 7 }, { 10, 11, 16, 20 }, { 23, 30, 34, 60 } };
		int target = 3;

		Search2DMatrix solution = new Search2DMatrix();

		boolean result = solution.searchMatrix(matrix, target);

		System.out.println(result);

	}

}
