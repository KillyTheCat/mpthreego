package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

/*
Function to play an mp3 file pointed to by the filepath, needs a channel to tell caller that song has finished and another channel to 
   example filepath: C:/mp3files/some_song.mp3
*/
func PlayFile(filepath string, done_main chan bool) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	
	done := make(chan bool)

	/* there's some concurrency shit going on in here
	/ Okay I understand now, this here beep.Seq function takes streamers as it's arguments and then plays them sequentially.
	/ Now instead of a streamer, we send it a callback function which signals a channel's value to true and that's basically it.
	*/
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<- done
}