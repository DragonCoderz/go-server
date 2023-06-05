package main

import(
	"fmt"
	"log"
	"net/http"
)

//This function is the handler for the "/form" URL. It takes two parameters: w http.ResponseWriter, which is used to send the response, and r *http.Request, which represents the HTTP request received.

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { //called to parse the form data sent in the request. If there is an error, it is handled by writing an error message to the response. The name and address values are obtained from the parsed form data using r.FormValue.
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


//This function is the handler for the "/hello" URL. It checks if the request URL path is "/hello" and if the request method is "GET". If either condition is not satisfied, it returns a 404 error response using http.Error.

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main(){ //entry point of the program. Sets up the server configuration and starts the server.
	fileServer := http.FileServer(http.Dir("./static")) //creates a file server handler that serves files from the "./static" directory.
	http.Handle("/", fileServer) //registers the file server handler to handle requests for the root URL ("/").
	http.HandleFunc("/form", formHandler) //registers the formHandler function to handle requests for the "/form" URL.
	http.HandleFunc("/hello", helloHandler) //registers the helloHandler function to handle requests

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}