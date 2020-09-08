package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the
characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”,
“I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

const (
	Prefix          = "i"
	Suffix          = "n"
	SearchCharacter = "a"
)

func main() {
	fmt.Printf("Enter a string:")
	// code reference: https://stackoverflow.com/questions/27414598/golang-accepting-input-with-spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if find_string(strings.ToLower(input)) {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}

func find_string(str string) bool {
	if strings.HasPrefix(str, Prefix) && strings.HasSuffix(str, Suffix) {
		return strings.Contains(str, SearchCharacter)
	} else {
		return false
	}

}
