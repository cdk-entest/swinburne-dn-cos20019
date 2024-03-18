package main

import (
	"net/http"
	"time"
)

func main ()  {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Minh Tran")
		http.ServeFile(w, r, "index.html")
	})

	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}