package helper

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/mochammadshenna/aplikasi-po/internal/model/api"
	"github.com/mochammadshenna/aplikasi-po/internal/util/exceptioncode"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := api.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		WriteToResponseBody(writer, webResponse)
		return true // konversi ke error handler
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(exceptioncode.NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := api.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, webResponse)
		return true // konversi ke error handler
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := api.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server error",
		Data:   err,
	}

	WriteToResponseBody(writer, webResponse)

}
