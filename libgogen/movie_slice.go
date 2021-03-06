// Generated by: main
// TypeWriter: slice
// Directive: +gen on Movie

package libgogen

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// MovieSlice is a slice of type Movie. Use it where you would use []Movie.
type MovieSlice []Movie

// GroupByInt groups elements into a map keyed by int. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv MovieSlice) GroupByInt(fn func(Movie) int) map[int]MovieSlice {
	result := make(map[int]MovieSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// SortBy returns a new ordered MovieSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv MovieSlice) SortBy(less func(Movie, Movie) bool) MovieSlice {
	result := make(MovieSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortMovieSlice(result, less, 0, n, maxDepth)
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapMovieSlice(rcv MovieSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapMovieSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapMovieSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownMovieSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapMovieSlice(rcv, first, first+i)
		siftDownMovieSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapMovieSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapMovieSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapMovieSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeMovieSlice(rcv MovieSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapMovieSlice(rcv, a+i, b+i)
	}
}

func doPivotMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeMovieSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeMovieSlice(rcv, less, m, m-s, m+s)
		medianOfThreeMovieSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeMovieSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapMovieSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapMovieSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapMovieSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeMovieSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeMovieSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortMovieSlice(rcv MovieSlice, less func(Movie, Movie) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortMovieSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotMovieSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortMovieSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortMovieSlice(rcv, mhi, b)
		} else {
			quickSortMovieSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortMovieSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortMovieSlice(rcv, less, a, b)
	}
}
