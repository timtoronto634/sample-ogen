package main

import "net/http"

func main() {
	s, err := newServer()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
