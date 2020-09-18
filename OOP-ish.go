package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func Eat(animale animal) {
	fmt.Println("Food:", animale.food)
}

func Move(animalm animal) {
	fmt.Println("Locomotion:", animalm.locomotion)
}

func Speak(animalsp animal) {
	fmt.Println("Sound:", animalsp.noise)
}

func main() {
	var cow animal = animal{"grass", "walk", "moo"}
	var bird animal = animal{"worms", "fly", "peep"}
	var snake animal = animal{"mice", "slither", "hsss"}
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
