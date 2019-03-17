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

func sumPotsWithPlants(pots string, origin int) (pwpSum int) {
  for i, pot := range pots {
    if pot == '#' {
      pwpSum += i + origin
    }
  }
  return pwpSum
}

func main() {
  fd, err := os.Open("day12.dat")
  check(err)

  dataReader := io.Reader(fd)
  pots, origin, plantPatternRules := readData(dataReader)

  prevPwpSum, prevPwpSumDiff, stablePwpSumCount := 0, 0, 0
  numGenerations, stabilityThreshold, genDelta := 50000000000, 100, 20
  for i := 0; i <= numGenerations; i++ {
    if i % genDelta == 0 {
      pwpSum := sumPotsWithPlants(pots, origin)
      if i == genDelta {
        fmt.Printf("After %d generations, the pots-with-plants sum is: %d.\n", genDelta, pwpSum)
      }
      pwpSumDiff := pwpSum - prevPwpSum
      if pwpSumDiff == prevPwpSumDiff {
        stablePwpSumCount++
      }
      if stablePwpSumCount == stabilityThreshold {
        fmt.Printf("Stable difference inferred after %d generations: %d.\n", i, pwpSumDiff)
        projectedPwpSum := pwpSum + ((numGenerations - i) / genDelta) * pwpSumDiff
        fmt.Printf("After %d generations, the project pot-with-plants sum is: %d.\n", numGenerations, projectedPwpSum)
        break
      }
      prevPwpSum = pwpSum
      prevPwpSumDiff = pwpSumDiff
    }

    pots, origin = step(pots, origin, plantPatternRules)
  }
}
