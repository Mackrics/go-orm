package main

import (
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

    number1 := r.Form.Get("number1")
    number2 := r.Form.Get("number2")

    number1n,_ := strconv.ParseFloat(number1, 64)
    number2n,_ := strconv.ParseFloat(number2, 64)
    eorm := calc_orm(number1n, number2n, 2, "Epley")
    // You can perform operations with the numeric values here
    // Convert the sum from float to string:
    eorms := strconv.FormatFloat(eorm, 'f', 2, 64)

    // how do I import orm.go?
    // a: you can't import it, but you can copy the functions from orm.go into this file.
    // can i make this a project somehow?
    // a: yes, you can create a new directory and put both files in it. Then you can run the program from the new directory.
    // how should the base file be structured?
    // a: you can structure it like this:
    //   - project
    //     - main.go
    //     - orm.go
    //     - template.html
    //     - go.mod
    //     - go.sum
    //     - README.md
    //     - LICENSE
    //     - .gitignore
    //     - .dockerignore
    // what is project?


    // Display a message
    message := "Numbers submitted: Number 1 = " + number1 + ", Number 2 = " + number2 + ", Sum = " + eorms

    // Render the template with the message
    err = tmpl.Execute(w, PageData{Message: message})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // Setup HTTP routes
    http.HandleFunc("/", handler)
    http.HandleFunc("/submit", submitHandler)

    // Start the server
    http.ListenAndServe(":8080", nil)
}

