package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() //parsing terlebih dahulu
	if err != nil {
		panic(err)
	}

	// request.PostFormValue("first_name")

	// ambil yang sudah di parsing
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=kamado&last_name=tanjirou")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4003", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") //penulisan sudah paten

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}