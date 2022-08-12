package errorx

const (
	DefaultCode  = 1001
	UserNotFound = 1002
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(DefaultCode, msg)
}

func NewUserNotFoundError(msg string) error {
	return NewCodeError(UserNotFound, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) ErrCode() int {
	return e.Code
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
