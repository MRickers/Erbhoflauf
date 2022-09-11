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
	"github.com/MRickers/Erbhoflauf/utils"
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

	userJson, err := utils.Serialize[models.User](user)

	if err != nil {
		t.Fatalf("Could not serialze user: %s", err)
	}

	payload := url.Values{}
	payload.Add("user", userJson)

	r, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(userJson))

	if err != nil {
		t.Logf("Request failed: %s", err)
		return
	}

	notify := mocks.NewMockNotifier()

	handler := controller.NewHandler(notify)

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if notify.To != "obiwan@sw.com" {
		t.Fatalf("Wrong email: %s != obiwan@sw.com", notify.To)
	}
	if notify.Message != user.FormatMail() {
		t.Fatalf("Wrong message: %s != %s", notify.Message, user.FormatMail())
	}
}
