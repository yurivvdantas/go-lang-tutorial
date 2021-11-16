package greetings

import (
	"errors"
	"math/rand"
	"time"

	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Check if name is not null
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
