package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/mackrics/go-orm/internal/calculation"
    "github.com/mackrics/go-orm/internal/fronend"
)

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
