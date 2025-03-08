package main

// https://en.wikipedia.org/wiki/Monty_Hall_problem

import (
	"fmt"
	"math/rand/v2"
)

var games = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func main() {
	var (
		wonWithoutChange int
		wonWithSwitch    int
		rounds           = 10000
	)

	for i := 0; i < rounds*2; i++ {
		gameType := rand.IntN(3)
		game := games[gameType]

		choice := rand.IntN(3)

		var possibleGoats []int
		for k := 0; k < 3; k++ {
			if k != choice && game[k] == 0 {
				possibleGoats = append(possibleGoats, k)
			}
		}

		goatIndex := possibleGoats[rand.IntN(len(possibleGoats))]

		if i < rounds {
			// without switch
			if game[choice] == 1 {
				wonWithoutChange++
			}
		} else {
			// with switch
			var newChoice int
			for k := 0; k < 3; k++ {
				if k != choice && k != goatIndex {
					newChoice = k
					break // We only need the one remaining door
				}
			}
			if game[newChoice] == 1 {
				wonWithSwitch++
			}
		}
	}

	fmt.Printf("Rounds => %v | Keep => %v (%.2f%%) | Switch => %v (%.2f%%)\n",
		rounds,
		wonWithoutChange, float64(wonWithoutChange)/float64(rounds)*100,
		wonWithSwitch, float64(wonWithSwitch)/float64(rounds)*100)
}
