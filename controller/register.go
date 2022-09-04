package controller

// Sends \p message string to \p to
// via Mail, Telegram, ...
type Notifier interface {
	Notify(to string, message string) utils.appError
}
