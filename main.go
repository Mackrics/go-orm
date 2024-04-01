package main

import (
    "fmt"
    "flag"
    "math"
    "os"
    "html/template"
    "net/http"
    "strconv"
)

var tmpl *template.Template

func init() {
    // Parse the template file
    var err error
    tmpl, err = template.ParseFiles("template.html")
    if err != nil {
        panic(err)
    }
}

type PageData struct {
    Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Execute the template
    err := tmpl.Execute(w, PageData{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the submitted values and convert them to numeric values:

    weight := r.Form.Get("weight")
    reps := r.Form.Get("reps")
    formula := r.Form.Get("formula")

    weightn,_ := strconv.ParseFloat(weight, 64)
    repsn,_ := strconv.ParseFloat(reps, 64)
    eorm := calc_orm(weightn, repsn, 2, formula)
    // You can perform operations with the numeric values here
    // Convert the sum from float to string:
    eorms := strconv.FormatFloat(eorm, 'f', 2, 64)

    // Display a message
    message := "Your estimated one rep max is: " + eorms

    // Render the template with the message
    err = tmpl.Execute(w, PageData{Message: message})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func serve_site(port string) {
    // Setup HTTP routes
    http.HandleFunc("/", handler)
    http.HandleFunc("/submit", submitHandler)

    // Start the server
    http.ListenAndServe(port, nil)
}



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


func main() {
  weight    := flag.Float64("w", 0, "The weight (required)")
  reps      := flag.Float64("r", 0, "The number of repetitions (required)")
  precision := flag.Int("p", 2, "Precision of estimated orm")
  formula   := flag.String("f", "Epley", "The formula used to calculate the estimated ORM (Mayhew or Epley)")

  if len(os.Args) == 2 {
    if os.Args[1] == "serve" {
      serve_site(":6080")
    } 
  } else {
    flag.Parse()
    if *weight == 0 || *reps == 0 {
      flag.Usage()
      os.Exit(1)
    }
    eorm := calc_orm(*weight, *reps, *precision, *formula)
    fmt.Println(eorm)
  }
}
