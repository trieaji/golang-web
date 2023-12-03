package golangweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{ //membuat server
		Addr: "localhost:4003",
	}

	err := server.ListenAndServe() //menjalankan server
	if err != nil {
		panic(err)
	}
}