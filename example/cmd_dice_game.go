package main

import (
	"fmt"
	"goMonetStrategy"
)

func main() {
	numGames := 1000000

	fmt.Println("Roll:")
	regularResults := goMonetStrategy.RunSimulation(numGames, false)
	goMonetStrategy.PrintResults(regularResults)

	fmt.Println("\nRollRealistic:")
	realisticResults := goMonetStrategy.RunSimulation(numGames, true)
	goMonetStrategy.PrintResults(realisticResults)
}
