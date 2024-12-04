//https://go.dev/doc/tutorial/random-greeting

package main

import (
	"fmt"
	"log"
	"tutorial/greetings"

	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
	fmt.Println(reverse.String(quote.Go()))
	fmt.Println(reverse.Int(24601))
}
