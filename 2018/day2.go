package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readIds(r io.Reader) (identifiers []string, err error) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    word := scanner.Text()
    identifiers = append(identifiers, word)
  }
  return identifiers, scanner.Err()
}

func checksum(identifiers []string) int {
  num_doubles, num_triples := 0, 0
  for _, id := range identifiers {
    character_map := make(map[rune]int)
    for _, c := range id {
      character_map[c]++
    }
    for _, v := range character_map {
      if v == 2 {
        num_doubles++
        break
      }
    }
    for _, v := range character_map {
      if v == 3 {
        num_triples++
        break
      }
    }
  }
  return num_doubles * num_triples
}

func getStringPair(identifiers []string) (string, string) {
  for i, base_id := range identifiers[:len(identifiers)-1] {
    for _, reference_id := range identifiers[i+1:] {
      mismatch_count := 0
      for j := 0; j < len(base_id); j++ {
        if base_id[j] != reference_id[j] {
          mismatch_count++
          if mismatch_count > 1 {
            break
          }
        }
      }
      if mismatch_count <= 1 {
        return base_id, reference_id
      }
    }
  }
  panic("Unable to find string pair.")
}


func main() {
  fd, err := os.Open("day2.dat")
  check(err)
  ioReader := io.Reader(fd)
  identifiers, err := readIds(ioReader)
  check(err)
  
  checksum := checksum(identifiers)
  fmt.Println("The checksum is:", checksum)

  common_one, common_two := getStringPair(identifiers)
  var common string
  for i := 0; i < len(common_one); i++ {
    if common_one[i] == common_two[i] {
      common += string(common_one[i])
    }
  }
  fmt.Println("The longest common substring is:", common)
}
