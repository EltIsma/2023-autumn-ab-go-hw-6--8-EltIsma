package models

import "errors"

// type ResponseError struct {
// 	Err     error
// }

// func (e *ResponseError) Error() string {
// 	return e.Err.Error()
// }

var ErrNotFound = errors.New("not found")

var ErrAlredyExist = errors.New("already exist")