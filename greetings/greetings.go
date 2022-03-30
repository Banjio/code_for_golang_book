package greetings

import (
	"errors"
	"fmt"
	"rsc.io/quote"
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
	return HelloIamAwesome(name), nil
}