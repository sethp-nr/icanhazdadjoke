package main

import (
	"fmt"

	"github.com/sethp-nr/icanhazdadjoke"
	"net/http"
)

func main() {
	joke, err := icanhazdadjoke.GetRandomJokeText()
	if err != nil {
		fmt.Println("Unable to retrieve joke:", err)
	}
	fmt.Println(joke)
}
