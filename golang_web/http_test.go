package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// example of testing http handler without starting a server
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello", nil)
	record := httptest.NewRecorder()

	HelloHandler(record, request)

	res := record.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.Status)
	fmt.Println(string(body))

}
