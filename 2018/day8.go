package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
)


func check(err error) {
  if err != nil {
    panic(err)
  }
}

func parseNumbers(r io.Reader) ([]int, error) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanWords)

  var numbers []int
  for scanner.Scan() {
    number, err := strconv.Atoi(scanner.Text())
    check(err)
    numbers = append(numbers, number)
  }

  return numbers, scanner.Err()
}

func accumulateMetadata(stream []int, streamPosition int) (int, int) {
  numChildren, numMetadata := stream[streamPosition], stream[streamPosition+1]
  streamPosition += 2

  childMetadataSum, cumulativeMetadataSum := 0, 0

  for i := 0; i < numChildren; i++ {
    streamPosition, childMetadataSum = accumulateMetadata(stream, streamPosition)
    cumulativeMetadataSum += childMetadataSum
  }

  for i := 0; i < numMetadata; i++ {
    cumulativeMetadataSum += stream[streamPosition]
    streamPosition++
  }

  return streamPosition, cumulativeMetadataSum
}

func main() {
  fd, err := os.Open("day8.dat")
  check(err)

  numReader := io.Reader(fd)
  numStream, err := parseNumbers(numReader)
  
  streamPosition, cumulativeMetadataSum := accumulateMetadata(numStream, 0)
  if streamPosition != len(numStream) {
    panic("Final stream position is not full length of stream.")
  }

  fmt.Println("The cumulative metadata sum is:", cumulativeMetadataSum)
}
