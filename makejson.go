package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

func main() {
	var name string
	var address string
	//map to hold the struct values to marshal into json
	personMap := make(map[string]string)

	fmt.Printf("Enter a name:")
	fmt.Scan(&name)
	fmt.Printf("Enter an address:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	address = scanner.Text()
	// store the data into the map
	personMap[name] = address
	// marshal to json
	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Printf("Error:", err)
	} else {
		fmt.Printf(string(jsonData))
	}

}
