// In this problem, a tree is an undirected graph that is connected and has no cycles.

// You are given a graph that started as a tree with n nodes labeled from 1 to n,
//  with one additional edge added. The added edge has two different vertices chosen
//   from 1 to n, and was not an edge that already existed. The graph is represented
// 	as an array edges of length n where edges[i] = [ai, bi] indicates that there is an
// 	edge between nodes ai and bi in the graph.

// Return an edge that can be removed so that the resulting graph is a tree of n nodes.
// If there are multiple answers, return the answer that occurs last in the input.

// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

// Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// Output: [1,4]

// Properties we need:
// 1. Tree = Connected + No cycles
// 2. n nodes = 5 nodes
// 3. Return last edge that creates cycle

function findRedundantConnectionBruteForce(edges: number[][]): number[] {
	const graph = new Map<number, number[]>();

	for (const [from, to] of edges) {
		// before adding the edge, check if the nodes are already connected
		// if there exists a path, there is a cycle, this is redundant
		if (hasPath(graph, to, from)) {
			return [from, to];
		}

		// otherwise extend and create the graph, we will get the last edge eventually
		if (!graph.has(from)) graph.set(from, []);
		if (!graph.has(to)) graph.set(to, []);

		graph.get(from)?.push(to);
		graph.get(to)?.push(from);
	}

	return [];
}

function hasPath(graph: Map<number, number[]>, source: number, target: number) {
	const visited = new Set<number>();
	const path = new Set<number>();

	function dfs(current: number, parent: number): boolean {
		visited.add(current);
		path.add(current);

		for (const neighbor of graph.get(current) || []) {
			if (neighbor === parent) continue;

			if (neighbor === target) return true; // i.e, there is an existing path

			if (!visited.has(neighbor)) {
				if (dfs(neighbor, current)) {
					return true;
				}
			}
		}

		path.delete(current);
		return false;
	}

	return dfs(source, -1);
}
