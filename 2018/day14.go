package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var kNumPlayers = 2

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day14.dat")
	check(err)

	input := strings.TrimSpace(string(dat))
	numRecipes, err := strconv.Atoi(input)
	check(err)

	fmt.Printf("The input number of recipes is: %d\n", numRecipes)

	scoreboard := []int{3, 7}
	e1pos, e2pos := 0, 1

	for len(scoreboard) < numRecipes+10 {
		e1score, e2score := scoreboard[e1pos], scoreboard[e2pos]
		newRecipes := strconv.Itoa(e1score + e2score)
		for _, newRecipe := range newRecipes {
			scoreboard = append(scoreboard, int(newRecipe-'0'))
		}
		e1pos = (1 + e1score) % len(scoreboard)
		e2pos = (1 + e2score) % len(scoreboard)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", scoreboard[numRecipes+i])
	}
	fmt.Println()
}
