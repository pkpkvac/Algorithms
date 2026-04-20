# go-craft

Go solutions for algorithm and data structure problems, organized by category.

## Repo Structure

```
go-craft/
  <category>/
    <problem_name>/
      README.md         # problem statement (plain text)
      solution.go       # solution with function stub or implementation
      solution_test.go  # table-driven tests using examples from the problem
      complexity.md     # complexity analysis (required for completion)
```

## Categories

| Directory            | Topic                        |
|----------------------|------------------------------|
| graphs               | BFS, DFS, union-find         |
| tree                 | binary trees, BST, traversal |
| greedy               | greedy algorithms            |
| dp / 1d_dp           | dynamic programming          |
| backtracking         | backtracking                 |
| heap_priority_queue  | heaps                        |
| binary_search        | binary search                |
| sliding_window       | sliding window               |
| two_pointers         | two pointer technique        |
| stack                | stack problems               |
| linked_list          | linked list problems         |
| array_hashing        | arrays, hash maps            |
| intervals            | interval problems            |

## Naming Conventions

- Directory: `snake_case` of problem name
- Go package: same as directory name
- Go function: `camelCase` matching LeetCode's function signature

## Go Module

Module name: `algorithms` (see go.mod)

## Solution Completion Requirements

A solution is **not complete** until it includes a `complexity.md` file with:

1. **Problem approach**: Brief description of the algorithm/technique used
2. **Time complexity**:
   - Best case
   - Average case
   - Worst case
3. **Space complexity**: Analysis of auxiliary space used
4. **Optimizations**: Discussion of alternative approaches and trade-offs

When running tests successfully, always create or update `complexity.md` before considering the solution complete.

## Skills

When asked to set up a new problem, follow `.claude/skills/algo-problem-setup/SKILL.md`.
