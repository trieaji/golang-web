package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

/* 
1. Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi
2. Kadang middleware juga bisa digunakan untuk melakukan error handler
3. Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan mengubah panic tersebut menjadi panic error response
4. Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan
*/

type LogMiddleware struct {
	Handler http.Handler
}

func(middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {//adanya pointer di "*LogMiddleware" supaya datanya tidak diduplikat terus-menerus
	fmt.Println("before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() { //defer adalah function yg bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi, defer function akan selalu di eksekusi walaupun terjadi error di function yg di eksekusi
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello middleware")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Foo Executed")
		panic("Ups")	
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	//error handler
	errorHandler := &ErrorHandler{
		Handler: &logMiddleware,//diberi pointer(&) karena "LogMiddleware" nya adalah function Handler
	}

	server := http.Server{
		Addr: "localhost:4003",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}