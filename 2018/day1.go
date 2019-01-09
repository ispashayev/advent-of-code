package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// ReadInts reads white-space separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanWords)
  var result []int
  for scanner.Scan() {
    x, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return result, err
    }
    result = append(result, x)
  }
  return result, scanner.Err()
}

// ComputeTotalFrequency returns the sum of all the values in the `numbers` array.
func ComputeTotalFrequency(numbers []int) (sum int) {
  for i := 0; i < len(numbers); i++ {
    sum += numbers[i]
  }
  return sum
}

func FindFirstRepeatedFrequency(numbers []int) (int) {
  i, frequency, visited_frequencies := 0, 0, make(map[int]bool)
  for {
    frequency += numbers[i]
    if visited_frequencies[frequency] {
      return frequency
    }
    visited_frequencies[frequency] = true
    i = (i + 1) % len(numbers)
  }
}

func main() {
  fd, err := os.Open("day1.dat")
  check(err)
  ioReader := io.Reader(fd)
  ints, err := ReadInts(ioReader)
  check(err)

  // Part One - sum all numbers
  totalFrequency := ComputeTotalFrequency(ints)
  fmt.Println("The total of all the number is", totalFrequency)

  // Part Two - find the first repeated "frequency" (i.e. running total)
  firstRepeatedFrequency := FindFirstRepeatedFrequency(ints)
  fmt.Println("The first repeated frequency is", firstRepeatedFrequency)
}
