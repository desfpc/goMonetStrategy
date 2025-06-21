package goMonetStrategy

import (
	"math/rand"
	"time"
)

// Dice represents a dice with a specified number of sides
type Dice struct {
	sides     int
	rng       *rand.Rand
	lastRoll  int
	rollCount int
}

// NewDice creates a new dice with the specified number of sides
// Returns nil if sides is less than or equal to 0
func NewDice(sides int) *Dice {
	if sides <= 0 {
		return nil
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &Dice{
		sides:     sides,
		rng:       rng,
		lastRoll:  0,
		rollCount: 0,
	}
}

// Roll returns a random number between 1 and the number of sides of the dice
func (d *Dice) Roll() int {
	if d == nil {
		return 0
	}

	result := d.rng.Intn(d.sides) + 1
	d.lastRoll = result
	d.rollCount++
	return result
}

// RollRealistic returns a number between 1 and the number of sides of the dice,
// simulating more realistic dice behavior by considering:
// - Physical biases (some sides are more likely to come up than others)
// - Momentum effects (previous rolls influence future rolls)
// - "Breaking in" effects (dice behavior changes after multiple rolls)
func (d *Dice) RollRealistic() int {
	if d == nil {
		return 0
	}

	// Base roll with slight randomness
	baseRoll := d.rng.Intn(d.sides) + 1

	// Apply physical bias - opposite sides of a die add up to n+1
	// (e.g., on a 6-sided die, 1 is opposite to 6, 2 is opposite to 5, etc.)
	// We'll make higher numbers slightly more likely (as if the die is slightly heavier on the lower-number sides)
	biasAdjustment := 0
	if d.sides > 3 && d.rng.Float64() < 0.3 { // 30% chance of bias effect
		if baseRoll <= d.sides/2 {
			// Slight chance to "roll over" to the opposite side (higher number)
			if d.rng.Float64() < 0.15 {
				biasAdjustment = d.sides + 1 - 2*baseRoll
			}
		}
	}

	// Apply momentum effect - previous roll influences current roll
	momentumAdjustment := 0
	if d.lastRoll > 0 && d.rng.Float64() < 0.2 { // 20% chance of momentum effect
		// Tendency to roll numbers adjacent to the previous roll
		diff := d.rng.Intn(3) - 1 // -1, 0, or 1
		candidateRoll := d.lastRoll + diff

		// Ensure the result is within valid range
		if candidateRoll >= 1 && candidateRoll <= d.sides {
			momentumAdjustment = candidateRoll - baseRoll
		}
	}

	// Apply "breaking in" effect - dice behavior stabilizes after multiple rolls
	breakingInFactor := 0
	if d.rollCount < 10 && d.rng.Float64() < 0.1 { // 10% chance when dice is "new"
		// New dice might be more erratic
		breakingInFactor = d.rng.Intn(3) - 1 // -1, 0, or 1
	}

	// Calculate final result with all adjustments
	result := baseRoll + biasAdjustment + momentumAdjustment + breakingInFactor

	// Ensure result is within valid range
	if result < 1 {
		result = 1
	} else if result > d.sides {
		result = d.sides
	}

	// Update dice state
	d.lastRoll = result
	d.rollCount++

	return result
}

// RollDice is a utility function that creates a dice with the specified number of sides
// and returns a random number between 1 and the number of sides
// Returns 0 if sides is less than or equal to 0
func RollDice(sides int) int {
	dice := NewDice(sides)
	if dice == nil {
		return 0
	}

	return dice.Roll()
}

// RollDiceRealistic is a utility function that creates a dice with the specified number of sides
// and returns a number between 1 and the number of sides using the realistic dice rolling simulation
// Returns 0 if sides is less than or equal to 0
func RollDiceRealistic(sides int) int {
	dice := NewDice(sides)
	if dice == nil {
		return 0
	}

	return dice.RollRealistic()
}
