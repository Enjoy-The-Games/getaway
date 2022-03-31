package apperror

import "encoding/json"

type AppError struct {
	Err     error  `json:"-"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}
func (e AppError) Unwrap() error {
	return e.Err
}
func (e AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}
func NewError(err error, code int, message string) *AppError {
	return &AppError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}
