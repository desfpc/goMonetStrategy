# goMonetStrategy

A Go library that provides various utilities for strategy games and simulations.

## Features

### Dice

The library includes a dice implementation with random distribution:

- Create a dice with any number of sides
- Roll the dice to get a random number between 1 and the number of sides
- Utility function for quick dice rolls

## Installation

```bash
go get github.com/yourusername/goMonetStrategy
```

## Usage

### Dice

```go
package main

import (
    "fmt"
    "github.com/yourusername/goMonetStrategy"
)

func main() {
    // Quick dice roll
    result := goMonetStrategy.RollDice(6) // Roll a 6-sided dice
    fmt.Println("Dice roll result:", result)

    // Create a dice object for multiple rolls
    dice := goMonetStrategy.NewDice(20) // Create a 20-sided dice
    if dice != nil {
        for i := 0; i < 5; i++ {
            fmt.Printf("Roll %d: %d\n", i+1, dice.Roll())
        }
    }
}
```

### Error Handling

The dice functions handle invalid inputs gracefully:

- `NewDice` returns `nil` if the number of sides is less than or equal to 0
- `RollDice` returns 0 if the number of sides is less than or equal to 0
- `Roll` method returns 0 if called on a `nil` dice

## Examples

See the [example](./example) directory for more usage examples.

## License

[MIT](LICENSE)
