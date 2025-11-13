For general tree problems:

Pre-order: tree creation, copying, serialization

Post-order: tree deletion, evaluation, most manipulations

In-order: mainly BST operations (sorted order)

Post-order: process left and right first, then swap at current. Since the subtrees are already
inverted, you're swapping already-inverted subtrees.

Pre-order: swap first, then recurse. After swapping, recurse on what are now the swapped children.
