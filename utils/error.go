package utils

import (
	"net/http"
)

type AppError struct {
	Error   error
	Message string
	Code    int
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError
