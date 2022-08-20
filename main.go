package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [dir] where dir is a directory with mp3 files", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("Searching for mp3 files in directory %s\n", os.Args[1])
	mp3s := lookForMp3sInDirectory(os.Args[1])
	if len(mp3s) == 0 {
		fmt.Printf("No MP3 files found in directory %s\n", os.Args[1])
		os.Exit(1)
	}
	// LOOK HERE
	done := make(chan bool)
	// playbackDone := make(chan bool)

	index := 0 //index of current song being played in queue
	numSongs := len(mp3s) // number of songs in the queue
	looping := true // if the queue loops or not
	playbackAltered := false // monitor if the index is altered in the playback coroutine. 
	var playbackCommand string // playback command

	for index < numSongs {
		playbackAltered = false
		fmt.Println("Playing song ", index+1)
		
		go PlayFile(mp3s[index], done)
		fmt.Println("Now playing file ", mp3s[index], "\nPress \"s\" to start the song over\nPress \"p\" to go to previous song\nPress \"n\" to go to next song")
		go PlaybackFunction(done, &index, &playbackCommand, &playbackAltered)
		<- done

		// if playback index has been messed with in playback coroutine then don't change index
		if !playbackAltered {
			index+=1
		}
		loopControls(looping, &index, numSongs)
	}
	fmt.Println("End of queue. Exitting...\n Thanks for using playmp3go")
}