package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your
program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

type names struct {
	fname string
	lname string
}

func main() {
	//declare a slice for the struct to store data in an array of structs format
	var sli []names
	fmt.Printf("Enter a file name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error reading the file!")
	}
	// this is a pretty cool feature: Immediately after getting a file object with createFile, we defer the closing of
	// that file with close. This will be executed at the end of the enclosing function (main),
	// after writeFile has finished. Ref: https://gobyexample.com/defer
	defer file.Close()
	if err == nil {
		//reads each line in a file
		scanner = bufio.NewScanner(file)
		//this moves the file cursor to the next line
		for scanner.Scan() {
			line := scanner.Text()
			fileContent := strings.Split(line, " ")
			fname, lname := fileContent[0], fileContent[1]
			name := names{fname, lname}
			sli = append(sli, name)
		}
		// iterate over the slice of structs and print them ... go uses for range to iterate slice (kind of like python)
		// range rteurns two things: first one is the index of the loop and second one is the actual value. Since I do
		// not care about the index I put _ so that an unnecessary variable is not allocated. This is also somewhat
		// similar to the _ concept in python. Also note the value return is a struct itself.
		for _, sliceNames := range sli {
			fmt.Println(sliceNames.fname, sliceNames.lname)
		}
	}

}
