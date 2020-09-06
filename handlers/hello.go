package handlers

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	
)

//Hello is the struct 
type Hello struct {
	l *log.Logger
}

//NewHello is a new instance of Hello struct
func NewHello (l *log.Logger) *Hello  {
	return &Hello{l}
}

//ServerHTTP is a new instance of the http responsea
func (h *Hello)  ServeHTTP(rw http.ResponseWriter, r *http.Request){
	log.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("There is an error")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Ooops"))
		return 
	}
	fmt.Fprintf(rw, "Hello %s" , d)
}
