package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Println("Hello was called")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello to you too!")
	})

	fmt.Println("Well, something happened")

	s := http.Server{Addr: ":8080", Handler: mux}

	err := s.ListenAndServe()
	fmt.Println(err)
}
