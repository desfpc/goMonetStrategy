package goMonetStrategy

import (
	"testing"
)

func TestNewDice(t *testing.T) {
	tests := []struct {
		name    string
		sides   int
		wantNil bool
	}{
		{
			name:    "valid sides",
			sides:   6,
			wantNil: false,
		},
		{
			name:    "zero sides",
			sides:   0,
			wantNil: true,
		},
		{
			name:    "negative sides",
			sides:   -1,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDice(tt.sides)
			if (got == nil) != tt.wantNil {
				t.Errorf("NewDice(%v) = %v, want nil: %v", tt.sides, got, tt.wantNil)
			}
			if got != nil && got.sides != tt.sides {
				t.Errorf("NewDice(%v).sides = %v, want %v", tt.sides, got.sides, tt.sides)
			}
		})
	}
}

func TestDice_Roll(t *testing.T) {
	sides := 6
	dice := NewDice(sides)

	// Test multiple rolls to ensure they're within the expected range
	for i := 0; i < 100; i++ {
		roll := dice.Roll()
		if roll < 1 || roll > sides {
			t.Errorf("Roll() = %v, want between 1 and %v", roll, sides)
		}
	}

	// Test nil dice
	var nilDice *Dice
	if nilDice.Roll() != 0 {
		t.Errorf("nil Dice.Roll() = %v, want 0", nilDice.Roll())
	}
}

func TestRollDice(t *testing.T) {
	tests := []struct {
		name     string
		sides    int
		wantZero bool
	}{
		{
			name:     "valid sides",
			sides:    6,
			wantZero: false,
		},
		{
			name:     "zero sides",
			sides:    0,
			wantZero: true,
		},
		{
			name:     "negative sides",
			sides:    -1,
			wantZero: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RollDice(tt.sides)
			if tt.wantZero && got != 0 {
				t.Errorf("RollDice(%v) = %v, want 0", tt.sides, got)
			}
			if !tt.wantZero && (got < 1 || got > tt.sides) {
				t.Errorf("RollDice(%v) = %v, want between 1 and %v", tt.sides, got, tt.sides)
			}
		})
	}
}

func TestDice_RollRealistic(t *testing.T) {
	sides := 6
	dice := NewDice(sides)

	// Test multiple rolls to ensure they're within the expected range
	for i := 0; i < 100; i++ {
		roll := dice.RollRealistic()
		if roll < 1 || roll > sides {
			t.Errorf("RollRealistic() = %v, want between 1 and %v", roll, sides)
		}
	}

	// Test nil dice
	var nilDice *Dice
	if nilDice.RollRealistic() != 0 {
		t.Errorf("nil Dice.RollRealistic() = %v, want 0", nilDice.RollRealistic())
	}
}

func TestRollDiceRealistic(t *testing.T) {
	tests := []struct {
		name     string
		sides    int
		wantZero bool
	}{
		{
			name:     "valid sides",
			sides:    6,
			wantZero: false,
		},
		{
			name:     "zero sides",
			sides:    0,
			wantZero: true,
		},
		{
			name:     "negative sides",
			sides:    -1,
			wantZero: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RollDiceRealistic(tt.sides)
			if tt.wantZero && got != 0 {
				t.Errorf("RollDiceRealistic(%v) = %v, want 0", tt.sides, got)
			}
			if !tt.wantZero && (got < 1 || got > tt.sides) {
				t.Errorf("RollDiceRealistic(%v) = %v, want between 1 and %v", tt.sides, got, tt.sides)
			}
		})
	}
}

func TestCompareRollDistributions(t *testing.T) {
	// This test checks that RollRealistic produces a different distribution than Roll
	// Note: This is a statistical test and may occasionally fail by chance
	sides := 6
	dice := NewDice(sides)

	// Reset the dice to ensure consistent state
	dice.lastRoll = 0
	dice.rollCount = 0

	// Track results for both methods
	regularRolls := make(map[int]int)
	realisticRolls := make(map[int]int)

	// Perform many rolls to get a good sample
	numRolls := 1000
	for i := 0; i < numRolls; i++ {
		// Create new dice for each roll to avoid state influence
		regularDice := NewDice(sides)
		realisticDice := NewDice(sides)

		regularRoll := regularDice.Roll()
		realisticRoll := realisticDice.RollRealistic()

		regularRolls[regularRoll]++
		realisticRolls[realisticRoll]++
	}

	// Check that the distributions are different
	// We'll use a simple chi-square test to compare distributions
	differentCount := 0
	for i := 1; i <= sides; i++ {
		regularCount := regularRolls[i]
		realisticCount := realisticRolls[i]

		// If the counts differ by more than 10%, we'll consider them different
		if float64(abs(regularCount-realisticCount))/float64(numRolls/sides) > 0.1 {
			differentCount++
		}
	}

	// We expect at least some of the sides to have different distributions
	if differentCount == 0 {
		t.Logf("Regular rolls distribution: %v", regularRolls)
		t.Logf("Realistic rolls distribution: %v", realisticRolls)
		t.Error("RollRealistic and Roll produced identical distributions, expected some differences")
	}
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
