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

func computeValue(stream []int, streamPosition int) (int, int) {
  numChildren, numMetadata := stream[streamPosition], stream[streamPosition+1]
  streamPosition += 2

  value := 0

  if numChildren == 0 {
    for i := 0; i < numMetadata; i++ {
      value += stream[streamPosition]
      streamPosition++
    }
  } else {
    childValues := make([]int, numChildren)
    for i := 0; i < numChildren; i++ {
      streamPosition, childValues[i] = computeValue(stream, streamPosition)
    }
    for i := 0; i < numMetadata; i++ {
      childIndex := stream[streamPosition] - 1
      if 0 <= childIndex && childIndex < len(childValues) {
        value += childValues[childIndex]
      }
      streamPosition++
    }
  }

  return streamPosition, value
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

  streamPosition, rootValue := computeValue(numStream, 0)
  if streamPosition != len(numStream) {
    panic("Final stream position is not full length of stream.")
  }

  fmt.Println("The value of the root node is:", rootValue)
}
