package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Trying to play file %s from folder %s\n", os.Args[2], os.Args[1])
	done := make(chan bool)
	go PlayFile(os.Args[1], os.Args[2], done)
	<- done
}