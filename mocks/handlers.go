package mocks

import (
	"fmt"

	"github.com/MRickers/Erbhoflauf/utils"
)

type MockNotifier struct {
	To      string
	Message string
}

func (m *MockNotifier) Notify(to string, message string) utils.AppError {

	m.To = to
	m.Message = message

	return utils.AppError{
		Error:   fmt.Errorf("OK"),
		Message: "",
		Code:    0,
	}
}

func NewMockNotifier() *MockNotifier {
	return &MockNotifier{
		To:      "",
		Message: "",
	}
}
