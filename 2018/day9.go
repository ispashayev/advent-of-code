package main

import (
  "fmt"
  "io"
  "os"
)


func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  fd, err := os.Open("day9.dat")
  check(err)
}
