package main

import ("net/http"
		 "log"
		 "fmt"
		"io/ioutil"
		)
func main (){

	http.HandleFunc("/", func (rw http.ResponseWriter, r *http.Request){
		//log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println("There is an error")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops"))
			return 
		}
		fmt.Fprintf(rw, "Hello %s" , d)



		//log.Printf("Data is %s\n", d)

	})

	http.HandleFunc("/goodbye", func (http.ResponseWriter, *http.Request){
		log.Println("Good Bye")
	})

	http.ListenAndServe(":9090",nil)
}

