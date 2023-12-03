package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"//berguna untuk menjalankan http di dalam unit test
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Giyuu")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4003/hailo", nil)//NewRequest merupakan representasi dari "request *http.Request"
	recorder := httptest.NewRecorder()//NewRecorder merupakan representasi dari "writer http.ResponseWriter"

	HelloHandler(recorder, request)

	//untuk mengecek hasil test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}