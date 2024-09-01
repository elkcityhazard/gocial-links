package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	msgChan := make(chan string)
	errorChan := make(chan error)
	go initServer(msgChan, errorChan)
	msgObserver(msgChan, errorChan)
}

func initServer(msgChan chan string, errorChan chan error) {

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           routes(),
		IdleTimeout:       time.Second * 30,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 10,
		ReadHeaderTimeout: time.Second * 10,
	}

	msgChan <- fmt.Sprintf("Starting server on %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		errorChan <- err
		log.Fatalln(err)
	}

}

func msgObserver(msgChan chan string, errorChan chan error) {
	for {
		select {
		case msg := <-msgChan:
			fmt.Println(msg)
		case err := <-errorChan:
			fmt.Println(err.Error())
		}
	}
}
