package main

import (
	"bufio"
	"fmt"
  "io"
  "os"
  "strconv"
  "strings"
)


type Claim struct {
  id, x, y, w, h int
}


type Coordinate struct {
  x, y int
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}


func readClaim(claimReader io.Reader) *Claim {
  claimDataScanner := bufio.NewScanner(claimReader)
  claimDataScanner.Split(bufio.ScanWords)

  // Read claim ID
  if !claimDataScanner.Scan() {
    check(claimDataScanner.Err())
  }
  id, err := strconv.Atoi(claimDataScanner.Text()[1:])
  check(err)

  // Skip @ symbol
  if !claimDataScanner.Scan() {
    check(claimDataScanner.Err())
  }

  // Read coordinate of upper-left corner of rectangle
  if !claimDataScanner.Scan() {
    check(claimDataScanner.Err())
  }
  coordinate_string := claimDataScanner.Text()
  coordinates := strings.Split(coordinate_string[:len(coordinate_string)-1], ",")
  if len(coordinates) != 2 {
    panic("Invalid data")
  }
  x, err := strconv.Atoi(coordinates[0])
  check(err)
  y, err := strconv.Atoi(coordinates[1])
  check(err)

  // Read dimensions of rectangle
  if !claimDataScanner.Scan() {
    check(claimDataScanner.Err())
  }
  dimensions := strings.Split(claimDataScanner.Text(), "x")
  if len(dimensions) != 2 {
    panic("Invalid data")
  }
  w, err := strconv.Atoi(dimensions[0])
  check(err)
  h, err := strconv.Atoi(dimensions[1])
  check(err)

  // Assert there are no more tokens
  if claimDataScanner.Scan() {
    panic("Invalid data")
  }
  check(claimDataScanner.Err())

  return &Claim{id, x, y, w, h}
}


func readDocument(documentReader io.Reader) ([]*Claim, error) {
  documentScanner := bufio.NewScanner(documentReader)
  documentScanner.Split(bufio.ScanLines)
  var claims []*Claim
  for documentScanner.Scan() {
    claimReader := strings.NewReader(documentScanner.Text())
    claim := readClaim(claimReader)
    claims = append(claims, claim)
  }
  return claims, documentScanner.Err()
}


func measureOverlappingArea(claims []*Claim) (int, map[Coordinate]int) {
  overlappingArea, fabricDistribution := 0, make(map[Coordinate]int)
  for _, baseClaim := range claims {
    xBase, yBase := (*baseClaim).x, (*baseClaim).y
    for xDelta := 0; xDelta < (*baseClaim).w; xDelta++ {
      for yDelta := 0; yDelta < (*baseClaim).h; yDelta++ {
        coordinate := Coordinate{xBase + xDelta, yBase + yDelta}
        fabricDistribution[coordinate] += 1
      }
    }
  }

  for _, count := range fabricDistribution {
    if count > 1 {
      overlappingArea += 1
    }
  }

  return overlappingArea, fabricDistribution
}


func claimIsDistinct(claim *Claim, fabricDistribution map[Coordinate]int) bool {
  xBase, yBase := (*claim).x, (*claim).y
  for xDelta := 0; xDelta < (*claim).w; xDelta++ {
    for yDelta := 0; yDelta < (*claim).h; yDelta++ {
      coordinate := Coordinate{xBase + xDelta, yBase + yDelta}
      if fabricDistribution[coordinate] > 1 {
        return false
      }
    }
  }
  return true
}


func getDistinctClaim(claims []*Claim, fabricDistribution map[Coordinate]int) *Claim {
  for _, claim := range claims {
    if claimIsDistinct(claim, fabricDistribution) {
      return claim
    }
  }

  panic("Unable to find claim with non-overlapping fabric")
}


func main() {
  fd, err := os.Open("day03.dat")
  check(err)
  ioReader := io.Reader(fd)
  claims, err := readDocument(ioReader)
  check(err)

  overlappingArea, fabricDistribution := measureOverlappingArea(claims)
  fmt.Println("Square inches of overlapping fabric:", overlappingArea)

  distinctClaim := getDistinctClaim(claims, fabricDistribution)
  fmt.Println("The claim with non-overlapping fabric is is:", (*distinctClaim).id)
}
