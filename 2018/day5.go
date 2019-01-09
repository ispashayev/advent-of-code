package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)


func check(err error) {
  if err != nil {
    panic(err)
  }
}

func unitsTriggered(one, two byte) bool {
  return two - ' ' == one || one - ' ' == two
}

func react(polymer string) string {
  i := 0
  for i < len(polymer)-1 {
    if unitsTriggered(polymer[i], polymer[i+1]) {
      polymer = polymer[:i] + polymer[i+2:]
      i = 0
    } else {
      i++
    }
  }
  return polymer
}

func diffReact(polymer string) map[rune]string {
  differentials := make(map[rune]string)
  var diffPolymer, diffReduced string
  for i := 'a'; i <= 'z'; i++ {
    diffPolymer = strings.Replace(polymer, string(i), "", -1)
    diffPolymer = strings.Replace(diffPolymer, string(i - ' '), "", -1)
    diffReduced = react(diffPolymer)
    differentials[rune(i)] = diffReduced
  }
  return differentials
}

func main() {
  sequence, err := ioutil.ReadFile("day5.dat")
  check(err)

  polymer := string(sequence)[:len(sequence)-1] // remove trailing new line
  fmt.Println("Number of units in original polymer:", len(polymer))
  
  reduced := react(polymer)
  fmt.Println("Number of units after reaction:", len(reduced))

  fmt.Println("--")
  
  differentials := diffReact(polymer)

  var maxReduceType rune
  minDiffReduced := polymer
  for unitType, diffReduced := range differentials {
    fmt.Println("Removing type", unitType, "yields a reduced polymer of length", len(diffReduced))
    if len(diffReduced) < len(minDiffReduced) {
      maxReduceType = unitType
      minDiffReduced = diffReduced
    }
  }
  fmt.Println("--")
  fmt.Println("The largest possible reduction is by removing type" , maxReduceType)
  fmt.Println("The reaction from the resulting polymer yields", len(minDiffReduced), "units")
}
