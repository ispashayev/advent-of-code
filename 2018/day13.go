package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var MAX_TICK = 1000000

type Cart struct {
	direction         rune
	turnOptionCounter int
	stepped           bool
	crashed           bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func resetCartSteps(carts []*Cart) {
	for _, cart := range carts {
		cart.stepped = false
	}
}

func numCartsLeft(carts []*Cart) (cartsLeft int) {
	for _, cart := range carts {
		if !cart.crashed {
			cartsLeft++
		}
	}
	return cartsLeft
}

func getCartsAndTurns(reader io.Reader) (carts []*Cart, cartMap [][]*Cart, trackMap [][]rune) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	directions := "<>^v"
	turns := "/\\+"

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		currentRowTurns := make([]rune, len(line))
		currentRowCarts := make([]*Cart, len(line))
		for x, state := range line {
			if state == ' ' {
				continue
			} else if strings.ContainsRune(directions, state) {
				newCart := Cart{state, 0, false, false}
				carts = append(carts, &newCart)
				currentRowCarts[x] = &newCart
			} else if strings.ContainsRune(turns, state) {
				currentRowTurns[x] = state
			}
		}
		cartMap = append(cartMap, currentRowCarts)
		trackMap = append(trackMap, currentRowTurns)
	}

	check(scanner.Err())

	return carts, cartMap, trackMap
}

func getNextCoordinate(direction rune, x int, y int) (int, int) {
	if direction == '>' {
		return x + 1, y
	} else if direction == '<' {
		return x - 1, y
	} else if direction == '^' {
		return x, y - 1
	} else if direction == 'v' {
		return x, y + 1
	}
	panic("Invalid cart direction")
}

func (cart *Cart) checkAndSetCartDirection(turn rune) {
	if turn == rune(0) {
		return
	} else if turn == '/' {
		if cart.direction == '>' {
			cart.direction = '^'
		} else if cart.direction == '<' {
			cart.direction = 'v'
		} else if cart.direction == '^' {
			cart.direction = '>'
		} else if cart.direction == 'v' {
			cart.direction = '<'
		}
	} else if turn == '\\' {
		if cart.direction == '>' {
			cart.direction = 'v'
		} else if cart.direction == '<' {
			cart.direction = '^'
		} else if cart.direction == '^' {
			cart.direction = '<'
		} else if cart.direction == 'v' {
			cart.direction = '>'
		}
	} else if turn == '+' {
		if cart.turnOptionCounter == 0 {
			// turn left
			if cart.direction == '>' {
				cart.direction = '^'
			} else if cart.direction == '<' {
				cart.direction = 'v'
			} else if cart.direction == '^' {
				cart.direction = '<'
			} else if cart.direction == 'v' {
				cart.direction = '>'
			}
		} else if cart.turnOptionCounter == 2 {
			// turn right
			if cart.direction == '>' {
				cart.direction = 'v'
			} else if cart.direction == '<' {
				cart.direction = '^'
			} else if cart.direction == '^' {
				cart.direction = '>'
			} else if cart.direction == 'v' {
				cart.direction = '<'
			}
		}
		cart.turnOptionCounter = (cart.turnOptionCounter + 1) % 3
	}
}

func getFirstCrash(carts []*Cart, cartMap [][]*Cart, trackMap [][]rune) (int, int, int) {
	for tick := 0; tick < MAX_TICK; tick++ {
		resetCartSteps(carts)
		for y, currentRowCarts := range cartMap {
			for x, cart := range currentRowCarts {
				if cart == nil || cart.stepped {
					continue
				}
				// advance the cart along its direction
				nextX, nextY := getNextCoordinate(cart.direction, x, y)
				if cartMap[nextY][nextX] != nil {
					// cart crash
					return tick, nextX, nextY
				}
				cartMap[y][x] = nil
				cartMap[nextY][nextX] = cart
				cart.checkAndSetCartDirection(trackMap[nextY][nextX])
				cart.stepped = true
			}
		}
	}
	panic("Unable to get first crash")
}

func getLastCart(carts []*Cart, cartMap [][]*Cart, trackMap [][]rune) (int, int, int) {
	for tick := 0; tick < MAX_TICK; tick++ {
		resetCartSteps(carts)
		var lastCartX, lastCartY int
		for y, currentRowCarts := range cartMap {
			for x, cart := range currentRowCarts {
				if cart == nil || cart.stepped {
					continue
				}
				// advance the cart along its direction
				nextX, nextY := getNextCoordinate(cart.direction, x, y)
				if cartMap[nextY][nextX] != nil {
					// cart crash
					cart.crashed = true
					cartMap[nextY][nextX].crashed = true
					cartMap[y][x] = nil
					cartMap[nextY][nextX] = nil
					continue
				}
				cartMap[y][x] = nil
				cartMap[nextY][nextX] = cart
				cart.checkAndSetCartDirection(trackMap[nextY][nextX])
				cart.stepped = true
				// Set metadata for simulation termination
				lastCartX, lastCartY = nextX, nextY
			}
		}
		if numCartsLeft(carts) == 1 {
			return tick, lastCartX, lastCartY
		}
	}
	panic("Unable to get last cart")
}

func main() {
	fd, err := os.Open("day13.dat")
	check(err)
	carts, cartMap, trackMap := getCartsAndTurns(io.Reader(fd))

	var tick, x, y int

	tick, x, y = getFirstCrash(carts, cartMap, trackMap)
	fmt.Printf("First crash at tick=%d, x=%d, y=%d\n", tick+1, x, y)

	// reload the input data
	_, err = fd.Seek(0, 0)
	check(err)
	carts, cartMap, trackMap = getCartsAndTurns(io.Reader(fd))

	tick, x, y = getLastCart(carts, cartMap, trackMap)
	fmt.Printf("Last cart at tick=%d, x=%d, y=%d\n", tick+1, x, y)
}
