// We are given an array asteroids of integers representing asteroids in
// a row. The indices of the asteriod in the array represent their relative position in space.

// For each asteroid, the absolute value represents its size,
//  and the sign represents its direction (positive meaning right,
// 	 negative meaning left). Each asteroid moves at the same speed.

// Find out the state of the asteroids after all collisions.
//  If two asteroids meet, the smaller one will explode.
//   If both are the same size, both will explode.
// 	 Two asteroids moving in the same direction will never meet.

// Example 1:

// Input: asteroids = [5,10,-5]
// Output: [5,10]
// Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
// Example 2:

// Input: asteroids = [8,-8]
// Output: []
// Explanation: The 8 and -8 collide exploding each other.
// Example 3:

// Input: asteroids = [10,2,-5]
// Output: [10]
// Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

function asteroidCollision(asteroids: number[]): number[] {
	const stack: number[] = [];

	for (const asteroid of asteroids) {
		// Flag to track if current asteroid survives
		let survives = true;

		// Handle collisions when current asteroid is moving left
		// and there are asteroids in the stack moving right
		while (
			survives &&
			asteroid < 0 &&
			stack.length > 0 &&
			stack[stack.length - 1] > 0
		) {
			// Current asteroid is moving left, top of stack is moving right
			const topAsteroid = stack[stack.length - 1];

			// Compare sizes
			if (Math.abs(topAsteroid) < Math.abs(asteroid)) {
				// Top asteroid is smaller, it explodes
				stack.pop();
			} else if (Math.abs(topAsteroid) === Math.abs(asteroid)) {
				// Equal size, both explode
				stack.pop();
				survives = false;
			} else {
				// Current asteroid is smaller, it explodes
				survives = false;
			}
		}

		// If current asteroid survives, add it to the stack
		if (survives) {
			stack.push(asteroid);
		}
	}

	return stack;
}

// const asteroids = [10, 2, -5];
const asteroids = [5, 7, -8, 5];

console.log(asteroidCollision(asteroids));
