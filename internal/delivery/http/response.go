package http

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const (
	layoutTemplateFilename        string = "layout/layout"
	viewGetSquareTemplateFilename string = "views/square/get-square"
	viewSignInTemplateFilename    string = "views/user/sign-in"
)

func (handler *Handler) getTemplateFilename(filename string) string {
	directory, _ := os.Getwd()
	return fmt.Sprintf("%s/templates/%s.html", directory, filename)
}

func (handler *Handler) ResponseError(response http.ResponseWriter, statusCode int, message string) {
	if len(message) == 0 {
		message = http.StatusText(statusCode)
	}

	http.Error(response, message, statusCode)
}

func (handler *Handler) ResponseHTML(response http.ResponseWriter, filename string, payload interface{}) {
	layoutFilename := handler.getTemplateFilename(layoutTemplateFilename)
	currentTemplateFilename := handler.getTemplateFilename(filename)
	tpl := template.Must(template.ParseFiles(layoutFilename, currentTemplateFilename))
	if err := tpl.ExecuteTemplate(response, "layout", payload); err != nil {
		handler.ResponseError(response, http.StatusInternalServerError, "")
		return
	}
}
