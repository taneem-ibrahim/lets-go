/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument
and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent
elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function
should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap()
function should return nothing, but it should swap the contents of the slice in position i with the contents
in position i+1.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputNumber := ""
	digits := make([]int, 0)
	for {
		fmt.Printf("Enter a digit to sort or type X when done entering:")
		fmt.Scan(&inputNumber)
		if strings.EqualFold(strings.TrimSpace(inputNumber), "X") {
			break
		} else {
			inputDigit, _ := strconv.Atoi(inputNumber)
			digits = append(digits, inputDigit)
		}
	}
	BubbleSort(digits)
	fmt.Printf("%d", digits)
}

// it is fine for this function be void since we are not adding a new element to the underlying array just updating the
// underlying array's elements. If we were actually modifying the length of the array then we would want to pass
// by reference so that they original underlying array and slice pointing to it is updated as well (and save on copy cost)
// using the swapped flag to save unnecessary comparisons
func BubbleSort(digits []int) {
	for i := 0; i < len(digits); i++ {
		swapped := false
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				Swap(digits, i, j)
				swapped = true
			}
		}
		// if no swap occurred in the inner loop then all the elements are already sorted so we can break the outer loop
		if !swapped {
			break
		}
	}
}

func Swap(digits []int, i int, j int) {
	tmp := digits[i]
	digits[i] = digits[j]
	digits[j] = tmp
}
