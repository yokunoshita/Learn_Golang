package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// CallServer is a helper function to start an HTTP server with the given address and handler
func CallServer(address string, handler http.Handler) {
	server := http.Server{
		Addr:    address,
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	CallServer("localhost:8000", handler)
}

// handler is not supported to have multiple endpoint urls
// instead we can use http.ServeMux to handle multiple endpoint urls

func TestServerMux(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Wzup")
	})
	mux.HandleFunc("/hello/yoku", func(writer http.ResponseWriter, request *http.Request) { // this is called url pattern
		fmt.Fprint(writer, "Wzup yoku")
	})

	CallServer("localhost:4000", mux)
}

// example to read request method and request uri
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, r.Method)
		fmt.Println(w, r.RequestURI)

	}
	CallServer("localhost:2000", handler)
}
