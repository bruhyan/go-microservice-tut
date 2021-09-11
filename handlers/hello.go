package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

// Dependency injection
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.logger.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// Using http to return error
		// Works the same as the commented block below
		http.Error(rw, "Oops Error Encountered", http.StatusBadRequest)
		return

		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops Error Encountered"))
		// return
	}

	// Use ResponseWriter to write data back to request
	fmt.Fprintf(rw, "Hello %s", d)
}
