package utils

import (
    "encoding/json"
)

func Deserialize[T any] (stream string) (T, error) {
	var t T
	err := json.Unmarshal([]byte(stream), &t)
	return t, err
}

func Serialize[T any] (t T) (string, error) {
	stream, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(stream), nil
}