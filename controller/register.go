package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/MRickers/Erbhoflauf/models"
)

type RegisterHandler struct {
	Notifier models.Notifier
}

func NewHandler(n models.Notifier) http.Handler {
	return &RegisterHandler{
		Notifier: n,
	}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Read HTTP body failed: %s", err)
		return
	}

	user := models.User{}

	err = user.Deserialize(string(payload))

	if err != nil {
		log.Fatalf("Deserialize to user failed: %s", err)
		return
	}

	error := h.Notifier.Notify(user.Email, user.FormatMail())

	if error.Code != 0 {
		log.Fatalf("Notify %s failed: %s (%d)", user.Email, error.Message, error.Code)
	}
}
