package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x  float64
	y  float64
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

func boundingBox(points []*Point) (int, int) {
	minX, minY, maxX, maxY := points[0].x, points[0].y, points[0].x, points[0].y
	for _, point := range points {
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	return int(maxX-minX) + 1, int(maxY-minY) + 1
}

func stepForward(points []*Point) {
	for _, point := range points {
		point.x += point.dx
		point.y += point.dy
	}
}

func stepBack(points []*Point) {
	for _, point := range points {
		point.x -= point.dx
		point.y -= point.dy
	}
}

func display(points []*Point) {
	minX, minY, maxX, maxY := points[0].x, points[0].y, points[0].x, points[0].y
	for _, point := range points {
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	dimX, dimY := int(maxX-minX)+1, int(maxY-minY)+1
	fmt.Printf("Displaying points in %d-by-%d coordinate grid.\n\n", dimX, dimY)

	grid := make([][]byte, dimY)
	for y := 0; y < dimY; y++ {
		grid[y] = make([]byte, dimX)
		for x := 0; x < dimX; x++ {
			grid[y][x] = '.'
		}
	}

	for _, point := range points {
		y := int(point.y - minY)
		x := int(point.x - minX)
		grid[y][x] = '#'
	}

	for _, row := range grid {
		for _, chrByte := range row {
			fmt.Printf("%s", string(chrByte))
		}
		fmt.Println()
	}
}

func main() {
	fd, err := os.Open("day10.dat")
	check(err)

	dataReader := io.Reader(fd)
	points := readData(dataReader)
	fmt.Printf("%d points loaded.\n", len(points))
	fmt.Printf("Simulating point steps until they coalesce maximally...\n\n")

	x, y := boundingBox(points)
	area := x * y
	var i, previousArea int
	for {
		fmt.Printf("The area covered by the points is %d.\n", area)
		previousArea = area
		stepForward(points)
		x, y = boundingBox(points)
		area = x * y
		if area > previousArea {
			break
		}
		i++
	}

	stepBack(points)

	fmt.Printf("Done in %d seconds.\n\n", i)
	display(points)
}
