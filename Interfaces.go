package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

// implement the interface methods
// notice the method receiver type is set to the struct type which associates the struct to the Animal interface
// in Go we don't have to explicitly say implements Some interface. The compiler does it behind the scene
// Also for ease of understanding think of the structs here as like classes in Java

func (cow Cow) Eat() {
	fmt.Println("I eat grass.")
}

func (cow Cow) Move() {
	fmt.Println("I walk.")
}

func (cow Cow) Speak() {
	fmt.Println("I make moo sound.")
}

type Bird struct {
	name string
}

func (bird Bird) Eat() {
	fmt.Println("I eat worms.")
}

func (bird Bird) Move() {
	fmt.Println("I fly.")
}

func (bird Bird) Speak() {
	fmt.Println("I make peep sound.")
}

type Snake struct {
	name string
}

func (snake Snake) Eat() {
	fmt.Println("I eat mice.")
}

func (snake Snake) Move() {
	fmt.Println("I slither.")
}

func (snake Snake) Speak() {
	fmt.Println("I make hsss sound.")
}

func processQuery(command string, namesMap map[string]string) {
	var animal Animal
	arguments := strings.Split(command, " ")
	name, request := arguments[1], arguments[2]
	//look up the animal type by name from the map
	animalType := namesMap[name]

	switch animalType {
	case "cow":
		//using the interface here to show polymorphism since Cow implements/satisfies Animal interface,
		// we can assign the generic animal type here
		//Animal = concrete type
		//Cow = dynamic type
		//name = dynamic value
		animal = Cow{name}
	case "bird":
		animal = Bird{name}
	case "snake":
		animal = Snake{name}
	}
	// since we are just calling the methods from the interface, we dont have to specify the specific type by name
	switch request {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {

	//map animal names to the associated type
	animalNames := make(map[string]string)

	for {
		fmt.Println("Enter a command: [newanimal name type] or [query name info] or type Q to exit")
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()
		if strings.EqualFold(strings.TrimSpace(command), "Q") {
			break
		} else if strings.HasPrefix(command, "query") {
			if len(strings.Split(command, " ")) == 3 {
				processQuery(command, animalNames)
			} else {
				fmt.Println("Incorrect command format. Please follow [query name info] such as [query clarabelle speak]")
			}
		} else if strings.HasPrefix(command, "newanimal") {
			if len(strings.Split(command, " ")) == 3 {
				arguments := strings.Split(command, " ")
				name, animalType := arguments[1], arguments[2]
				animalNames[name] = animalType
				fmt.Println("Created it!")
			} else {
				fmt.Println("Incorrect command format. Please follow [newanimal name type] such as [newanimal clarabelle cow]")
			}
		} else {
			fmt.Println("Incorrect command")
		}
	}
}
