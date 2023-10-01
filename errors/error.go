package errors

const (
	ErrorMessageKey           = "error"
	ErrorCodeInvalidParameter = "INVALID-PARAMETER"
	ErrorCodeGeneral          = "GENERAL-ERROR"
	ErrorCodeNotFound         = "NOT-FOUND"
	ErrorCodeUnauthorized     = "UNAUTHORIZED"
	ErrorCodeInternal         = "INTERNAL-ERROR"
)

type Err struct {
	errorCode      string
	message        string
	httpStatusCode int
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) GetHttpStatusCode() int {
	return e.httpStatusCode
}

func (e *Err) GetErrorCode() string {
	return e.errorCode
}

func NewErr(httpStatusCode int, errorCode, message string) *Err {
	return &Err{
		httpStatusCode: httpStatusCode,
		errorCode:      errorCode,
		message:        message,
	}
}
