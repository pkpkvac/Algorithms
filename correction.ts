// for an array of numbers, some with positive values, some with negative values,
// take the positive values, and return 5 random positive values from the array,
// there is a function randInt(x,y) that returns a random integer between x and y

function randInt(x: number, y: number): number {
	// unspecified, returns a random integer (in this case, the index of the array) between x and y
	// example implementation:
	return Math.floor(Math.random() * (y - x + 1)) + x;
}

function getRandomPositiveValues(arr: number[], count: number): number[] {
	const positiveValues = arr.filter((num) => num > 0);
	const randomPositiveValues: number[] = [];
	for (let i = 0; i < count; i++) {
		const randomIndex = randInt(0, positiveValues.length - 1);
		randomPositiveValues.push(positiveValues[randomIndex]);
	}
	return randomPositiveValues;
}

// example usage:
const arr = [
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10,
];
const randomPositiveValues = getRandomPositiveValues(arr, 5);
console.log(randomPositiveValues);
