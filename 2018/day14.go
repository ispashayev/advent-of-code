package main

import (
  "fmt"
  "io/ioutil"
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

  scoreboardInput := strings.TrimSpace(string(dat))
  scoreboard := make([]int, len(scoreboardInput))
  for i, score := range scoreboardInput {
    scoreboard[i] = int(score - '0')
  }
  
  fmt.Printf("The input scoreboard is: %s\n", scoreboardInput)
  fmt.Println("The recipe scores are:")
  for i, score := range scoreboard {
    fmt.Printf("Recipe %d: %d\n", i, score)
  }
}
