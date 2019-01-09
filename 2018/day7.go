package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
)

var jobNames = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numWorkers = 5
var baseTimeRequired = 60

type Job struct {
  name byte
  dependencies []*Job
  timeToCompletion int
  finished bool
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func insertDependencyOrdered(job, predecessor *Job) {
  dependencies := job.dependencies

  i := 0
  for i < len(dependencies) && dependencies[i].name < predecessor.name {
    i++
  }

  job.dependencies = append(dependencies[:i], append([]*Job{predecessor}, dependencies[i:]...)...)
}

func parseDependencies(r io.Reader, jobs map[byte]*Job) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    tokens := strings.Split(scanner.Text(), " ")
    predecessor, job := jobs[tokens[1][0]], jobs[tokens[7][0]]
    insertDependencyOrdered(job, predecessor)
  }
}

func traverseDependencies(jobs map[byte]*Job) string {
  var jobSequence []string
  i, visited := 0, make(map[byte]bool)
  for i < len(jobNames) {
    jobName := jobNames[i]
    if !visited[jobName] {
      allDependenciesMet := true
      for _, dependency := range jobs[jobName].dependencies {
        if !visited[dependency.name] {
          allDependenciesMet = false
          break
        }
      }
      if allDependenciesMet {
        jobSequence = append(jobSequence, string(jobName))
        visited[jobName] = true
        i = 0
        continue
      }
    }
    i++
  }
  return strings.Join(jobSequence, "")
}


func completeJob(jobQueue []*Job) ([]*Job, *Job) {
  if len(jobQueue) == 0 {
    panic("Cannot complete job from an empty queue")
  }

  mindex := 0
  for i := 1; i < len(jobQueue); i++ {
    if jobQueue[i].timeToCompletion < jobQueue[mindex].timeToCompletion {
      mindex = i
    }
  }

  completedJob := jobQueue[mindex]
  completedJob.finished = true
  jobQueue = append(jobQueue[:mindex], jobQueue[mindex+1:]...)
  for _, job := range jobQueue {
    job.timeToCompletion -= completedJob.timeToCompletion
  }

  return jobQueue, completedJob
}


func scheduleJob(jobSequence string, jobs map[byte]*Job) *Job {
  for i := 0; i < len(jobSequence); i++ {
    job, allDependenciesMet := jobs[jobSequence[i]], true
    
    if job.finished || job.timeToCompletion > 0 {
      continue
    }

    for _, dependency := range job.dependencies {
      if !dependency.finished {
        allDependenciesMet = false
        break
      }
    }

    if allDependenciesMet {
      job.timeToCompletion = baseTimeRequired + int(job.name-'A') + 1
      return job
    }
  }

  return nil
}

func simulateJobCompletion(jobSequence string, jobs map[byte]*Job) (string, int) {
  timeElapsed, completedSequence := 0, make([]string, 0)

  var jobQueue []*Job
  var nextJob, completedJob *Job

  for {
    // Schedule a job, completing queued jobs if necessary
    nextJob = scheduleJob(jobSequence, jobs)
    for nextJob == nil {
      jobQueue, completedJob = completeJob(jobQueue)
      timeElapsed += completedJob.timeToCompletion
      completedSequence = append(completedSequence, string(completedJob.name))

      if len(completedSequence) == len(jobSequence) {
        return strings.Join(completedSequence, ""), timeElapsed
      }

      nextJob = scheduleJob(jobSequence, jobs)
    }

    // If all workers are busy, wait until the next job completes
    if len(jobQueue) == numWorkers {
      jobQueue, completedJob = completeJob(jobQueue)
      timeElapsed += completedJob.timeToCompletion
      completedSequence = append(completedSequence, string(completedJob.name))

      if len(completedSequence) == len(jobSequence) {
        return strings.Join(completedSequence, ""), timeElapsed
      }
    }

    jobQueue = append(jobQueue, nextJob)
  }

  // Flush out the job queue
  for len(jobQueue) > 0 {
    jobQueue, completedJob = completeJob(jobQueue)
    timeElapsed += completedJob.timeToCompletion
    completedSequence = append(completedSequence, string(completedJob.name))
  }

  return strings.Join(completedSequence, ""), timeElapsed
}

func main() {
  jobs := make(map[byte]*Job, len(jobNames))
  for i := 0; i < len(jobNames); i++ {
    job := Job{jobNames[i], make([]*Job, 0), -1, false}
    jobs[jobNames[i]] = &job
  }

  fd, err := os.Open("day7.dat")
  check(err)

  dataReader := io.Reader(fd)
  parseDependencies(dataReader, jobs)
  jobSequence := traverseDependencies(jobs)
  fmt.Println("The job order is:", jobSequence)

  completedSequence, timeElapsed := simulateJobCompletion(jobSequence, jobs)
  fmt.Println("The jobs are completed in the following order:", completedSequence)
  fmt.Println("The time elapsed is", timeElapsed)
}
