package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/MRickers/Erbhoflauf/controller"
	"github.com/MRickers/Erbhoflauf/mocks"
	"github.com/MRickers/Erbhoflauf/models"
)

func TestMailHandler(t *testing.T) {
	user := models.User{
		Name:     "Obi",
		Lastname: "Wan",
		Gender:   models.Male,
		Email:    "obiwan@sw.com",
		Year:     1990,
		Team:     "Rebellion",
		City:     "Tetuin",
		Distance: models.Short5km,
	}

	payload := url.Values{}
	payload.Add("name", "Name")
	payload.Add("lastname", "Lastname")

	r, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(payload.Encode()))

	if err != nil {
		t.Logf("Request failed: %s", err)
		return
	}

	notify := mocks.New()

	handler := controller.New(notify)

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

}
