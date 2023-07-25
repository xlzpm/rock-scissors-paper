package main

import (
	"fmt"
	"math/rand"
)

var number int

const (
	text = "Input number where: \nrock = 0 \npaper = 1 \nscissors = 2"
)

func RandomComputer() int {
	values := rand.Intn(3)
	return values
}

func RSP() {
	fmt.Scan(&number)
	random := RandomComputer()
	switch number {
	case 0:
		if random == 2 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You win\n\n", number, random)
		} else if random == 1 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You loser\n\n", number, random)
		} else {
			fmt.Printf("Result: your choice = %d and computer = %d\n Tie\n\n", number, random)
		}
	case 1:
		if random == 0 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You win\n\n", number, random)
		} else if random == 2 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You loser\n\n", number, random)
		} else {
			fmt.Printf("Result: your choice = %d and computer = %d\n Tie\n\n", number, random)
		}
	case 2:
		if random == 1 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You win\n\n", number, random)
		} else if random == 0 {
			fmt.Printf("Result: your choice = %d and computer = %d\n You loser\n\n", number, random)
		} else {
			fmt.Printf("Result: your choice = %d and computer = %d\n Tie\n\n", number, random)
		}
	default:
		fmt.Print("Entered incorrect data.... Once again\n\n")
	}

}

func textRSP(text string) {
	fmt.Println(text)
}

func main() {
	for {
		textRSP(text)
		RSP()
	}
}
