package main

// Imports 
import (
	"fmt"
	"net/http"
	".."
)

//Function to server eliza 
func resquestHandler(w http.ResponseWriter, r *http.Request) {
	
	//takes in user input from the webpage
	input := r.URL.Query().Get("input")

	// gets output 
	output := eliza.Reply(input)

	// Outputs result to the webpage
	fmt.Fprintf(w, output)
}

//function main to server webpage.
func main() {
	// servers webpage server
	dir := http.Dir("../webpage")
	fileServer := http.FileServer(dir)

	// HTTP request handler
	http.Handle("/", fileServer)
	http.HandleFunc("/elizaPrompt", resquestHandler)
	http.ListenAndServe(":8080", nil)
}