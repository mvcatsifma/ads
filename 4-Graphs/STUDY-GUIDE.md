# Chapter 4: Graphs - Study Guide

## For Working Engineers

Graphs model relationships: networks, maps, dependencies, social connections. This chapter is immediately practical.

**Note:** You already have implementations in `4.1-Undirected-Graphs/` and `4.2-Directed-Graphs/`!

## The Big Picture
- **Section 4.1**: Undirected graphs - roads, networks, friendships
- **Section 4.2**: Directed graphs - dependencies, web links, state machines
- **Section 4.3**: Minimum spanning trees - network design, clustering
- **Section 4.4**: Shortest paths - GPS, routing, network flow

## Section 4.1: Undirected Graphs
**Essential - You have implementations!**

### What Is a Graph?
- **Vertices (nodes)**: The things
- **Edges**: Connections between things
- **Undirected**: Connection works both ways

### Representations

#### 1. Adjacency List (Standard)
```go
// Your implementation in 4.1-Undirected-Graphs/graph.go
type Graph struct {
    V   int        // number of vertices
    E   int        // number of edges
    Adj [][]int    // adjacency lists
}
```
- **Space**: O(V + E)
- **Check if edge exists**: O(degree(v))
- **Iterate over neighbors**: O(degree(v))
- **Best for sparse graphs** (most real graphs)

#### 2. Adjacency Matrix
- **Space**: O(V²)
- **Check if edge exists**: O(1)
- **Only good for dense graphs** (rare in practice)

### Core Algorithms (You Have These!)

#### Depth-First Search (DFS)
**File**: `dfs_paths.go`, `dfs_reachable.go`

**What it does:**
- Explores as far as possible before backtracking
- Uses recursion (or explicit stack)

**Applications:**
- Find path between two vertices
- Check connectivity
- Detect cycles
- Topological sort (for DAGs)

**Complexity:** O(V + E)

#### Breadth-First Search (BFS)
**File**: `bfs_paths.go`

**What it does:**
- Explores neighbors level by level
- Uses queue
- Finds **shortest path** in unweighted graph

**Applications:**
- Shortest path (unweighted)
- Level-order traversal
- Find all vertices within k edges
- Web crawling

**Complexity:** O(V + E)

#### Connected Components
**File**: `connected_components.go`

**What it does:**
- Groups vertices into separate components
- Uses DFS or union-find

**Applications:**
- Network connectivity
- Image segmentation
- Social network groups
- Finding islands

#### Cycle Detection
**File**: `cycle.go`

**What it does:**
- Detects if graph has a cycle
- Uses DFS

**Applications:**
- Deadlock detection
- Circular dependency checking

#### Bipartite Check
**File**: `bipartite.go`

**What it does:**
- Check if graph can be 2-colored
- Uses BFS or DFS

**Applications:**
- Matching problems
- Scheduling (no conflicts)
- Detecting odd cycles

### Real-World Applications
1. **Social Networks**: Friends, connections, communities
2. **Network Design**: Routers, connections
3. **Map Navigation**: Intersections and roads
4. **Dependency Resolution**: Build systems
5. **Compiler**: Control flow graphs

## Section 4.2: Directed Graphs (Digraphs)
**Essential - You have implementations!**

### Directed vs. Undirected
- Edge has direction: A → B (doesn't imply B → A)
- Represents: Dependencies, web links, inheritance, state transitions

### Your Implementations

#### Directed DFS
**Files**: `directed_dfs_paths.go`, `directed_dfs_reachable.go`
- Same as undirected but follow edge direction

#### Directed BFS
**File**: `directed_bfs_paths.go`
- Shortest directed path

#### Directed Cycle Detection
**File**: `directed_cycle.go`

**Critical for:**
- Deadlock detection
- Circular dependency detection
- Build system validation

#### DFS Orderings
**File**: `dfs_order.go`

Three orderings:
1. **Pre-order**: Record before recursion
2. **Post-order**: Record after recursion
3. **Reverse post-order**: Topological order for DAGs

#### Topological Sort
**File**: `topological.go`

**What it does:**
- Linear ordering where all edges point forward
- Only possible for DAGs (Directed Acyclic Graphs)
- Uses reverse post-order of DFS

**Applications:**
- Task scheduling (prerequisites)
- Course prerequisites
- Build systems (compile A before B)
- Spreadsheet formula evaluation
- Package dependency resolution

**Algorithm:**
1. Check for cycles (no topological order if cycle exists)
2. Do DFS, record vertices in reverse post-order
3. That's your topological order!

**Your implementation**: Returns (hasOrder, pre, post, reverse) where reverse is the topological order.

#### Strong Components (Kosaraju-Sharir)
**File**: `kosaraju_sharir_scc.go`

**What it does:**
- Finds strongly connected components
- Two vertices are strongly connected if there's a path in both directions

**Applications:**
- Web graph analysis (strongly connected web pages)
- Compiler optimization (variables always defined together)
- Ecology (food webs)

**Algorithm (clever!):**
1. Do DFS on graph, get reverse post-order
2. Reverse all edges
3. Do DFS on reversed graph in reverse post-order
4. Each DFS tree is a strong component

### Real-World Applications
1. **Package Managers**: Dependency resolution (npm, cargo, go mod)
2. **Build Systems**: Compile order (make, bazel)
3. **Version Control**: Commit graphs, merge detection
4. **Web Crawling**: Follow links, detect cycles
5. **Spreadsheets**: Cell dependencies

## Section 4.3: Minimum Spanning Trees
**Essential for network design**

### What Is MST?
Given weighted, undirected, connected graph:
- **Spanning tree**: Connects all vertices with V-1 edges
- **Minimum**: Sum of edge weights is minimum

### Why It Matters
**Problem**: Connect all vertices with minimum total cost
**Applications:**
- Network design (power grid, telecom)
- Approximation algorithms
- Cluster analysis
- Image segmentation

### Algorithms

#### Prim's Algorithm
**Idea**: Grow tree one edge at a time, always adding cheapest edge that connects to a new vertex

**Implementation**: Use priority queue for edges
**Complexity**: O(E log V) with binary heap

**When to use**: Dense graphs

#### Kruskal's Algorithm
**Idea**: Sort edges by weight, add edges if they don't create cycle

**Implementation**: Uses Union-Find! (You have this from Chapter 1)
**Complexity**: O(E log E)

**When to use**: Sparse graphs, already have union-find

### Practical Tips
- For most practical graphs (sparse), use Kruskal with union-find
- For dense graphs or online algorithms, use Prim
- Both give same MST (though may differ if edge weights equal)

## Section 4.4: Shortest Paths
**Essential - GPS, routing, network flow**

### Single-Source Shortest Paths
Given source vertex, find shortest paths to all other vertices

### Dijkstra's Algorithm
**Idea**: Like Prim's MST, but minimizing distance from source

**Requirements**: Non-negative edge weights

**Implementation**: Priority queue of vertices by distance
**Complexity**: O(E log V) with binary heap

**Applications:**
- GPS navigation
- Network routing (OSPF)
- Game AI pathfinding
- Flight planning

**When to use:**
- Non-negative weights
- Need shortest paths from one source
- Can use with A* heuristic for faster search

### Edge-Weighted DAGs
**Special case**: DAG (Directed Acyclic Graph)

**Algorithm**: Process vertices in topological order
**Complexity**: O(E + V) - linear!

**Applications:**
- Project scheduling (critical path)
- Job scheduling with dependencies
- PERT/CPM charts

### Bellman-Ford Algorithm
**Idea**: Relax all edges V times

**Advantage**: Handles negative weights
**Disadvantage**: Slower - O(VE)
**Detects**: Negative cycles

**When to use:**
- Have negative edge weights
- Need to detect negative cycles
- Can't use Dijkstra

### All-Pairs Shortest Paths
**Floyd-Warshall**: Find shortest paths between all pairs
**Complexity**: O(V³)

**When to use:**
- Dense graphs
- Need distances between all pairs
- Graph fits in memory

## Practical Decision Guide

### Graph Traversal
- **Need shortest path (unweighted)?** → BFS
- **Need any path?** → DFS
- **Need all reachable vertices?** → DFS or BFS

### Directed Graphs
- **Need process order?** → Topological sort
- **Check for cycles?** → Cycle detection
- **Find strongly connected groups?** → Kosaraju-Sharir

### Weighted Graphs
- **Minimum spanning tree?** → Kruskal (sparse) or Prim (dense)
- **Shortest paths, non-negative weights?** → Dijkstra
- **Shortest paths in DAG?** → Topological sort + relaxation
- **Negative weights?** → Bellman-Ford

## What to Skip
- Detailed proofs of correctness
- Every MST variant
- Network flow algorithms (unless you need them)
- Max flow/min cut (advanced topic)

## What to Focus On
- Understanding BFS vs. DFS (you have implementations!)
- Topological sort for DAGs
- Dijkstra's algorithm (most used shortest path)
- MST algorithms (Kruskal with union-find)
- When to use which algorithm

## Study Time Estimate
- Section 4.1: 4-5 hours (study your implementations, run tests)
- Section 4.2: 4-5 hours (directed algorithms, topological sort)
- Section 4.3: 3-4 hours (MST algorithms, understand Kruskal)
- Section 4.4: 4-5 hours (Dijkstra, shortest path variants)

**Total: 15-19 hours**

## Hands-On Learning
1. **Run your implementations**: Test the code in `4.1` and `4.2`
2. **Trace algorithms**: Walk through BFS/DFS on paper
3. **Solve problems**: LeetCode graph problems
4. **Implement Dijkstra**: Build on your existing graph code
5. **Implement Kruskal**: Use your union-find from Chapter 1!

## Real Projects to Try
1. **Route planner**: Dijkstra on a city graph
2. **Dependency resolver**: Topological sort
3. **Network analyzer**: Find connected components
4. **Build system**: Detect circular dependencies

## Next Steps
Chapter 5 (Strings) focuses on text processing - a different domain but builds on the graph/tree concepts you've learned.
