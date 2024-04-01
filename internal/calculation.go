package main

import (
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
