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
  "time"
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
  cityPositions := newCities(20);
  algorithm.putRandomCandidateSolutions(cityPositions, 50)
}

func newAlorithm() *algorithm {
  return &algorithm{
    [][]point{},
  }
}

func (algorithm *algorithm) putRandomCandidateSolutions(cityPositions []point, amount int){
  for i := 0; i < amount; i++ {
    algorithm.candidateSolutions = append(algorithm.candidateSolutions, randomSolution(cityPositions))
  }

}

func randomSolution(cities []point) []point {

  localCitySliceCopy := make([]point, len(cities))
  copy(localCitySliceCopy, cities)
  solution := []point{cp[0]}
  cities = append(localCitySliceCopy[:0], localCitySliceCopy[1:]...)

  for i := 0; i < len(localCitySliceCopy); i++ {
    randomPoint := randInt(0, len(localCitySliceCopy))
    solution = append(solution, localCitySliceCopy[randomPoint])
    cp = append(localCitySliceCopy[:randomPoint],localCitySliceCopy[randomPoint+1:]...)
  }
  solution = append(solution, solution[0])
  return solution

}

func newPointSet(amount int) []point {
  points := []point{}
  for i := 0; i < amount; i++ {
    point1 := randInt(0,100)
    point2 := randInt(0,100)
    points = append(points,point{point1,point2})
  }
  return points
}

func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max-min)
}
