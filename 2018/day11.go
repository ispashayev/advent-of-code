package main

import (
	"fmt"
  "io/ioutil"
  "strconv"
  "strings"
)


func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

type FuelCell struct {
  x int
  y int
  powerLevel int
}

func (cell *FuelCell) SetPowerLevel(gridSerialNumber int) {
  rackId := cell.x + 10
  powerLevel := rackId * (rackId * cell.y + gridSerialNumber)
  powerLevel = ((powerLevel / 100) % 10) - 5
  cell.powerLevel = powerLevel
}

func makeGrid(gridSerialNumber int) [][]*FuelCell {
  grid := make([][]*FuelCell, 300)
  for i := 0; i < 300; i++ {
    grid[i] = make([]*FuelCell, 300)
    for j := 0; j < 300; j++ {
      cell := &FuelCell{i, j, 0}
      cell.SetPowerLevel(gridSerialNumber)
      grid[i][j] = cell
    }
  }
  return grid
}

func getPatchPower(grid [][]*FuelCell, i int, j int) (patchPower int) {
  patchPower += grid[i-1][j-1].powerLevel + grid[i-1][j].powerLevel + grid[i-1][j+1].powerLevel
  patchPower += grid[i][j-1].powerLevel + grid[i][j].powerLevel + grid[i][j+1].powerLevel
  patchPower += grid[i+1][j-1].powerLevel + grid[i+1][j].powerLevel + grid[i+1][j+1].powerLevel
  return patchPower
}

func main() {
  data, err := ioutil.ReadFile("day11.dat")
  check(err)

  gridSerialNumber, err := strconv.Atoi(strings.TrimSpace(string(data)))
  check(err)

  grid := makeGrid(gridSerialNumber)

  var highestPatchPower int
  var source *FuelCell
  for i := 1; i < 299; i++ {
    for j := 1; j < 299; j++ {
      current := grid[i][j]
      currentPatchPower := getPatchPower(grid, i, j)
      if currentPatchPower > highestPatchPower {
        highestPatchPower = currentPatchPower
        source = current
      }
    }
  }

  fmt.Printf("The highest level patch is centered at (%d,%d).\n", source.x, source.y)
  fmt.Printf("The patch's power level is: %d.\n", highestPatchPower)
}
