package errors

import "errors"

var ErrPostAlredyExist = errors.New("Error Post alredy Exits")
var ErrPostNotFound = errors.New("Error Post not found")
var ErrBadRequest = errors.New("Error Bad Request")
var ErrFromEncodeJson = errors.New("Error From Encode Json")
var ErrInternalError = errors.New("Internal Errors")
