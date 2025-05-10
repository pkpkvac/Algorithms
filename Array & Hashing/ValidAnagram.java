import java.sql.Array;
import java.util.*;

public class ValidAnagram {

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

	public boolean isAnagramHashMap(String s, String t) {

		if (s.length() != t.length())
			return false;

		HashMap<Character, Integer> map = new HashMap<>();

		for (char c : s.toCharArray()) {
			map.put(c, map.getOrDefault(c, 0) + 1);
		}

		for (char c : t.toCharArray()) {
			if (!map.containsKey(c) || map.get(c) == 0)
				return false;

			map.put(c, map.get(c) - 1);
		}

		for (int count : map.values()) {
			if (count != 0)
				return false;
		}

		return true;
	}

	// 🧪 Test cases go here
	public static void main(String[] args) {
		String s = "anagram";
		String t = "nagaram";
		ValidAnagram solution = new ValidAnagram();
		boolean result = solution.isAnagram(s, t);
		System.out.println(result);
	}
}
