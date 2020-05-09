package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type pairWeb struct {
	x, y int
}

func (p pairWeb) String() string { // p is called the "receiver"
	// Sprintf is another public function in package fmt.
	// Dot syntax references fields of p.
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// A single function from package http starts a web server.
func learnWebProgramming() {
	// First parameter of ListenAndServe is TCP address to listen to.
	// Second parameter is an interface, specifically http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pairWeb{})
		fmt.Println(err) // don't ignore errors
	}()
	requestServer()
}

// Make pair an http.Handler by implementing its only method, ServeHTTP.
func (p pairWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve data with a method of http.ResponseWriter.
	w.Write([]byte("You learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}

func main() {
	learnWebProgramming()
}
