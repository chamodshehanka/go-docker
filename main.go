package main

import (
	"net/http"
	"fmt"
)

func main()  {
	fmt.Println("Go Docker!!")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		
	})
}