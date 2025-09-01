## 160. Intersection of Two Linked Lists

https://leetcode.com/problems/intersection-of-two-linked-lists/description/

## Benchmark results

```
goos: darwin
goarch: arm64
pkg: leetcode/160-Intersection-of-Two-Linked-Lists
cpu: Apple M3 Pro
BenchmarkGetIntersectionNode
BenchmarkGetIntersectionNode/Small
BenchmarkGetIntersectionNode/Small-12         	75822302	        15.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetIntersectionNode/Medium
BenchmarkGetIntersectionNode/Medium-12        	 3049366	       403.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetIntersectionNode/Large
BenchmarkGetIntersectionNode/Large-12         	  276906	      4315 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetIntersectionNode/XLarge
BenchmarkGetIntersectionNode/XLarge-12        	   22990	     52723 ns/op	       0 B/op	       0 allocs/op
PASS
```

The provided benchmark results for the `getIntersectionNode` function in the LeetCode 160 problem show the performance characteristics of the solution across different input sizes. Let's analyze the key takeaways:

1. **Time Complexity**: The benchmark results demonstrate that the time complexity of the solution is linear with respect to the input size. As the size of the linked lists increases, the time taken per operation also increases linearly.

    - For the "Small" case (10 and 15 nodes), the time per operation is around 15.99 ns.
    - For the "Medium" case (100 and 150 nodes), the time per operation is around 403.1 ns.
    - For the "Large" case (1,000 and 1,500 nodes), the time per operation is around 4,315 ns.
    - For the "XLarge" case (10,000 and 15,000 nodes), the time per operation is around 52,723 ns.

2. **Space Complexity**: The solution has a constant space complexity, as indicated by the 0 B/op and 0 allocs/op results across all input sizes. This means the solution does not allocate any additional memory proportional to the input size.

3. **Efficiency**: The solution is highly efficient, with very low time and space requirements. Even for the "XLarge" case with 10,000 and 15,000 nodes, the time per operation is only around 52 microseconds, which is well within acceptable performance bounds for most real-world applications.

4. **Scalability**: The linear time complexity of the solution suggests that it can handle even larger input sizes without significant performance degradation. The solution should be able to scale well as the size of the linked lists increases.

5. **Consistency**: The benchmark results show consistent performance across the different input sizes, with no significant outliers or unexpected behavior. This indicates a robust and reliable implementation.

Overall, the benchmark results demonstrate that the provided `getIntersectionNode` solution is highly efficient, scalable, and well-optimized for the LeetCode 160 problem. The linear time complexity and constant space complexity make it a suitable choice for real-world applications dealing with linked list intersection problems.
