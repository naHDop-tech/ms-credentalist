package responser

//default codes for all microservices
const (
	UNDEFINED_CODE     ErrorCode = "undefined"
	BAD_REQUEST_CODE   ErrorCode = "bad_request"
	UNAUTHORIZED_CODE  ErrorCode = "unauthorized"
	ACCESS_DENIED_CODE ErrorCode = "access_denied"
)

var errorCodeToMessageMap = map[ErrorCode]string{
	UNDEFINED_CODE:     "Unknown error",
	BAD_REQUEST_CODE:   "Bad request",
	UNAUTHORIZED_CODE:  "Unauthorized",
	ACCESS_DENIED_CODE: "Access denied",
}

type DpError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Type    ErrorType `json:"-"`
}

func (m *DpError) Error() string {
	return m.Message
}

type ErrorType int

type ErrorCode string

const (
	ET_INTERNAL ErrorType = iota
	ET_BUSINESS
)

func NewDpUndefinedError(err error) *DpError {
	if err == nil {
		return nil
	}
	return &DpError{
		Message: err.Error(),
		Type:    ET_INTERNAL,
		Code:    UNDEFINED_CODE,
	}
}

func WrapError(err error) *DpError {
	if e, ok := err.(*DpError); ok {
		return e
	}

	return NewDpUndefinedError(err)
}
