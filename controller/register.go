package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MRickers/Erbhoflauf/models"
)

type RegisterHandler struct {
	Notify models.Notifier
}

func New(n models.Notifier) http.Handler {
	return &RegisterHandler{
		Notify: n,
	}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("ServeHTTP :D")

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user models.User
	// Convert to User type
	err = json.Unmarshal(payload, &user)
	if err != nil {
		panic(err)
	}

}
