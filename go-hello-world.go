package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("[go-hello-world] %s serving %s %s\n", req.RemoteAddr, req.Method, req.RequestURI)
		resp.Write([]byte("Hello World"))
	})

	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8080"
	}

	log.Printf("[go-hello-world] Starting server on PORT %s\n", port)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
