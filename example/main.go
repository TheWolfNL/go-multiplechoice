package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thewolfnl/go-multiplechoice"
)

func main() {
	var question = "What are Go Programmers called?"
	var options = [...]string{"Go Programmers", "Gogrammers", "Gophers", "Go Nerds"}

	// Single selection
	selection := MultipleChoice.Selection(question, options[:])
	fmt.Printf("Final selection: %s, this can be used in your program.\n", selection)

	// Multi selection
	selections := MultipleChoice.MultiSelection(question, options[:])
	json, _ := json.Marshal(selections)
	log.Printf("Result: %+v\n", string(json))
}
