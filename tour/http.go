package main

import (
	"fmt"
	"net/http"
)

type String string

func (this String) ServeHTTP(
	writer http.ResponseWriter,
	Request *http.Request) {
	fmt.Fprint(writer, this)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (this Struct) ServeHTTP(
	writer http.ResponseWriter,
	Request *http.Request) {
	fmt.Fprintf(
        writer, "%s %s %s",
        this.Greeting, this.Punct, this.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
