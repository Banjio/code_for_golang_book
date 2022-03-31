package greetings

import (
	"errors"
	"fmt"
	"rsc.io/quote"
	"math/rand"
	"time"

)


func HelloMyNameIs(name string) string {
	//Sprint f works similiar to f strings in python.
	message := fmt.Sprintf("Hi, my name is %v. And i sSuuuuuuuuuuck!", name)
	return message
}

func HelloIamAwesome(name string) string {
	//Sprint f works similiar to f strings in python.
	message := fmt.Sprintf("Hi my name is %s and my ability is %s", name, quote.Glass())
	return message
}

func Hello(name string) (string, error) {
	// Without a name given return an error message
	if name == "" {
		return "", errors.New("Oy mate ya need a name ya bloody wanker.")

	}

	// With a name given return the usual message
	// This breaks the hello function to see a failing test
	//message := fmt.Sprint(randomFormat())

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func Hellos(names []string) (map[string]string, error){
	// A map to associate names with messages
	messages := make(map[string]string)
	//Loop through received slices of names, calling the hello function each time to get a message
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats. 
	formats := []string{
		"Hi, %v. Welcome!",
		"Oy, Get on that train, %v!",
		"HAIL HYDRA, %v!",
	}

	// Return a randomly selected message by specifying a random slice
	return formats[rand.Intn(len(formats))]
}

