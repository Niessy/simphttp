// Simple static file hosting.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	port string
	dir  string
	err  error
)

func init() {
	flag.StringVar(&port, "port", ":8000", "The port to listen on, 8000 by default")
	flag.StringVar(&port, "p", ":8000", "Shortform for port")
	flag.StringVar(&dir, "dir", ".", "Directory to serve files from, current directory by default")
	flag.StringVar(&dir, "d", ".", "Shortform for dir")
}

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Listening on 127.0.0.1%s\n", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		select {
		case s := <-c:
			log.Printf("Recieved signal %v. Gracefully shutting down...\n", s)
			os.Exit(1)
		}
	}()

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Oh no %v\n", err)
	}

}
