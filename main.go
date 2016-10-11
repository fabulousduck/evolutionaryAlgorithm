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

func main() {
  algorithm := newAlorithm()
  dataPoints := newPointSet(20);
  algorithm.putRandomCandidateSolutions(dataPoints, 50)
}

func newAlorithm() *algorithm {
  return &algorithm{
    [][]point{},
  }
}

func (algorithm *algorithm) putRandomCandidateSolutions(dataPoints []point, amount int){
  dp := dataPoints
  for i := 0; i < amount; i++ {
    fmt.Println("randomized : ", randomSolution(dp))
  }

}

func randomSolution(dataPoints []point) []point {
  startPoint := dataPoints[0]
  solution := []point{startPoint}
  dataPoints = append(dataPoints[:0], dataPoints[1:]...)
  for i := 0; i < len(dataPoints); i++{
    randomPos := rand.Intn(len(dataPoints));
    solution = append(solution, dataPoints[randomPos])
    dataPoints = append(dataPoints[:randomPos],dataPoints[randomPos:]...)
  }
  solution = append(solution, startPoint)
  return solution
}

func newPointSet(amount int) []point {
  points := []point{}
  for i := 0; i < amount; i++ {
    point1 := rand.Intn(100)
    point2 := rand.Intn(100)
    points = append(points,point{point1,point2})
  }
  return points
}
