package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * Write a program which prompts the user to enter a floating point number and prints the integer which is a
 * truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 */

func main() {
	var inputNumber float64
	fmt.Printf("Enter a floating point number:")
	fmt.Scan(&inputNumber)
	// code reference: https://golang.org/pkg/math/#example_Trunc
	var truncatedNumber = math.Trunc(inputNumber)
	// code reference: https://golang.org/pkg/strconv/#:~:text=FormatFloat%20converts%20the%20floating%2Dpoint,%2C%20'e'%20(%2Dd.
	fmt.Printf(strconv.FormatFloat(truncatedNumber, 'f', -1, 64))
}
