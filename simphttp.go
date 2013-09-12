// Simple static file hosting.
// Similar to python's SimpleHTTPServer

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	port = flag.String("port", "1234", "The port you want to run from localhost")
	err  error
)

func main() {
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Oh no %v\n", err)
	}

	addr := fmt.Sprintf("localhost:%s", *port)
	log.Printf("Current directory is %v\n", dir)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Running on port %s\n", addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		select {
		case s := <-c:
			log.Printf("Recieved signal %v. Gracefully shutting down...\n", s)
			os.Exit(1)
		}
	}()

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Oh no %v\n", err)
	}

}
