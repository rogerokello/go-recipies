package main

import "net/http"

func main() {
	/*
		- Parameters of HandleFunc are
		1. The path to listen for a request
		2. A handler function that will be called when a request comes on that path

		The Handler function takes two Parameters
		1. A ResponseWriter - We shall use this to write responses back to the requester
		2. A Pointer to the request object

		The handler function wraps all the information that is got from the incoming request
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// This method is used to write back the response to requester.
		// The response sent must be a slice of bytes
		// []byte("String")-This converts a string to a slice of bytes
		w.Write([]byte("Hello Go!"))
	})

	// This is how a server is set up
	// When you pass nil as the second parameter, it is going to use the default
	// mux server. It just means it is going to give us a default instance
	// of a web server.
	http.ListenAndServe(":8000", nil)
}
