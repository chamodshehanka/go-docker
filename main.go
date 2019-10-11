package main

import (
	"log"
	"net/http"
	"fmt"
)

func main()  {
	fmt.Println("Go Docker!!")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))	
}