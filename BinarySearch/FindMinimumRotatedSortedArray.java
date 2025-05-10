import java.util.*;

public class FindMinimumRotatedSortedArray {

	// You are given an m x n integer matrix matrix with the following two
	// properties:

	// Each row is sorted in non-decreasing order.
	// The first integer of each row is greater than the last integer of the
	// previous row.
	// Given an integer target, return true if target is in matrix or false
	// otherwise.

	// You must write a solution in O(log(m * n)) time complexity.

	public int findMin(int[] nums) {

	}

	public static void main(String[] args) {

		int[] nums = { 3, 4, 5, 1, 2 };

		FindMinimumRotatedSortedArray solution = new FindMinimumRotatedSortedArray();

		int result = solution.findMin(nums);

		System.out.println(result);

	}

}
