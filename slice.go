package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter
an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints
the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers
which the user decides to enter. The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer.
*/

func main() {
	inputNumber := ""
	// create a slice object of len=0 which is a data type and used more widely than arrays in go
	// The slice grows as needed and append works on nil (ie NULL) slices too.
	// the brackets [] without ... or data element size makes it a slice object
	sli := make([]int, 0)
	for {
		fmt.Printf("Enter an integer or enter 'X' to exit:")
		fmt.Scan(&inputNumber)
		// make it case insensitive
		if strings.EqualFold(strings.TrimSpace(inputNumber), "X") {
			os.Exit(3)
		} else {
			// interesting that Atoi returns two values so we cant directly append to the slide of int[]
			input, err := strconv.Atoi(inputNumber)
			//Note that we need to accept a return value from append as we may get a new slice value
			//we can even multiple values to a slice as well like append(sli, 2, 3, 4)
			if err == nil {
				sli = append(sli, input)
				sort.Ints(sli)
				fmt.Printf("Numbers stored (sorted): %v\n", sli)
			}
		}

	}
}
