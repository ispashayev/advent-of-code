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

func makeGrid(gridSize int, gridSerialNumber int) [][]*FuelCell {
  grid := make([][]*FuelCell, gridSize)
  for i := 0; i < gridSize; i++ {
    grid[i] = make([]*FuelCell, gridSize)
    for j := 0; j < gridSize; j++ {
      cell := &FuelCell{i, j, 0}
      cell.SetPowerLevel(gridSerialNumber)
      grid[i][j] = cell
    }
  }
  return grid
}

func getPatchPower(grid [][]*FuelCell, x int, y int, patchSize int) (patchPower int) {
  for i := 0; i < patchSize; i++ {
    for j := 0; j < patchSize; j++ {
      patchPower += grid[x+i][y+j].powerLevel
    }
  }
  return patchPower
}

func main() {
  data, err := ioutil.ReadFile("day11.dat")
  check(err)

  gridSerialNumber, err := strconv.Atoi(strings.TrimSpace(string(data)))
  check(err)

  gridSize := 300
  grid := makeGrid(gridSize, gridSerialNumber)

  var sourcePatchPower, sourcePatchSize int
  var source *FuelCell
  for patchSize := 1; patchSize <= gridSize; patchSize++ {
    for i := 0; i <= gridSize - patchSize; i++ {
      for j := 0; j <= gridSize - patchSize; j++ {
        currentPatchPower := getPatchPower(grid, i, j, patchSize)
        if currentPatchPower > sourcePatchPower {
          source = grid[i][j]
          sourcePatchPower = currentPatchPower
          sourcePatchSize = patchSize
        }
      }
    }
  }

  fmt.Printf("The strongest patch's power level is: %d.\n", sourcePatchPower)
  fmt.Printf("Its top-left corner is (%d, %d).\n", source.x, source.y)
  fmt.Printf("The patch size is %d.\n", sourcePatchSize)
}
