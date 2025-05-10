import java.sql.Array;
import java.util.*;

public class Solution {

	public boolean isAnagram(String s, String t) {

		int[] charCount = new int[26];

		if (s.length() != t.length()) {
			return false;
		}

		for (int i = 0; i < s.length(); i++) {
			charCount[s.charAt(i) - 'a']++;

			charCount[t.charAt(i) - 'a']--;
		}

		// for (int count : charCount) {
		// if (count != 0)
		// return false;
		// }

		// boolean anyNonZero = Arrays.stream(charCount).anyMatch(count -> count != 0);
		// return !anyNonZero;

		return Arrays.stream(charCount).allMatch(count -> count == 0);

	}

	// 🧪 Test cases go here
	public static void main(String[] args) {
		String s = "anagram";
		String t = "nagaram";
		Solution solution = new Solution();
		boolean result = solution.isAnagram(s, t);
		System.out.println(result);
	}
}
