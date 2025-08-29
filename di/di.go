package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// As bytes.Buffer implements we can pass the io.Writer interface to accept
// anything that implements the Write(p []byte) (n int, err error) function from io.Writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implement io.Writer
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Gwen")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
