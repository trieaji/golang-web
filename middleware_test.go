package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func(middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
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
		Handler: &logMiddleware,
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