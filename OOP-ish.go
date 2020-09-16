package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func Eat(animal Animal) {
	fmt.Println("Food:", animal.food)
}

func Move(animal Animal) {
	fmt.Println("Locomotion:", animal.locomotion)
}

func Speak(animal Animal) {
	fmt.Println("Sound:", animal.noise)
}

func main() {
	var cow Animal = Animal{"grass", "walk", "moo"}
	var bird Animal = Animal{"worms", "fly", "peep"}
	var snake Animal = Animal{"mice", "slither", "hsss"}
	var choice string
	for {
		fmt.Println("Enter an animal name (cow, bird, or snake) and attribute (eat, move, or speak) separated by space OR type q to exit:")
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice = scanner.Text()
		if strings.EqualFold(strings.TrimSpace(choice), "Q") {
			break
		} else {
			line := strings.Split(choice, " ")
			animalName, attributeName := line[0], line[1]
			if strings.EqualFold(strings.TrimSpace(attributeName), "eat") {
				if strings.EqualFold(strings.TrimSpace(animalName), "cow") {
					Eat(cow)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "bird") {
					Eat(bird)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "snake") {
					Eat(snake)
				}
			}
			if strings.EqualFold(strings.TrimSpace(attributeName), "move") {
				if strings.EqualFold(strings.TrimSpace(animalName), "cow") {
					Move(cow)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "bird") {
					Move(bird)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "snake") {
					Move(snake)
				}
			}
			if strings.EqualFold(strings.TrimSpace(attributeName), "speak") {
				if strings.EqualFold(strings.TrimSpace(animalName), "cow") {
					Speak(cow)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "bird") {
					Speak(bird)
				}
				if strings.EqualFold(strings.TrimSpace(animalName), "snake") {
					Speak(snake)
				}
			}
		}
	}
}
