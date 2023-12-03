package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {//handler berguna untuk menerima HTTP Request yg masuk ke server
		fmt.Fprint(writer, "Hello World")//pointnya tugas heandler yaitu menerima request yang masuk ke server
		//writer berfungsi untuk mengembalikan response dari si client
		//request berfungsi untuk menerima apa yg dikirim dari si client
	}

	server := http.Server{
		Addr: "localhost:4003",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) { //servemux adalah implementasi handler yg bisa mendukung multiple endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	server := http.Server{
		Addr: "localhost:4003",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//request berguna untuk mengetahui semua informasi yg dikirim oleh web browser