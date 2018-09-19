# Double booked

This solution is designed and implemented by Xindi Wang

## Brief Introduction

1. I started off with bruce force where compare each interval with each other and the time complexity is O(n2) -> n square.
2. Optimization: applying interval tree would reduce the average time complexity to O(nlogn) -> each interval only takes logn time to search on the interval tree.
3. Note that insertion function is implemented in regular BST but not red-black tree. This makes the worse case time complexity still O(n2) -> n square.

## How to run the code
* `go build` under double-booked folder. Make sure that you have golang installed on your computer.
* `go run itvtree.go`
* then you will see the result, for example:
```
given sequence: {[1, 2], [1, 3], [4, 5], [3, 6], {2, 8}, {10, 100}
for every single event, the output would list out all of the overlapped events, as shown below:
Event {1 2} overlaps with event(s) : [{1 3}]
Event {1 3} overlaps with event(s) : [{1 2} {2 8}]
Event {4 5} overlaps with event(s) : [{3 6} {2 8}]
Event {3 6} overlaps with event(s) : [{4 5} {2 8}]
Event {2 8} overlaps with event(s) : [{1 3} {4 5} {3 6}]
```
