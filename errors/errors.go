package errors

import "net/http"

// HandleError TO DO
type HandleError struct {
	Message    string
	Name       string
	Params     interface{}
	HTTPStatus int
}

// CreateNoContent TO DO
func CreateNoContent() *HandleError {
	return &HandleError{
		Message:    "No data found",
		Name:       "no_data_found",
		HTTPStatus: http.StatusNoContent, //204
	}
}

// CreateErrorMultiRows TO DO
func CreateErrorMultiRows() *HandleError {
	return &HandleError{
		Message:    "Multiple rows found when expecting single row",
		Name:       "multiple_rows",
		HTTPStatus: http.StatusBadRequest, //400
	}
}

// CreateErrorNoAuthorization TO DO
func CreateErrorNoAuthorization() *HandleError {
	return &HandleError{
		Message:    "Missing Bearer token in Authorization request header",
		Name:       "no_authorization",
		HTTPStatus: http.StatusUnauthorized, //401
	}
}

// CreateErrorNoRows TO DO
func CreateErrorNoRows() *HandleError {
	return &HandleError{
		Message:    "No row found with specified id",
		Name:       "not_found",
		HTTPStatus: http.StatusBadRequest, //400
	}
}

// CreateValidationError TO DO
func CreateValidationError() *HandleError {
	return &HandleError{
		Message:    "The request is not valid",
		Name:       "validation_error",
		HTTPStatus: http.StatusBadRequest, //400
	}
}

// CreateServerError TO DO
func CreateServerError() *HandleError {
	return &HandleError{
		Message:    "An error has occurred",
		Name:       "server_error",
		HTTPStatus: http.StatusInternalServerError, //500
	}
}

// CreateUnknownField TO DO
func CreateUnknownFieldError() *HandleError {
	return &HandleError{
		Message:    "Unknown field for idea",
		Name:       "request_error",
		HTTPStatus: http.StatusNotFound, //404
	}
}
