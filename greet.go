package main

import (
    "fmt"
    "io"
    // "log"
    "net/http"
)

func Greets(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greets(w, "world")
}

// func main() {
//     log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }