package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//header adalah informasi tambahan yg biasa dikirim dari client ke server atau sebaliknya

// cara menangkap header dari client, yang dikirim oleh si client
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type") //"content-type" adalah header atau informasi tambahan yang dikirim oleh si client
	fmt.Fprint(writer, contentType)
}

// cara menangkap header dari client, yang dikirim oleh si client
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4003/", nil)
	request.Header.Add("Content-Type", "meliodas") //application/json -> adalah header yang kita kirim

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// cara meresponse header dari client yang dikirim oleh si client
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("dipersembahkan-oleh", "nanatsu no taizai") //digunakan untuk menambahkan header pada response
	fmt.Fprint(writer, "OK")
}

// cara meresponse header dari client yang dikirim oleh si client
func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4003/", nil)
	request.Header.Add("Content-Type", "meliodas")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	//cek header response nya
	fmt.Println(response.Header.Get("dipersembahkan-oleh"))
}
