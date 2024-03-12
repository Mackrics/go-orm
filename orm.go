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

func main() {

  weight   := flag.Float64("w", 0, "The weight (required)")
  reps     := flag.Float64("r", 0, "The number of repetitions (required)")
  precision := flag.Int("p", 2, "Precision of estimated orm")
  formula  := flag.String("f", "Epley", "The formula used to calculate the estimated ORM (Mayhew or Epley)")
  
  flag.Parse()

  if *weight == 0 || *reps == 0 {
    flag.Usage()
    os.Exit(1)
  }


  if *formula == "Epley" {
    eorm := *weight * (1 + *reps/30)
    prettyEORM := roundFloat(eorm, *precision)
    fmt.Println(prettyEORM)
  } else if *formula == "Mayhew" {
    eorm := *weight * ( 1 / (0.55 + (0.419 * math.Exp((-0.055 * *reps)))))
    prettyEORM := roundFloat(eorm, *precision)
    fmt.Println(prettyEORM)
  } else {
    flag.Usage()
    os.Exit(1)
  }

}
