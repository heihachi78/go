//https://go.dev/doc/tutorial/random-greeting

package main

import (
	"fmt"

	"tutorial/greetings"

	"log"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Hachi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(quote.Go())
}
