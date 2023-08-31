package httphelper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/mochammadshenna/aplikasi-po/model/api"
	"github.com/mochammadshenna/aplikasi-po/state"
	"github.com/mochammadshenna/aplikasi-po/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/logger"
)

var Decoder = schema.NewDecoder()

func init() {
	Decoder.RegisterConverter([]string{}, convertStringCommaSeparated)
}

func convertStringCommaSeparated(value string) reflect.Value {
	return reflect.ValueOf(strings.Split(value, ","))
}

func Read(request *http.Request, result interface{}) error {
	err := Decoder.Decode(result, request.URL.Query())

	if err != nil {
		return parseError(err)
	}

	if request.Method == http.MethodPost || request.Method == http.MethodPut || request.Method == http.MethodPatch {
		jsonDecoder := json.NewDecoder(request.Body)
		err = jsonDecoder.Decode(result)
		if err != nil && err != io.EOF {
			logger.Error(request.Context(), err)
			return api.ErrorResponse{
				Code:    exceptioncode.CodeInvalidRequest,
				Message: err.Error(),
			}
		}
	}

	logger.Info(request.Context(), strings.Replace(fmt.Sprintf("request: %+v", result), "\u0026", "", 1))
	return nil
}

func ReadJsonMultipart(request *http.Request, result interface{}) error {
	err := Decoder.Decode(result, request.URL.Query())
	if err != nil {
		return parseError(err)
	}

	if request.Method == http.MethodPost || request.Method == http.MethodPut || request.Method == http.MethodPatch {
		err := json.Unmarshal([]byte(request.FormValue("json")), result)
		if err != nil && err != io.EOF {
			logger.Error(request.Context(), err)
			return api.ErrorResponse{
				Code:    exceptioncode.CodeInvalidRequest,
				Message: err.Error(),
			}
		}
	}

	logger.Info(request.Context(), strings.Replace(fmt.Sprintf("request: %+v", result), "\u0026", "", 1))
	return nil
}

func Write(ctx context.Context, writer http.ResponseWriter, data interface{}) {
	response := api.ApiResponse{
		Header: getHeader(writer),
		Data:   data,
	}
	write(ctx, writer, response)
}

func WriteError(ctx context.Context, writer http.ResponseWriter, errorResponse error) {
	writer.WriteHeader(http.StatusBadRequest)
	response := api.ApiResponse{
		Header: getHeader(writer),
		Error:  errorResponse,
	}
	write(ctx, writer, response)
}

func WriteErrorWithData(ctx context.Context, writer http.ResponseWriter, errorResponse error, data interface{}) {
	writer.WriteHeader(http.StatusBadRequest)
	response := api.ApiResponse{
		Header: getHeader(writer),
		Data:   data,
		Error:  errorResponse,
	}
	write(ctx, writer, response)
}

func DownloadFile(filePath string, url string, overwrite bool) (err error) {
	filePath = "./files" + filePath
	_, err = os.Stat(filePath)
	if err == nil && !overwrite {
		return nil
	}

	dir := filepath.Dir(filePath)
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func write(ctx context.Context, writer http.ResponseWriter, response api.ApiResponse) {
	// if response.Error != nil {
	// 	logger.Errorf(ctx, "response: %+v", response)
	// } else {
	// 	logger.Infof(ctx, "response: %+v", response)
	// }

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	helper.PanicError(err)
}

func getHeader(writer http.ResponseWriter) api.HeaderResponse {
	headerResponse := api.HeaderResponse{
		ServerTimeMs: time.Now().Unix(),
		RequestId:    writer.Header().Get(string(state.HttpHeaders().RequestId)),
	}

	startTimeHeader := writer.Header().Get(string(state.HttpHeaders().StartTime))
	if len(startTimeHeader) > 0 {
		startTime, _ := strconv.ParseInt(startTimeHeader, 10, 64)
		headerResponse.ProcessTimeMs = time.Since(time.Unix(0, startTime)).Milliseconds()
	}

	return headerResponse
}

func parseError(err error) error {
	errors := []api.ErrorValidate{}
	new := err.(schema.MultiError)
	for i, a := range new {
		errors = append(errors, api.ErrorValidate{
			Key:     i,
			Code:    "VALIDATION",
			Message: a.Error(),
		})
	}
	return api.ErrorResponse{
		Code:    exceptioncode.CodeInvalidValidation,
		Message: "validation error",
		Errors:  errors,
	}
}
