## What Makes a Problem "Dynamic Programming"?

**DP solves decision/optimization questions**, not enumeration:

- ✅ "Can this be done?" (yes/no) — e.g., Word Break
- ✅ "What's the minimum/maximum way?" (value) — e.g., Coin Change, House Robber
- ✅ "How many ways?" (count) — e.g., Climbing Stairs
- ❌ "Give me all combinations" — that's **Backtracking**

**Key characteristics:**

1. **Overlapping subproblems** — same subproblem appears multiple times
2. **Optimal substructure** — optimal solution contains optimal solutions to subproblems
3. **Memoization helps** — caching subproblem results avoids redundant work

**DP vs. Backtracking:**

- **Backtracking**: Explores all possibilities, builds actual solutions (Permutations, Subsets, Combination Sum)
- **DP**: Answers questions about possibilities without building them (Word Break, Coin Change, Climbing Stairs)

---

## Learning Path

climbing stairs is a kind of canonical problem
for understanding:
a) recursion
b) recursion with memoization (top down)
c) tabulation (iterative, bottom up)

reference it for DP.

House Robber → Coin Change (or Min Cost Climbing Stairs if you have it)
