package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

type Marble struct {
  number int
  prev *Marble
  next *Marble
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func getMax(scores map[int]int) (maxScore int) {
  for _, score := range scores {
    if score > maxScore {
      maxScore = score
    }
  }
  return maxScore
}

func playMarbleGame(numPlayers, finalPoints int) map[int]int {
  scores := make(map[int]int)

  currentMarble, currentPlayer := &Marble{0, nil, nil}, 0
  currentMarble.prev = currentMarble
  currentMarble.next = currentMarble

  for i := 1; i <= finalPoints; i++ {
    if i % 23 == 0 {
      scores[currentPlayer] += i

      for j := 0; j < 7; j++ {
        currentMarble = currentMarble.prev
      }

      scores[currentPlayer] += currentMarble.number
      currentMarble.prev.next = currentMarble.next
      currentMarble.next.prev = currentMarble.prev

      currentMarble = currentMarble.next
      currentPlayer = (currentPlayer + 1) % numPlayers
      continue
    }

    newMarble := Marble{i, currentMarble.next, currentMarble.next.next}

    currentMarble.next.next = &newMarble
    newMarble.next.prev = &newMarble

    currentMarble = &newMarble
    currentPlayer = (currentPlayer + 1) % numPlayers
  }

  return scores
}

func main() {
  dat, err := ioutil.ReadFile("day9.dat")
  check(err)

  tokens := strings.Split(string(dat), " ")
  numPlayers, err := strconv.Atoi(tokens[0])
  check(err)
  finalPoints, err := strconv.Atoi(tokens[len(tokens)-2])
  check(err)

  fmt.Printf("Simulating game with %d players and final marble %d.\n", numPlayers, finalPoints)
  scores := playMarbleGame(numPlayers, finalPoints)
  fmt.Println("The winning elf's score is", getMax(scores))

  fmt.Printf("Simulating game with %d players and final marble %d.\n", numPlayers, 100 * finalPoints)
  scores2 := playMarbleGame(numPlayers, 100 * finalPoints)
  fmt.Println("IfThe winning elf's score is", getMax(scores2))
}
