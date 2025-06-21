package main

import (
	"fmt"
	"goMonetStrategy"
)

func main() {
	// Example 1: Using the RollDice utility function
	fmt.Println("Rolling a 6-sided dice:", goMonetStrategy.RollDice(6))
	fmt.Println("Rolling a 20-sided dice:", goMonetStrategy.RollDice(20))
	fmt.Println("Rolling a 100-sided dice:", goMonetStrategy.RollDice(100))

	// Example 2: Creating and using a Dice object
	d6 := goMonetStrategy.NewDice(6)
	if d6 != nil {
		fmt.Println("\nRolling a 6-sided dice 5 times:")
		for i := 0; i < 5; i++ {
			fmt.Printf("Roll %d: %d\n", i+1, d6.Roll())
		}
	}

	d100 := goMonetStrategy.NewDice(100)
	if d100 != nil {
		fmt.Println("\nRolling a 100-sided dice 5 times:")
		for i := 0; i < 5; i++ {
			fmt.Printf("Roll %d: %d\n", i+1, d100.Roll())
		}
	}

	// Example 3: Handling invalid input
	invalidDice := goMonetStrategy.NewDice(-1)
	if invalidDice == nil {
		fmt.Println("\nInvalid dice (negative sides) returns nil")
	}

	// Using the utility function with invalid input
	result := goMonetStrategy.RollDice(0)
	fmt.Println("Rolling a 0-sided dice returns:", result)
}
