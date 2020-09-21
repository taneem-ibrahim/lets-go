/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array
should print the subarray that it will sort. When sorting is complete, the main goroutine
should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//wrapper around strconv.AtoI() to handle the error properly
func atoi(s string) int {
	digit, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return digit
}

func appendSlice(numbers []string) []int {
	sli := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		sli = append(sli, atoi(numbers[i]))
	}
	return sli
}

func sortPart1(slice1ptr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*slice1ptr)
	fmt.Println("Goroutine Part 1 sorted elements: ", *slice1ptr)
}
func sortPart2(slice2ptr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*slice2ptr)
	fmt.Println("Goroutine Part 2 sorted elements: ", *slice2ptr)
}
func sortPart3(slice3ptr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*slice3ptr)
	fmt.Println("Goroutine Part 3 sorted elements: ", *slice3ptr)
}
func sortPart4(slice4ptr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*slice4ptr)
	fmt.Println("Goroutine Part 4 sorted elements: ", *slice4ptr)
}

func main() {
	finalSortedSlice := make([]int, 0)
	fmt.Println("Enter a list of integers separated by comma")
	fmt.Printf(">")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	sli := appendSlice(numbers)
	// partition the numbers into 4 equal sub arrays
	mid := int(math.Ceil(float64(len(sli) / 2)))
	firstHalf := sli[0:mid]
	secondHalf := sli[mid:len(sli)]
	firstHalfMid := int(math.Ceil(float64(len(firstHalf) / 2)))
	secondHalfMid := int(math.Ceil(float64(len(secondHalf) / 2)))
	slice1 := firstHalf[0:firstHalfMid]
	slice2 := firstHalf[firstHalfMid:len(firstHalf)]
	slice3 := secondHalf[0:secondHalfMid]
	slice4 := secondHalf[secondHalfMid:len(secondHalf)]
	//sort the 4 parts concurrently
	var wg sync.WaitGroup
	wg.Add(1)
	go sortPart1(&slice1, &wg)
	wg.Add(1)
	go sortPart2(&slice2, &wg)
	wg.Add(1)
	go sortPart3(&slice3, &wg)
	wg.Add(1)
	go sortPart4(&slice4, &wg)
	wg.Wait()
	//To avoid "cannot use []int literal (type []int) as type int in append" error
	//just like any other variadic function we have to add ... at the end of the second argument
	//reference: https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
	finalSortedSlice = append(finalSortedSlice, slice1...)
	finalSortedSlice = append(finalSortedSlice, slice2...)
	finalSortedSlice = append(finalSortedSlice, slice3...)
	finalSortedSlice = append(finalSortedSlice, slice4...)
	sort.Ints(finalSortedSlice)
	fmt.Println("Fully sorted list: ", finalSortedSlice)
}
