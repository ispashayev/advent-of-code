package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day14.dat")
	check(err)

	input := strings.TrimSpace(string(dat))

	/* Part One */
	numRecipes, err := strconv.Atoi(input)
	check(err)
	scoreboard := []int{3, 7}
	e1pos, e2pos := 0, 1
	for len(scoreboard) < numRecipes+10 {
		e1score, e2score := scoreboard[e1pos], scoreboard[e2pos]
		newRecipes := strconv.Itoa(e1score + e2score)
		for _, newRecipe := range newRecipes {
			scoreboard = append(scoreboard, int(newRecipe-'0'))
		}
		e1pos = (e1pos + 1 + e1score) % len(scoreboard)
		e2pos = (e2pos + 1 + e2score) % len(scoreboard)
	}

	fmt.Printf("The scores of the ten recipes immediately after %d recipes: ", numRecipes)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", scoreboard[numRecipes+i])
	}
	fmt.Println()

	/* Part Two */
	scoreboard = []int{3, 7}
	e1pos, e2pos = 0, 1
	patternMatchCounter := 0
	for {
		e1score, e2score := scoreboard[e1pos], scoreboard[e2pos]
		newRecipes := strconv.Itoa(e1score + e2score)
		for _, newRecipe := range newRecipes {
			if byte(newRecipe) == input[patternMatchCounter] {
				patternMatchCounter++
			} else if byte(newRecipe) == input[0] {
				patternMatchCounter = 1
			} else {
				patternMatchCounter = 0
			}
			scoreboard = append(scoreboard, int(newRecipe-'0'))
			if patternMatchCounter == len(input) {
				fmt.Printf("There are %d recipes to the left of the input score sequence\n", len(scoreboard)-len(input))
				return
			}
		}
		e1pos = (e1pos + 1 + e1score) % len(scoreboard)
		e2pos = (e2pos + 1 + e2score) % len(scoreboard)
	}
}
