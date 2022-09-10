package models

import (
	"encoding/json"
	"fmt"

	"github.com/MRickers/Erbhoflauf/utils"
)

type TrackLength uint16
type Gender string

const (
	Short900m TrackLength = iota
	Short5km
	Long10km
)

const (
	Male    Gender = "male"
	Female         = "female"
	Diverse        = "diverse"
)

// User to register
type User struct {
	Name     string      `json:"name"`
	Lastname string      `json:"lastname"`
	Gender   Gender      `json:"gender"`
	Email    string      `json:"email"`
	Year     uint16      `json:"year"`
	Team     string      `json:"team"`
	City     string      `json:"city"`
	Distance TrackLength `json:"distance"`
}

func (u *User) Serialize() (string, error) {
	stream, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(stream), nil
}

func (u *User) Deserialize(stream string) error {
	err := json.Unmarshal([]byte(stream), u)
	return err
}

type EmailNotifier struct {
	Host string
	Pw   string
}

func (e *EmailNotifier) Notify(to string, message string) utils.AppError {
	return utils.AppError{
		Error:   fmt.Errorf("Error"),
		Message: "Error message",
		Code:    1,
	}
}

// interface
// Sends message to to via mail, whatsapp, telegram, etc..
type Notifier interface {
	Notify(to string, message string) utils.AppError
}

// Convert string stream to user and back
type Converter interface {
	Deserialize(stream string) (User, error)
	Serialize(User) (string, error)
}
