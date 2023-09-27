package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nama kosong")
	}

	message := fmt.Sprintf("Hai, %v. Apa kabar?", name)

	return message, nil
}
