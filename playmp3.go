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
Function to play an mp3 file pointed to by the filepath and filename
   example filepath: C:/mp3files/
   example filename: some_song.mp3
*/
func PlayFile(filepath string, filename string, done_main chan bool) {
	f, err := os.Open(filepath+filename)
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

	// there's some concurrency shit going on in here
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<- done
	done_main <- true
}