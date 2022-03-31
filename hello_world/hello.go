// Declare the package as main
package main

// Functions for formatting text from standard library, including printing
import (
	"log"
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
)

func main_old() {
	//fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Printf("Hi my name is %s and my ability is %s\n", "Max", quote.Glass())
}

func main_not_quite_so_old() {
	message := greetings.HelloMyNameIs("Grewe")
	fmt.Println(message)
	message2 := greetings.HelloIamAwesome("Tim")
	fmt.Println(message2)
}

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	//Request greetings
	names := []string{"Tim", "Grewe", "Max"}
	message, err := greetings.Hellos(names)
	// If an err is returned print to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	fmt.Println(message)
}
