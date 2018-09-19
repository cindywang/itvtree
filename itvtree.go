/*
This solution of double booked problem is written by Xindi Wang.
Reference(s):
Segment tree data structure: Introduction to Algorithms by T. Cormen et al.
*/
package main

import "fmt"

//A struct to represent the interval in each node
type Interval struct {
	low    int
  high   int
}

//A struct to represent a node in interval search tree
type ITVNode struct {
	interval    Interval
	max         int
	left, right *ITVNode
}

//utility func to initialize a new interval search tree node
func NewITVNode(i Interval) *ITVNode {
	return &ITVNode{
		interval: i,
		max:   i.high,
		left:  nil,
		right: nil,
	}
}

//utility func to insert a new interval search tree node
//low value in the interval is the key to maintain the tree
func insert(root *ITVNode, i Interval) *ITVNode {
	//base case empty tree
	if root == nil {
		return NewITVNode(i)
	}

	l := root.interval.low

	if i.low < l {
		root.left = insert(root.left, i)
	} else {
		root.right = insert(root.right, i)
	}

	if root.max < i.high {
		root.max = i.high
	}

	return root
}

//check if two intervals conflict with each other.
//if two intervals overlaps in anykind (doesn't matter inclusive or exclusive)
// they are conflict events with each other
//e.g [1,2] and [2,3] / [2, 5] and [3, 6]
func checkOverlap(l1, l2 Interval) bool {
	if l1.low == l2.low && l1.high == l2.high {
		return false
	} else if l1.low < l2.high && l2.low < l1.high {
		return true
	}
	return false
}

//seach for overlapping intervals
//conlict events will be appended to the end of result slice
// and pass along in the recursion
func search(res []Interval, root *ITVNode, i Interval) []Interval {
	if root == nil {
		return res
	}

	if checkOverlap(root.interval, i) {
		res = append(res, root.interval)
	}

	//condition to go search on left subtree
	//left child is not nil and max value of left tree is greater or equal to given interval
	if root.left != nil && root.left.max >= i.low {
		return search(res, root.left, i)
	}
	return search(res, root.right, i)
}

func main() {
	x := []Interval{
    Interval{low: 1, high: 2},
		Interval{low: 1, high: 3},
    Interval{low:4, high: 5},
		Interval{low:3, high: 6},
		Interval{low:2, high: 8},
		Interval{low:10, high: 100},
  }
  var root *ITVNode
	res := make(map[Interval][]Interval)

	//insert into trees
	for _, itv := range x {
		root = insert(root, itv)
	}

	//search one by one
	for _, itv := range x {
		s := make([]Interval, 0)
		res[itv] = search(s, root, itv)
		fmt.Printf("Event %v overlaps with event(s) : %v \n", itv, res[itv])
	}
}
