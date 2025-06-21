package goMonetStrategy

import (
	"fmt"
	"math"
)

// Strategy represents a decision-making strategy for the dice game
type Strategy struct {
	Name       string
	Thresholds []int
}

// Define the three strategies
var (
	strategy1 = Strategy{
		Name:       "Стратегия 1",
		Thresholds: []int{93, 90, 87, 84, 80, 76, 71, 66, 51, 0}, // 0 for the 10th throw (always take)
	}
	strategy2 = Strategy{
		Name:       "Стратегия 2",
		Thresholds: []int{100, 100, 100, 99, 98, 95, 88, 76, 51, 0}, // 0 for the 10th throw (always take)
	}
	strategy3 = Strategy{
		Name:       "Стратегия 3",
		Thresholds: []int{80, 70, 65, 60, 60, 60, 60, 55, 51, 0}, // 0 for the 10th throw (always take)
	}
)

// PlayGame simulates one game with the given strategy and dice rolling method
// Returns the winnings in dollars
func PlayGame(strategy Strategy, useRealistic bool) int {
	dice := NewDice(100)

	for throw := 0; throw < 10; throw++ {
		var roll int
		if useRealistic {
			roll = dice.RollRealistic()
		} else {
			roll = dice.Roll()
		}

		// On the 10th throw, always take the result
		if throw == 9 || roll >= strategy.Thresholds[throw] {
			return roll * 1000 // Return winnings in dollars
		}
	}

	// This should never happen as we always take the 10th throw
	return 0
}

// RunSimulation runs the specified number of games for each strategy
// and returns statistics for each strategy
func RunSimulation(numGames int, useRealistic bool) map[string]map[string]float64 {
	strategies := []Strategy{strategy1, strategy2, strategy3}
	results := make(map[string]map[string]float64)

	for _, strategy := range strategies {
		minWin := math.MaxInt32
		maxWin := 0
		totalWin := 0

		for i := 0; i < numGames; i++ {
			winnings := PlayGame(strategy, useRealistic)

			if winnings < minWin {
				minWin = winnings
			}
			if winnings > maxWin {
				maxWin = winnings
			}
			totalWin += winnings
		}

		avgWin := float64(totalWin) / float64(numGames)

		results[strategy.Name] = map[string]float64{
			"min": float64(minWin),
			"max": float64(maxWin),
			"avg": avgWin,
		}
	}

	return results
}

// PrintResults prints the simulation results in the required format
func PrintResults(results map[string]map[string]float64) {
	for strategy, stats := range results {
		fmt.Printf("%s:\n", strategy)
		fmt.Printf("мин: %.0f$, макс: %.0f$, среднее: %.0f$\n", stats["min"], stats["max"], stats["avg"])
	}
}
