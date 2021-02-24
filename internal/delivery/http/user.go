package http

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (handler *Handler) userRoutes(router chi.Router) {
	router.Get("/sign-in", handler.SignIn)
	router.Post("/sign-in", handler.SignInAction)
}

func (handler *Handler) SignIn(response http.ResponseWriter, request *http.Request) {
	flash, _ := handler.cookie.Get(request, "session-name-flash")
	message := ""
	if value, ok := flash.Values["message"]; ok {
		message = value.(string)
	}

	handler.ResponseHTML(response, viewSignInTemplateFilename, struct {
		FlashMessage string
	}{
		FlashMessage: message,
	})
}

func (handler *Handler) SignInAction(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		handler.ResponseError(response, http.StatusBadRequest, err.Error())
		return
	}

	email := request.FormValue("email")
	if err := validation.Validate(&email, validation.Required, is.Email); err != nil {
		flash, _ := handler.cookie.Get(request, "session-name-flash")
		flash.Values["message"] = err.Error()
		if err := flash.Save(request, response); err != nil {
			handler.ResponseError(response, http.StatusInternalServerError, err.Error())
			return
		}
		http.Redirect(response, request, "/user/sign-in", http.StatusFound)
		return
	}

	session, _ := handler.cookie.Get(request, "session-name")
	session.Values["is_current_user_admin"] = true
	if err := session.Save(request, response); err != nil {
		handler.ResponseError(response, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(response, request, "/square", http.StatusFound)
}
