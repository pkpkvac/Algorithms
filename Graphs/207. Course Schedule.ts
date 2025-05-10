// Indegree is very common for:
// Dependency resolution
// Build systems
// Task scheduling
// Any "topological sort" problem

function canFinishDFS(numCourses: number, prerequisites: number[][]): boolean {
	return false;
}

function canFinishTopo(numCourses: number, prerequisites: number[][]): boolean {
	return false;
}

const numCourses = 5;
const prerequisites = [
	[1, 4],
	[2, 4],
	[3, 1],
	[3, 2],
];

// Joint (Connected AND Directed):
//     4
//    ↙ ↘
//   1   2   All courses connected
//    ↘ ↙    through dependencies
//     3

// Cyclical (Creates Impossible Schedule):
// 1 → 2
// ↑   ↓     Can't start anywhere!
// 3 ←       (Impossible to complete)

// Key Concerns:
// Direction matters (prerequisites)
// Cycles make schedule impossible
// Can have multiple independent groups
// Need to handle all these cases!
// This is why it's a graph problem!

function canFinishMapAndQueue(
	numCourses: number,
	prerequisites: number[][]
): boolean {
	// Validate course numbers
	for (const [course, prereq] of prerequisites) {
		if (course >= numCourses || prereq >= numCourses) {
			return false; // Invalid course number
		}
	}

	const prereqs = new Map<number, number[]>();

	// Initialize all courses (including isolated ones)
	for (let i = 0; i < numCourses; i++) {
		prereqs.set(i, []);
	}

	// Build adjacency list
	// add courses with prereqs
	for (const [course, prereq] of prerequisites) {
		if (!prereqs.has(course)) {
			prereqs.set(course, []);
		}
		prereqs.get(course)!.push(prereq);
	}

	// add all prereq courses with empty arrays (tricky).
	for (const [_, prereq] of prerequisites) {
		if (!prereqs.has(prereq)) {
			prereqs.set(prereq, []);
		}
	}

	// find initial courses with no prerequisites
	const queue: number[] = [];
	for (const [course, deps] of prereqs) {
		if (deps.length === 0) queue.push(course);
	}

	let coursesCompleted = 0;

	while (queue.length > 0) {
		const prereq = queue.shift();
		if (prereq === undefined) continue;

		coursesCompleted++;

		for (const [course, deps] of prereqs) {
			if (deps.includes(prereq)) {
				// remove the prereq from dep, and add it to the queue if dep.length is 0
				const newDeps = deps.filter((val) => val !== prereq);
				prereqs.set(course, newDeps);

				if (newDeps.length === 0 && !queue.includes(course)) {
					queue.push(course);
				}
			}
		}
	}

	for (const [_, deps] of prereqs) {
		if (deps.length > 0) {
			return false;
		}
	}

	return coursesCompleted === numCourses;
}
console.log(canFinishMapAndQueue(numCourses, prerequisites));
