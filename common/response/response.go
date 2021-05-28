package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrForbiddenResource   = errors.New("Forbidden resource")
	ErrNotFound            = errors.New("Not Found")
	ErrTimeoutError        = errors.New("Timeout error")
	ErrAlreadyRegistered   = errors.New("User already registered")
	ErrInternalServerError = errors.New("Internal server error")
	ErrNoValidUserFound    = errors.New("No Valid User Found")
)

const (
	STATUSCODE_GENERICSUCCESS = "200"
	STATUSCODE_BADREQUEST     = "400"
	STATUS_FORBIDDEN          = "403"
	STATUSCODE_NOT_FOUND      = "404"
	STATUSCODE_INTERNAL_ERROR = "500"
	STATUSCODE_TIMEOUT_ERROR  = "504"
)

func GetErrorCodeStr(err error) string {
	switch err {
	case ErrBadRequest:
		return STATUSCODE_BADREQUEST
	case ErrForbiddenResource:
		return STATUS_FORBIDDEN
	case ErrNotFound:
		return STATUSCODE_NOT_FOUND
	case ErrInternalServerError:
		return STATUSCODE_INTERNAL_ERROR
	case ErrTimeoutError:
		return STATUSCODE_TIMEOUT_ERROR
	case ErrAlreadyRegistered:
		return STATUSCODE_BADREQUEST
	case nil:
		return STATUSCODE_GENERICSUCCESS
	case ErrNoValidUserFound:
		return STATUSCODE_BADREQUEST
	default:
		return STATUSCODE_INTERNAL_ERROR
	}
}

func GetHTTPCode(code string) int {
	s := code[0:3]
	i, _ := strconv.Atoi(s)
	return i
}

type JSONResponse struct {
	Code        string      `json:"code"`
	Message     string      `json:"message,omitempty"`
	ErrorString string      `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	StatusCode  int         `json:"-"`
	Error       error       `json:"-"`
	Success     *int        `json:"success,omitempty"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{Code: STATUSCODE_GENERICSUCCESS, StatusCode: GetHTTPCode(STATUSCODE_GENERICSUCCESS)}
}

func (r *JSONResponse) SetSuccess(code int) *JSONResponse {
	r.Success = &code
	return r
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func getErrType(err error) error {
	return err
}

func (r *JSONResponse) SetError(err error, a ...string) *JSONResponse {
	err = getErrType(err)
	r.Error = err
	r.ErrorString = err.Error()
	r.Code = GetErrorCodeStr(err)
	r.StatusCode = GetHTTPCode(r.Code)
	if r.StatusCode != http.StatusForbidden {
		log.Println(err)
	}
	if r.StatusCode == http.StatusInternalServerError {
		r.ErrorString = "Internal Server error"
	}
	if len(a) > 0 {
		r.ErrorString = a[0]
	}
	return r
}

func (r *JSONResponse) Send(w http.ResponseWriter) {
	b, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	w.Write(b)
}
