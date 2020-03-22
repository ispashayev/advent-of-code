package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getBattleGrid(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	check(scanner.Err())
}

func main() {
	fd, err := os.Open("day15.dat")
	check(err)
	getBattleGrid(io.Reader(fd))
}
