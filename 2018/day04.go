package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "sort"
  "strconv"
  "strings"
  "time"
)

type Record struct {
  timestamp time.Time
  id int
  action []string
}

type By func(r1, r2 *Record) bool

func (by By) Sort(records []Record) {
  rs := &recordSorter{
    records: records,
    by: by, // The Sort method's receiver is the function (closure) that defines the sort order.
  }
  sort.Sort(rs)
}

// recordSorter joins a By function and a slice of Records to be sorted.
type recordSorter struct {
  records []Record
  by func(r1, r2 *Record) bool // Closure used in the Less method
}

// Len is part of sort.Interface.t
func (s *recordSorter) Len() int {
  return len(s.records)
}

// Swap is part of sort.Interface
func (s *recordSorter) Swap(i, j int) {
  s.records[i], s.records[j] = s.records[j], s.records[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *recordSorter) Less(i, j int) bool {
  return s.by(&s.records[i], &s.records[j])
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func parseLog(logReader io.Reader) ([]Record, error) {
  var records []Record

  logScanner := bufio.NewScanner(logReader)
  logScanner.Split(bufio.ScanLines)

  for logScanner.Scan() {
    recordLine := strings.Split(logScanner.Text(), " ")
    
    year, err := strconv.Atoi(recordLine[0][1:5])
    check(err)
    month, err := strconv.Atoi(recordLine[0][6:8])
    check(err)
    day, err := strconv.Atoi(recordLine[0][9:11])
    check(err)
    hour, err := strconv.Atoi(recordLine[1][0:2])
    check(err)
    minute, err := strconv.Atoi(recordLine[1][3:5])
    check(err)
    timestamp := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

    id, action := -1, recordLine[2:]
    if recordLine[2] == "Guard" {
      id, err = strconv.Atoi(recordLine[3][1:])
      check(err)
      action = recordLine[4:]
    }
    new_record := Record{timestamp, id, action}
    records = append(records, new_record)
  }

  return records, logScanner.Err()
}

func Sum(a []int) (sum int) {
  for _, a_i := range a {
    sum += a_i
  }
  return sum
}

func ArgMax(a []int) (argmax int) {
  for i, a_i := range a {
    if a_i > a[argmax] {
      argmax = i
    }
  }
  return argmax
}

func computeSleepDistribution(records []Record) (sleepDistribution map[int][]int) {
  currentGuardId := -1
  sleepStart, sleepEnd := -1, -1
  sleepDistribution = make(map[int][]int)
  for _, record := range records {
    if record.id != -1 {
      currentGuardId = record.id
      if sleepDistribution[currentGuardId] == nil {
        sleepDistribution[currentGuardId] = make([]int, 60)
      }
    }
    action := strings.Join(record.action, " ")
    if action == "falls asleep" {
      sleepStart = record.timestamp.Minute()
    }
    if action == "wakes up" {
      sleepEnd = record.timestamp.Minute()
      for i := sleepStart; i < sleepEnd; i++ {
        sleepDistribution[currentGuardId][i]++
      }
    }
  }

  return sleepDistribution
}

func computeStrategyOne(sleepDistribution map[int][]int) {
  fmt.Println("STRATEGY ONE DATA:")
  fmt.Println("=================")

  sleepiestGuardId, mostMinutesSlept := -1, -1
  for guardId, minuteDistribution := range sleepDistribution {
    minutesSlept := Sum(minuteDistribution)
    if minutesSlept > mostMinutesSlept {
      sleepiestGuardId = guardId
      mostMinutesSlept = minutesSlept
    }
  }

  minuteMostSlept := ArgMax(sleepDistribution[sleepiestGuardId])
  fmt.Println("The sleepiest guard is", sleepiestGuardId)
  fmt.Println("He slept for", mostMinutesSlept, "minutes")
  fmt.Println("Minute most slept:", minuteMostSlept)
  fmt.Println()
}

func computeStrategyTwo(sleepDistribution map[int][]int) {
  fmt.Println("STRATEGY TWO DATA:")
  fmt.Println("=================")

  mostPredictableGuardId, minute, timesSlept := -1, -1, -1
  for guardId, minuteDistribution := range sleepDistribution {
    minuteMostSlept := ArgMax(minuteDistribution)
    if minuteDistribution[minuteMostSlept] > timesSlept {
      mostPredictableGuardId = guardId
      minute = minuteMostSlept
      timesSlept = minuteDistribution[minuteMostSlept]
    }
  }

  fmt.Println("The most predictable guard is", mostPredictableGuardId)
  fmt.Println("He slept on the same minute", timesSlept, "times")
  fmt.Println("That minute was minute", minute)
}

func main() {
  fd, err := os.Open("day04.dat")
  check(err)

  logReader := io.Reader(fd)
  records, err := parseLog(logReader)
  check(err)

  // Closure that orders the Record structure
  timestamp := func(r1, r2 *Record) bool {
    one, two := r1.timestamp, r2.timestamp
    if one.Year() != two.Year() {
      return one.Year() < two.Year()
    }
    if one.Month() != two.Month() {
      return one.Month() < two.Month()
    }
    if one.Day() != two.Day() {
      return one.Day() < two.Day()
    }
    if one.Hour() != two.Hour() {
      return one.Hour() < two.Hour()
    }
    if one.Minute() != two.Minute() {
      return one.Minute() < two.Minute()
    }
    return false
  }

  By(timestamp).Sort(records)
  sleepDistribution := computeSleepDistribution(records)

  computeStrategyOne(sleepDistribution)
  computeStrategyTwo(sleepDistribution)

}
