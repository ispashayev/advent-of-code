package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
)


func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func trim(pots string, origin int) (string, int) {
  for i := 0; i < len(pots); i++ {
    if pots[i] == '#' {
      pots = pots[i:]
      origin += i
      break
    }
  }
  for i := len(pots); i > 0; i-- {
    if pots[i-1] == '#' {
      pots = pots[:i]
      break
    }
  }
  return pots, origin
}

func readData(r io.Reader) (string, int, map[string]byte) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanLines)

  // Read the initial state
  scanner.Scan()
  pots, origin := trim(scanner.Text()[15:], 0)
  scanner.Scan()

  // Read the plant pattern matching rules
  plantPatternRules := make(map[string]byte)
  for scanner.Scan() {
    plantPatternRule := strings.Split(scanner.Text(), " => ")
    if len(plantPatternRule) != 2 {
      panic("Improperly formatted plant pattern rule.")
    }
    pattern, result := plantPatternRule[0], plantPatternRule[1]
    if len(result) != 1 {
      panic("Improperly formatted plant pattern rule.")
    }
    plantPatternRules[pattern] = result[0]
  }

  check(scanner.Err())

  return pots, origin, plantPatternRules
}

func step(pots string, origin int, plantPatternRules map[string]byte) (string, int) {
  pots = "....." + pots + "....."
  
  var pattern, potsNext string
  for i := 0; i <= len(pots)-5; i++ {
    pattern = pots[i:i+5]
    potsNext += string(plantPatternRules[pattern])
  }

  return trim(potsNext, origin-3)
}

func main() {
  fd, err := os.Open("day12.dat")
  check(err)

  dataReader := io.Reader(fd)
  pots, origin, plantPatternRules := readData(dataReader)

  for i := 0; i < 20; i++ {
    pots, origin = step(pots, origin, plantPatternRules)
  }

  sumNumbers := 0
  for i, pot := range pots {
    if pot == '#' {
      sumNumbers += i + origin
    }
  }
  fmt.Println("The total sum of the pots-with-plants numbers is:", sumNumbers)
}
