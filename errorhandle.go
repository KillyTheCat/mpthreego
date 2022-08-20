package main

import (
	"log"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func DeferPushChan(ch chan bool) {
	ch <- true
}