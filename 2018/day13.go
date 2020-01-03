package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Cart struct {
	direction         rune
	turnOptionCounter int
	stepped           bool
	startX            int
	startY            int
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
				newCart := Cart{state, 0, false, x, y}
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

func getNextCoordinate(direction rune, x int, y int, cart *Cart) (int, int) {
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

func stepAndCheckCrash(cartMap [][]*Cart, trackMap [][]rune) bool {
	for y, currentRowCarts := range cartMap {
		for x, cart := range currentRowCarts {
			if cart == nil || cart.stepped {
				continue
			}
			// advance the cart along its direction
			nextX, nextY := getNextCoordinate(cart.direction, x, y, cart)
			if cartMap[nextY][nextX] != nil {
				fmt.Printf("Crash at x=%d, y=%d\n", nextX, nextY)
				return true
			}
			cartMap[y][x] = nil
			cartMap[nextY][nextX] = cart
			cart.checkAndSetCartDirection(trackMap[nextY][nextX])
			cart.stepped = true
		}
	}
	return false
}

func main() {
	fd, err := os.Open("day13.dat")
	check(err)
	carts, cartMap, trackMap := getCartsAndTurns(io.Reader(fd))

	var tick int
	for tick = 0; !stepAndCheckCrash(cartMap, trackMap); tick++ {
		resetCartSteps(carts)
	}
	fmt.Printf("Crashed on tick %d\n", tick+1)
}
