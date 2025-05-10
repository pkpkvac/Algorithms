import java.util.*;

public class ValidParentheses {

	public boolean isValid(String s) {
		Stack<Character> stack = new Stack<>();

		for (char c : s.toCharArray()) {
			// If opening bracket, push to stack
			if (c == '(' || c == '[' || c == '{') {
				stack.push(c);
			}
			// If closing bracket
			else if (c == ')' || c == ']' || c == '}') {
				// Check if stack is empty (no matching opening bracket)
				if (stack.isEmpty()) {
					return false;
				}

				char top = stack.pop();

				// Check if brackets match
				if ((c == ')' && top != '(') ||
						(c == ']' && top != '[') ||
						(c == '}' && top != '{')) {
					return false;
				}
			}
		}

		// Check if all brackets were matched (stack should be empty)
		return stack.isEmpty();
	}

	public static void main(String[] args) {

		ValidParentheses solution = new ValidParentheses();

		String s = "()[]{}";

		boolean result = solution.isValid(s);

		System.out.println(result);

	}

}
