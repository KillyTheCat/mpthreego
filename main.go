package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Searching for mp3 files in directory %s\n", os.Args[1])
	mp3s := lookForMp3sInDirectory(os.Args[1])
	if len(mp3s) == 0 {
		fmt.Printf("No MP3 files found in directory %s\n", os.Args[1])
		os.Exit(1)
	}
	for index, filepath := range mp3s {
		fmt.Println("Playing song ", index)
		done := make(chan bool)
		exit := make(chan bool)
		go PlayFile(filepath, done, exit)
		fmt.Println("Now playing file ", filepath, "\nPress \"s\" and then enter to skip current file")
		var temp string
		go func() {
			for {
				fmt.Scanln(&temp)
				switch temp {
				case "s":
					done <- true
					return
				}
			}
		}()
		<- done
	}
	fmt.Println("End of queue. Exitting...\n Thanks for using playmp3go")
}