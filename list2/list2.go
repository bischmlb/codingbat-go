package list2

import (
	"sort"
)

// Return the number of even ints in the given array. Note: the % "mod" operator
// computes the remainder, e.g. 5 % 2 is 1.
func CountEvens(a []int) int {
	evenCount := 0
	for _, k := range a {
		if k%2 == 0 {
			evenCount++
		}
	}
	return evenCount
}

// Return the sum of the numbers in the array, returning 0 for an empty array.
// Except the number 13 is very unlucky, so it does not count and numbers that
// come immediately after a 13 also do not count.
func Sum13(a []int) int {
	/* lidt snyd */
	/* apparently there is not "null" or "undefined" value initially when assign with "var index13 int". It will be initialized to 0 (always lowest value of type)*/
	index13 := 99
	sum := 0
	if len(a) != 0 {
		for i, k := range a {
			if k != 13 && i != (index13+1) {
				sum += k
			} else if k == 13 {
				index13 = i
			}
		}
	}
	return sum
}

// Given an array length 1 or more of ints, return the difference between the
// largest and smallest values in the array.
func BigDiff(a []int) int {
	lowest, largest := a[0], a[0]
	for _, k := range a {
		if k > largest {
			largest = k
		}
		if k < lowest {
			lowest = k
		}
	}
	return largest - lowest
}

// Return the sum of the numbers in the array, except ignore sections of numbers
//starting with a 6 and extending to the next 7 (every 6 will be followed by at
// least one 7). Return 0 for no numbers.
func Sum67(a []int) int {
	seen6 := false
	sum := 0
	for _, k := range a {
		if seen6 != true && k != 6 {
			sum += k
		} else if k == 6 {
			seen6 = true
		} else if k == 7 && seen6 == true {
			seen6 = false
		}
	}
	return sum
}

// Return the "centered" average of an array of ints, which we'll say is the mean
// average of the values, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and
// likewise for the largest value. Use int division to produce the final average.
// You may assume that the array is length 3 or more.
func CenteredAverage(a []int) int {

	/* Elegant way (use sort)*/
	sort.Ints(a)
	trimmed := a[1 : len(a)-1]

	sum := 0
	for _, i := range trimmed {
		sum += i
	}
	return sum / len(trimmed)

	/* INITIAL WAY I USED (not so elegant)

	lowest, largest := a[0], a[0]
	lowestIndex, largestIndex := 0, 0
	skips := 0
	sum := 0

	for i := 0; i < len(a); i++ {
		if a[i] > largest {
			largest = a[i]
			largestIndex = i

		}
		if a[i] < lowest {
			lowest = a[i]
			lowestIndex = i
		}
	}

	for i, k := range a {
		if i != largestIndex && i != lowestIndex {
			sum += k
		} else {
			skips++
		}
	}

	if skips != 2 {
		sum -= largest
		return sum / (len(a) - 2)
	}
	return sum / (len(a) - 2)

	*/

}

// Given an array of ints, return true if the array contains a 2 next to a 2 somewhere.
func Has22(a []int) bool {
	for i, k := range a {
		if k == 2 {
			if !(i+1 > len(a)-1) {
				if a[i+1] == 2 {
					return true
				}
			}
		}
	}

	return false
}
