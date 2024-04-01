package main

import (
    "fmt"
    "flag"
    "os"
    "math"
)

func roundFloat(number float64, decimals int) float64 {
  factor := math.Pow(10, float64(decimals))
  roundedNumber := math.Round(number*factor) / factor
  return roundedNumber
}

func calc_orm(weight float64, reps float64, precision int, formula string) float64 {
  var prettyEORM float64
  if formula == "Epley" {
    eorm := weight * (1 + reps/30)
    prettyEORM = roundFloat(eorm, precision)
  } else if formula == "Mayhew" {
    eorm := weight * ( 1 / (0.55 + (0.419 * math.Exp((-0.055 * reps)))))
    prettyEORM = roundFloat(eorm, precision)
  } else {
    flag.Usage()
    os.Exit(1)
  }
  return(prettyEORM)
}

// q: why is prettyEROM not defined?
// a: because it is defined in the if block, and not available outside of it.
// q: how do I fix this?


func main() {
  weight    := flag.Float64("w", 0, "The weight (required)")
  reps      := flag.Float64("r", 0, "The number of repetitions (required)")
  precision := flag.Int("p", 2, "Precision of estimated orm")
  formula   := flag.String("f", "Epley", "The formula used to calculate the estimated ORM (Mayhew or Epley)")
  
  flag.Parse()

  if *weight == 0 || *reps == 0 {
    flag.Usage()
    os.Exit(1)
  }
  eorm := calc_orm(*weight, *reps, *precision, *formula)
  fmt.Println(eorm)
}
