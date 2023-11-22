package golangweb

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
	"net/http/httptest"
	"io"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
		))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "template data map",
		"Name" : "sukuna",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4003", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}	