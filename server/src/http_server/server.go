package http_server

import (
	"fmt"
	"io"
	"net/http"
)

type HttpUrlsHandler struct {
	server http.Server
}

func addHelloHandler(m *http.ServeMux) {
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Println("Hello was called")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello to you too!")
	})
}

// TODO pass settings, probably store settings
func (h *HttpUrlsHandler) Initialize() {
	mux := http.NewServeMux()

	addHelloHandler(mux)

	h.server = http.Server{Addr: ":8080", Handler: mux}
}

func (h *HttpUrlsHandler) Start() error {
	return h.server.ListenAndServe()
}
