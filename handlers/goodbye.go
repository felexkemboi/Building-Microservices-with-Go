package handlers

import (
	"log"
	"net/http"
)

//Goodbye to just avoid errors
type Goodbye struct {
	l *log.Logger
}

//NewGoodbye is  a comment just to avoid errors
func NewGoodbye( l *log.Logger) *Goodbye{
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Byee"))
}


