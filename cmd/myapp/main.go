package main

import (
	"fmt"
  "os"
	"strconv"
	"math/rand/v2"
)

func rollDice() int {
	return rand.IntN(7 - 1) + 1 // Returns numbers in the range 1 to 6
}

func main() {
	fmt.Println("Dice Roller App: ")

	if len(os.Args) == 1 {
		fmt.Println(rollDice())	
		return
	}

	numOfDices, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	i := 0
	for i < numOfDices {
		fmt.Printf("%d ", rollDice())
		i += 1
	}

	fmt.Println()
}

