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
		t.Fatalf("Invalid name: %s != %s", user.Name, "ObiWan")
	}
	if user.Lastname != "Kenobi" {
		t.Fatalf("Invalid Lastname: %s != %s", user.Lastname, "Kenobi")
	}
	if user.Gender != "male" {
		t.Fatalf("Invalid Gender: %s != %s", user.Gender, "male")
	}
	if user.Email != "obiwan@sw.com" {
		t.Fatalf("Invalid Email: %s != %s", user.Email, "obiwan@sw.com")
	}
	if user.Year != 2000 {
		t.Fatalf("Invalid Year: %d != %d", user.Year, 2000)
	}
	if user.Team != "Rebellion" {
		t.Fatalf("Invalid Team: %s != %s", user.Team, "Rebellion")
	}
	if user.City != "Tetuin" {
		t.Fatalf("Invalid City: %s != %s", user.City, "Tetuin")
	}
	if user.Distance != 0 {
		t.Fatalf("Invalid Distance: %d != %d", user.Distance, 0)
	}
}

func TestUserModelDeserialzeFailed(t *testing.T) {
	userJson := `{"name": "ObiWan", "gender":"male", "email":"obiwan@sw.com", "year":2000,"team":"Rebellion", "city":"Tetuin", "distance":0}`

	var user models.User

	user.Deserialize(userJson)

	if user.Name != "ObiWan" {
		t.Fatalf("Invalid name: %s != %s", user.Name, "ObiWan")
	}
	if user.Gender != "male" {
		t.Fatalf("Invalid Gender: %s != %s", user.Gender, "male")
	}
	if user.Email != "obiwan@sw.com" {
		t.Fatalf("Invalid Email: %s != %s", user.Email, "obiwan@sw.com")
	}
	if user.Year != 2000 {
		t.Fatalf("Invalid Year: %d != %d", user.Year, 2000)
	}
	if user.Team != "Rebellion" {
		t.Fatalf("Invalid Team: %s != %s", user.Team, "Rebellion")
	}
	if user.City != "Tetuin" {
		t.Fatalf("Invalid City: %s != %s", user.City, "Tetuin")
	}
	if user.Distance != 0 {
		t.Fatalf("Invalid Distance: %d != %d", user.Distance, 0)
	}
	if user.Lastname != "" {
		t.Fatalf("Invalid Lastname not empty")
	}
}
