package main

import (
  "bufio"
  "fmt"
  "io"
  "math"
  "os"
  "strconv"
  "strings"
)

type Coordinate struct {
  x int
  y int
  centroid *Coordinate
  torn bool
  label int
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func parseDocument(r io.Reader) ([]Coordinate, error) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanLines)
  
  var coordinates []Coordinate
  counter := 0
  for scanner.Scan() {
    coordinate_slice := strings.Split(scanner.Text(), " ")
    x, err := strconv.Atoi(coordinate_slice[0][:len(coordinate_slice[0])-1])
    check(err)
    y, err := strconv.Atoi(coordinate_slice[1])
    check(err)
    coordinate := Coordinate{x, y, nil, false, counter}
    coordinate.centroid = &coordinate
    coordinates = append(coordinates, coordinate)
    counter++
  }
  return coordinates, scanner.Err()
}

func manhattanDistance(one, two *Coordinate) int {
  xDiff, yDiff := float64(one.x - two.x), float64(one.y - two.y)
  return int(math.Abs(xDiff) + math.Abs(yDiff))
}

func checkAndSetCentroid(grid [][]*Coordinate, dstX int, dstY int, src *Coordinate) bool {
  dst := grid[dstX][dstY]

  if dst == nil {
    newCoordinate := Coordinate{dstX, dstY, src.centroid, false, -1}
    if manhattanDistance(&newCoordinate, src) != 1 {
      panic("Can only extend cluster coverage by distance of one.")
    }
    grid[dstX][dstY] = &newCoordinate
    return true
  }

  if src.centroid != dst.centroid {
    if manhattanDistance(dst, src.centroid) == manhattanDistance(dst, dst.centroid) {
      dst.torn = true
    }
  }

  return false
}

func extendCluster(grid [][]*Coordinate, pos *Coordinate) (neighbors []*Coordinate) {
  max_x, max_y := len(grid)-1, len(grid[0])-1
  x, y := pos.x, pos.y
  
  if x > 0 { // Extend left
    if checkAndSetCentroid(grid, x-1, y, pos) {
      neighbors = append(neighbors, grid[x-1][y])
    }
  }

  if x < max_x { // Extend right
    if checkAndSetCentroid(grid, x+1, y, pos) {
      neighbors = append(neighbors, grid[x+1][y])
    }
  }

  if y > 0 { // Extend down
    if checkAndSetCentroid(grid, x, y-1, pos) {
      neighbors = append(neighbors, grid[x][y-1])
    }
  }

  if y < max_y { // Extend up
    if checkAndSetCentroid(grid, x, y+1, pos) {
      neighbors = append(neighbors, grid[x][y+1])
    }
  }

  return neighbors
}

func computeClusters(coordinates []Coordinate) [][]*Coordinate {
  max_x, min_x := coordinates[0].x, coordinates[0].x
  max_y, min_y := coordinates[0].y, coordinates[0].y
  for _, coordinate := range coordinates {
    if coordinate.x > max_x { max_x = coordinate.x }
    if coordinate.x < min_x { min_x = coordinate.x }
    if coordinate.y > max_y { max_y = coordinate.y }
    if coordinate.y < min_y { min_y = coordinate.y }
  }

  // Initialize a 2D grid of the coordinates
  grid := make([][]*Coordinate, max_x - min_x + 1)
  for x := 0; x <= max_x - min_x; x++ {
    grid[x] = make([]*Coordinate, max_y - min_y + 1)
  }
  fmt.Println("Making a grid of size", len(grid), "x", len(grid[0]))
  
  // Mark the input coordinates in the grid and initialize the bfs queue
  var coordinate *Coordinate
  bfsQueue := make([]*Coordinate, len(coordinates))
  for i := range coordinates {
    coordinate = &coordinates[i]
    coordinate.x -= min_x
    coordinate.y -= min_y
    grid[coordinate.x][coordinate.y] = coordinate
    bfsQueue[i] = coordinate
  }

  for len(bfsQueue) > 0 {
    coordinate := bfsQueue[0]
    neighbors := extendCluster(grid, coordinate)
    bfsQueue = append(bfsQueue, neighbors...)[1:]
  }

  return grid
}

func centroidsDistance(pos *Coordinate, centroids []Coordinate) (distance int) {
  for _, centroid := range centroids {
    distance += manhattanDistance(pos, &centroid)
  }
  return distance
}

func computeLargestCluster(grid [][]*Coordinate, centroids []Coordinate) {
  infiniteClusters := make(map[int]bool)
  clusters := make(map[int]int)

  max_x, max_y := len(grid)-1, len(grid[0])-1
  boundedDistanceClusterArea := 0
  for x, slice := range grid {
    for y, coordinate := range slice {
      if !coordinate.torn {
        if x == 0 || x == max_x || y == 0 || y == max_y {
          infiniteClusters[coordinate.centroid.label] = true
        }
        clusters[coordinate.centroid.label]++
      }
      if centroidsDistance(coordinate, centroids) < 10000 {
        boundedDistanceClusterArea++
      }
    }
  }

  maxFiniteCluster := 0
  for cluster, size := range clusters {
    if !infiniteClusters[cluster] && size > clusters[maxFiniteCluster] {
      maxFiniteCluster = cluster
    }
  }

  fmt.Println("The largest finite cluster belongs to centroid", maxFiniteCluster)
  fmt.Println("The size of that cluster is", clusters[maxFiniteCluster])
  fmt.Println("The size of the region containing all locations with a total")
  fmt.Println("distance to centroids of less than 10000 is:", boundedDistanceClusterArea)
}

func main() {
  fd, err := os.Open("day06.dat")
  check(err)

  documentReader := io.Reader(fd)
  coordinates, err := parseDocument(documentReader)
  check(err)
  grid := computeClusters(coordinates)
  computeLargestCluster(grid, coordinates)
}
