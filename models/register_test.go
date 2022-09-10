package models_test

import (
	"testing"

	"github.com/MRickers/Erbhoflauf/models"
)

func TestUserModelDeserialze(t *testing.T) {
	userJson := `{"name": "ObiWan", "lastname":"Kenobi", "gender":"male", "email":"obiwan@sw.com", "year":2000,"team":"Rebellion", "city":"Tetuin", "distance":0}`

	var user models.User

	user.Deserialize(userJson)

	if user.Name != "ObiWan" {
		t.Logf("Invalid name: %s != %s", user.Name, "ObiWan")
	}
	if user.Lastname != "Kenobi" {
		t.Logf("Invalid name: %s != %s", user.Name, "Kenobi")
	}
	if user.Gender != "male" {
		t.Logf("Invalid name: %s != %s", user.Name, "male")
	}
	if user.Email != "obiwan@sw.com" {
		t.Logf("Invalid name: %s != %s", user.Name, "obiwan@sw.com")
	}
	if user.Year != 2000 {
		t.Logf("Invalid name: %d != %d", user.Year, 2000)
	}
	if user.Team != "Rebellion" {
		t.Logf("Invalid name: %s != %s", user.Name, "Rebellion")
	}
	if user.City != "Tetuin" {
		t.Logf("Invalid name: %s != %s", user.Name, "Tetuin")
	}
	if user.Distance != 0 {
		t.Logf("Invalid name: %d != %d", user.Distance, 0)
	}
}
