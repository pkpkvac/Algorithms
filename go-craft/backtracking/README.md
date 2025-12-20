def backtrack(state):
if is_valid_solution(state): # 1. Base Case
solutions.append(state.copy()) # Save solution
return

    for choice in get_possible_choices(state):
        # 2. Make a choice
        state.add(choice)

        # 3. Recurse with new state
        backtrack(state)

        # 4. Undo choice (backtrack)
        state.remove(choice)

The Core Pattern

Think of backtracking as exploring a decision tree:

1. At each step, you have choices (include this element? skip it? what value to
   try?)
2. Make a choice (modify state)
3. Recurse deeper (explore that path)
4. Undo the choice (backtrack - restore state)
5. Try the next choice

The Template Shape

function backtrack(state, choices):
if (base case - no more decisions):
record/process result
return

      for each choice in available choices:
          make choice (modify state)
          backtrack(new_state, remaining_choices)
          undo choice (restore state)

Key Questions to Ask

When you see a problem:

- What are my choices at each step? (include/exclude, pick a number, place
  something)
- What's my state that I'm building? (current subset, current path, current
  arrangement)
- When do I stop? (processed all elements, reached target, filled all positions)
- Do I need to undo changes? (if using mutable state that's shared across recursive
  calls)

Common Backtracking Flavors

- Subsets/Combinations: include or exclude each element
- Permutations: which unused element goes next
- Path finding: which neighbor to visit next
- Constraint satisfaction: which valid value fits here

That's it. The pattern is simpler than it seems - just systematically try all
possibilities by making choices, exploring, and undoing.
