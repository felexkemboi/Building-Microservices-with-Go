package main

import ("net/http"
	"context"
		 "time"
		 "os/signal"
		 "log"
		"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
		"os"
		)
func main (){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr : ":6060",                     //configure the bind address
		Handler: sm,                        //set the default handler
		ErrorLog : l,						// set the logger for the server
		IdleTimeout:120*time.Second,        // 
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	} ()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan,os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate", sig)


	s.ListenAndServe()

	tc,cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	//tc, _ :=  context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}

