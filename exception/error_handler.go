package exception

import (
	"caturilham05/product/helper"
	"caturilham05/product/model/web"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if conflictError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		var validationErrors validator.ValidationErrors
		var textValidationErrors string

		if errors.As(exception, &validationErrors) {
			var errorMessages []string
			for _, e := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf("%s %s", e.Field(), e.Tag()))
			}
			textValidationErrors = strings.Join(errorMessages, ",")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   textValidationErrors,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   nil,
	}

	fmt.Println("Internal Server Error:", err)

	helper.WriteToResponseBody(w, webResponse)
}

func conflictError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(ConflictError)
	if !ok {
		return false
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusConflict)

	webResponse := web.WebResponse{
		Code:   http.StatusConflict,
		Status: "CONFLICT",
		Data:   exception.Error,
	}

	helper.WriteToResponseBody(w, webResponse)
	return true
}

func badRequestError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(BadRequestError)
	if !ok {
		return false
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data:   exception.Error,
	}

	helper.WriteToResponseBody(w, webResponse)
	return true
}
