import java.util.*;
import java.sql.Array;

// A palindrome is a string that reads the same forward and 
// backward. It is also case-insensitive and ignores all non-alphanumeric characters.

public class ValidPalindrome {

	public boolean isPalindrome(String s) {

		int left = 0;
		int right = s.length() - 1;
		System.out.println(s);

		while (left < right) {

			while (left < right && !Character.isLetterOrDigit(s.charAt(left))) {
				left++;
			}

			while (left < right && !Character.isLetterOrDigit(s.charAt(right))) {
				right--;
			}
			char leftChar = Character.toLowerCase(s.charAt(left));
			char rightChar = Character.toLowerCase(s.charAt(right));

			if (leftChar != rightChar)
				return false;

			left++;
			right--;

		}

		return true;
	}

	public static void main(String[] args) {

		// String s = "Was it a car or a cat I saw?";
		// String s = "race. car?";
		String s = "tab a cat";

		ValidPalindrome solution = new ValidPalindrome();

		Boolean result = solution.isPalindrome(s);

		System.out.println(result);

	}

}
