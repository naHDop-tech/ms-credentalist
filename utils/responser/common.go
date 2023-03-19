package responser

import (
	"time"
)

type Responser interface {
	New(data any, err error, code ResponseCode) (response Response)
}

type Response struct {
	Status    int          `json:"-"`
	Code      ResponseCode `json:"code"`
	Message   string       `json:"message"`
	Timestamp time.Time    `json:"timestamp"`
	Data      any          `json:"data"`
	Error     *DpError     `json:"error"`
}

type ResponseCode string

//example implementation
type responserExampleStruct struct {
	codeToMessageMap map[ResponseCode]string
	codeToStatusMap  map[ResponseCode]int
}

//example implementation
func NewResponserExample(codeToMessageMap map[ResponseCode]string, codeToStatusMap map[ResponseCode]int) Responser {
	return &responserExampleStruct{
		codeToMessageMap: codeToMessageMap,
		codeToStatusMap:  codeToStatusMap,
	}
}

//example implementation
func (r *responserExampleStruct) New(data any, err error, code ResponseCode) (response Response) {
	response.Status = r.codeToStatusMap[code]
	response.Code = code
	response.Error = WrapError(err)
	response.Data = data
	response.Timestamp = time.Now()
	response.Message = r.codeToMessageMap[code]
	return
}
