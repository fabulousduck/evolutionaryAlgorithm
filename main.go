/*
this algorithm will attempt to solve the traveling salesman problem

Evolution is based on survival of the fittest.
there fore :

  Finite resources + Life froms with basic intinct towards reproduction = natural selection

there are 2 sorts of evolution :
  1. recombination  > take 2 inputs and combine to make 1 output
  2. mutation       > take 2 inputs and generate a completely new output

step by step :

  1. initialize pupulation with random candidate solutions

    if there are 6 cities to go to, a candidate solution might be : 1>2>4>3>6>5>1

    this enables :
      calculating the destance (fitness function)
      mutation
      recombination


*/
package main

import(
  "fmt"
  "math/rand"
)

type point struct {
  x int
  y int
}

type algorithm struct {
  candidateSolutions [][]point
}

type opts struct {
  points int
  initCandidateSolutionsCount int
}

func main() {
  opts := newOpts(20, 1000)
  algorithm := opts.newAlorithm()
  dataPoints := newPointSet(opts.points);
  algorithm.putRandomCandidateSolutions(dataPoints, opts.initCandidateSolutionsCount)
  fmt.Println(algorithm.candidateSolutions)
}

func newOpts(points int, icsc int) *opts {
  return &opts{
      points,
      icsc,
  }
}

func (opts *opts)newAlorithm() *algorithm {
  return &algorithm{
    make([][]point{},opts.initCandidateSolutionsCount),
  }
}

func (algorithm *algorithm) putRandomCandidateSolutions(dataPoints []point, amount int){
  for i := 0; i < amount; i++ {
    algorithm.candidateSolutions = append(algorithm.candidateSolutions, randomSolution(dataPoints))
  }
}

func randomSolution(dataPoints []point) []point {
  candidateSolution := []point{dataPoints[0]}
  dataPoints = append(dataPoints[:0], dataPoints[1:]...)
  for i := 0; i < len(dataPoints)-2; i++ {
    candidateSolution = append(candidateSolution, dataPoints[rand.Intn(len(dataPoints))]);

  }
  candidateSolution = append(candidateSolution, candidateSolution[0]);
  return candidateSolution
}

func newPointSet(amount int) []point {
  points := []point{}
  for i := 0; i < amount; i++ {
    points = append(points,point{rand.Intn(100),rand.Intn(100)})
  }
  return points
}
