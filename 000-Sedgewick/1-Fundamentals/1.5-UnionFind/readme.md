Here you go — clean, structured, markdown-ready.

---

# Union-Find (Disjoint Set Union) — Invariants

Understanding Union-Find means understanding its invariants.
Everything else is implementation detail.

---

## Core Structural Invariant

At any moment:

> The data structure represents a **forest of disjoint trees**, where each tree corresponds to one connected component (set).

Each tree:

* Has exactly one root
* The root is its own parent
* All nodes eventually point (directly or indirectly) to that root

Formally:

```text
For every node x:
Following parent[x] repeatedly must eventually reach a node r
such that parent[r] == r.
```

If this is ever false, the structure is corrupted.

---

## 1. Parent Structure Invariant

For every node `x`:

```text
find(x) always returns a root r such that parent[r] == r.
```

Implications:

* No cycles
* Trees remain valid
* Roots are stable representatives

---

## 2. Representative Invariant

For any elements `a` and `b`:

```text
find(a) == find(b)  <=>  a and b belong to the same set
```

This is the correctness guarantee of the structure.

Everything — union, path compression, rank — must preserve this.

---

## 3. Union Invariant

During `union(x, y)`:

1. Find roots:

   ```text
   rx = find(x)
   ry = find(y)
   ```

2. Only roots may be attached to other roots:

   ```text
   parent[rx] = ry
   ```

Golden rule:

> Only roots may become children of other roots.

Never attach a non-root directly.
That risks breaking tree structure and invariants.

---

## 4. Rank / Size Invariant

### Union by Rank

```text
rank[root] is an upper bound on the height of the tree.
```

When merging:

* Attach smaller-rank tree under larger-rank tree
* If ranks equal, increment new root's rank

This guarantees near-constant amortized time.

---

### Union by Size

```text
size[root] correctly tracks number of elements in the component.
```

When merging:

* Attach smaller tree under larger tree
* Update size of new root

Correctness depends on updating size only at roots.

---

## 5. Path Compression Invariant

During `find(x)`:

```text
parent[x] is updated to point closer to the root.
```

Key guarantee:

> Path compression never changes connectivity.

It only shortens paths.
It never merges unrelated sets.

---

## Correctness Summary

At all times:

1. The structure forms a forest.
2. Each tree represents one equivalence class.
3. Two nodes are connected if and only if they share the same root.
4. Union operations merge trees.
5. Find operations preserve connectivity while flattening trees.

---

## Deep View (What You're Actually Maintaining)

Union-Find maintains an **equivalence relation**:

* Reflexive
* Symmetric
* Transitive

The invariant is:

> The forest structure correctly encodes this equivalence relation at all times.

---

## Where Invariants Commonly Break

Typical implementation mistakes:

* Attaching non-root nodes during union
* Forgetting to find roots before union
* Incorrect rank updates
* Updating size on non-root nodes
* Incorrect path compression recursion

When debugging:

Ask:

```text
Which invariant did I just violate?
```

Not:

```text
Why doesn't this work?
```

That shift is professional-level thinking.

---

If you want, next step:
We can write a fully annotated Union-Find implementation with invariants explicitly stated above each function.
