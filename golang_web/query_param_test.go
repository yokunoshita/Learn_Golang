package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// SayHello is an example handler to demonstrate query parameters
func SayHello(writter http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writter, "At least say your name !!")
	} else {
		fmt.Fprintf(writter, "Wzup %s", name)
	}
}

// example of query parameter
func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=Yoku", nil)
	record := httptest.NewRecorder()

	SayHello(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// example of multiple query parameters
func MultipleParams(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("firstName")
	lastname := request.URL.Query().Get("lastName")
	fmt.Fprintf(writer, "%s %s", firstname, lastname)
}

// multiple value query params
func MultipleParams2(writter http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query()
	var names []string = query["name"]
	fmt.Fprintln(writter, strings.Join(names, ","))
}

func TestMultipleParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/hello?firstName=Yoku&lastName=Dake", nil)
	record := httptest.NewRecorder()

	MultipleParams(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleParams2(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/hello?name=Yoku&name=Dake&name=Han", nil)
	record := httptest.NewRecorder()

	MultipleParams2(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
