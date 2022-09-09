package controller

import (
	"github.com/MRickers/Erbhoflauf/utils"
)

type EmailNotifier struct {
	
}
type RegisterHandler struct {
	notifier Notifier
}

// Sends \p message string to \p to
// via Mail, Telegram, ...
type Notifier interface {
	Notify(to string, message string) utils.AppError
}

func 
// Register Handler

