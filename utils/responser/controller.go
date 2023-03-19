package responser

import (
	"net/http"
	"time"
)

type responser struct {
}

func NewResponser() Responser {
	return &responser{}
}

const (
	OK           ResponseCode = "code.100001"
	BAD_REQUEST  ResponseCode = "code.100002"
	FAIL         ResponseCode = "code.100003"
	UNAUTHORIZED ResponseCode = "code.100004"
	NOT_FOUND    ResponseCode = "code.100005"
)

var codeToMessageMap = map[ResponseCode]string{
	OK:           "ok",
	BAD_REQUEST:  "bad request",
	FAIL:         "failed",
	UNAUTHORIZED: "Unauthorized",
	NOT_FOUND:    "Resource not found",
}

var codeToStatusMap = map[ResponseCode]int{
	OK:           http.StatusOK,
	BAD_REQUEST:  http.StatusBadRequest,
	FAIL:         http.StatusInternalServerError,
	UNAUTHORIZED: http.StatusUnauthorized,
	NOT_FOUND:    http.StatusNotFound,
}

func (r *responser) New(data any, err error, code ResponseCode) (response Response) {
	response.Status = codeToStatusMap[code]
	response.Code = code

	response.Data = data
	response.Timestamp = time.Now()
	response.Message = codeToMessageMap[code]
	response.Error = WrapError(err)
	return
}
