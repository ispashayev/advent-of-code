package main

import (
	"fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

type FuelGrid interface {
  PowerLevel(int) int
}

type FuelCell struct {
  x int
  y int
}

func (cell *FuelCell) PowerLevel(gridSerialNumber int) (powerLevel int) {
  rackId := cell.x + 10
  powerLevel = rackId * (rackId * cell.y + gridSerialNumber)
  powerLevel = ((powerLevel / 100) % 10) - 5
  return powerLevel
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func makeGrid() [][]*FuelCell {
  grid := make([][]*FuelCell, 300)
  for i := 0; i < 300; i++ {
    grid[i] = make([]*FuelCell, 300)
    for j := 0; j < 300; j++ {
      grid[i][j] = &FuelCell{i,j}
    }
  }
  return grid
}

func testPowerLevel(cell *FuelCell, gridSerialNumber int, expected int) {
  fmt.Printf("Testing fuel cell at (%d,%d).", cell.x, cell.y)
  actual := cell.PowerLevel(gridSerialNumber)
  pass := expected == actual
  if pass {
    fmt.Printf("Pass.\n")
  } else {
    fmt.Printf("\nFAIL\n")
    fmt.Printf("EXPECTED: %d\n", expected)
    fmt.Printf("ACTUAL: %d\n\n", actual)
  }
}

func main() {
  data, err := ioutil.ReadFile("day11.dat")
  check(err)

  gridSerialNumber, err := strconv.Atoi(strings.TrimSpace(string(data)))
  check(err)

  grid := makeGrid()
  fmt.Println("10,10 powerlevel", grid[10][10].PowerLevel(gridSerialNumber))

  testPowerLevel(grid[122][79], 57, -5)
  testPowerLevel(grid[217][196], 39, 0)
  testPowerLevel(grid[101][153], 71, 4)
}
