package main

import (
  "bufio"
  "fmt"
  "io"
  "math"
  "os"
  "regexp"
  "strconv"
  "strings"
)

type Point struct {
  x float64
  y float64
  dx float64
  dy float64
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func readData(r io.Reader) (points []*Point) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanLines)

  positionRe := regexp.MustCompile("position=<.+?, .+?>")
  velocityRe := regexp.MustCompile("velocity=<.+?, .+?>")
  for scanner.Scan() {
    pointData := scanner.Text()

    matchedPosition := positionRe.FindString(pointData)
    positionTokens := strings.Split(matchedPosition[9:], ",")
    xString := positionTokens[0][1:]
    x, err := strconv.Atoi(strings.TrimSpace(xString))
    check(err)
    yString := positionTokens[1][:len(positionTokens[1])-1]
    y, err := strconv.Atoi(strings.TrimSpace(yString))
    check(err)

    matchedVelocity := velocityRe.FindString(pointData)
    check(err)
    velocityTokens := strings.Split(matchedVelocity[9:], ",")
    dxString := velocityTokens[0][1:]
    dx, err := strconv.Atoi(strings.TrimSpace(dxString))
    check(err)
    dyString := velocityTokens[1][:len(velocityTokens[1])-1]
    dy, err := strconv.Atoi(strings.TrimSpace(dyString))
    check(err)

    points = append(points, &Point{float64(x), float64(y), float64(dx), float64(dy)})
  }

  check(scanner.Err())

  return points
}

func distance(a, b *Point) (d float64) {
  return math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2))
}

func averageDistance(points []*Point) (d float64) {
  for i := 0; i < len(points); i++ {
    for j := i + 1; j < len(points); j++ {
      d += distance(points[i], points[j])
    }
  }
  return d / float64(len(points))
}

func step(points []*Point) {
  for _, point := range points {
    point.x += point.dx
    point.y += point.dy
  }
}

func main() {
  fd, err := os.Open("day10.dat")
  check(err)

  dataReader := io.Reader(fd)
  points := readData(dataReader)

  dispersion := averageDistance(points)
  for averageDistance(points) <= dispersion {
    step(points)
    dispersion = averageDistance(points)
  }

  fmt.Println("Done... now we look at the message")
}
