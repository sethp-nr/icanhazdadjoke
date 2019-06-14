package main

import (
	"fmt"

	"github.com/sethp-nr/icanhazdadjoke"
)

func main() {
	joke, err := icanhazdadjoke.GetRandomJokeText()
	if err != nil {
		fmt.Println("Unable to retrieve joke:", err)
	}
	fmt.Println(joke)
}
