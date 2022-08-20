package main

import "fmt"

func PlaybackFunction(sig chan bool, index *int, playbackCommand *string, playbackAltered *bool) {
	fmt.Scanln(playbackCommand)
	*playbackAltered = true
	switch *playbackCommand {
	case "s":
		sig <- true
		return
	case "n":
		*index++
		sig <- true
		return
	case "p":
		*index--
		sig <- true
		return
	}
}

func loopControls(looping bool, index *int, numSongs int) {
	if *index == numSongs && looping {
		*index = 0
	} else if *index < 0 && looping {
		*index = numSongs - 1
	}
}