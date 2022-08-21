package main

import (
	// "fmt"
	"os"
	"sync"
	"time"

	tm "github.com/buger/goterm"
)
var tm_mutex sync.Mutex

func cleanup() {
	tm_mutex.Lock()
	tm.Clear()
	tm.MoveCursor(1, 7)
	tm.Println("Thanks for using mpthreego")
	tm_mutex.Unlock()
	time.Sleep(time.Second*5)
}

func main() {
	tm.Clear()
	programName := "MP3GO"
	tm.MoveCursor(1, 1)
	tm.Print(tm.Color("killythecat", tm.RED))
	tm.MoveCursor((tm.Width()/2) - (len(programName)/2), 1)
	tm.Println(programName)

	for i := 2 ; i < tm.Height(); i++ {
		tm.MoveCursor(2*(tm.Width()/3), i)
		tm.Print("|")
	}
	timeLen := 30
	go func() {
		for {
			tm_mutex.Lock()
			tm.MoveCursor(tm.Width()-timeLen, 1)
			tm.Print(time.Now().Format(time.RFC1123))
			tm.Flush()
			tm_mutex.Unlock()
		}
	}()
	tm_mutex.Lock()
	tm.MoveCursor(1,3)
	tm_mutex.Unlock()
	if len(os.Args) < 2 {
		tm_mutex.Lock()
		tm.Printf("Usage: %s [dir] where dir is a directory with mp3 files", os.Args[0])
		tm_mutex.Unlock()
		time.Sleep(time.Second*10)
		tm.Clear()
		os.Exit(1)
	}
	tm_mutex.Lock()
	tm.MoveCursor(1, 2)
	tm.Printf("Searching for mp3 files in directory %s\n", os.Args[1])
	mp3s := lookForMp3sInDirectory(os.Args[1])
	if len(mp3s) == 0 {
		tm.Printf("No MP3 files found in directory %s\n", os.Args[1])
		os.Exit(1)
	}
	tm_mutex.Unlock()
	// LOOK HERE
	done := make(chan bool)
	// playbackDone := make(chan bool)

	index := 0 //index of current song being played in queue
	numSongs := len(mp3s) // number of songs in the queue
	looping := true // if the queue loops or not
	playbackAltered := false // monitor if the index is altered in the playback coroutine. 
	var playbackCommand string // playback command

	

	for index < numSongs {

		go func() {
			tm_mutex.Lock()
			tm.Flush()
			startpos := 2
			for a := 0; a < numSongs; a++ {
				tm.MoveCursor(2*(tm.Width()/3)+2, startpos)
				
				if a == index {
					tm.Print("-> ")
				}
				x := tm.Width() - 2*(tm.Width()/3)+2
				if len(mp3s[a]) < x {
					x = len(mp3s[a]) - 1
				}
				{
					tm.Print(tm.Color(mp3s[a][:x], tm.MAGENTA))
				}
				startpos+=1
			}
			tm_mutex.Unlock()
		}()
		tm_mutex.Lock()
		tm.MoveCursor(1,6)
		// playbackControlsWindow := tm.NewBox(50|tm.PCT, tm.Height()-10, 0)
		playbackAltered = false
		tm.Println(tm.Color("Playing song ", tm.GREEN), index+1)
		
		go PlayFile(mp3s[index], done)
		tm.Println(tm.Color("Now playing file ", tm.CYAN), mp3s[index], "\nPress \"s\" to start the song over\nPress \"p\" to go to previous song\nPress \"n\" to go to next song")
		tm_mutex.Unlock()

		// tm.Print(tm.MoveTo(playbackControlsWindow.String(), 1, 5))
		go PlaybackFunction(done, &index, &playbackCommand, &playbackAltered)
		<- done

		// if playback index has been messed with in playback coroutine then don't change index
		if !playbackAltered {
			index+=1
		}
		loopControls(looping, &index, numSongs)
		tm_mutex.Lock()
		tm.Flush()
		tm_mutex.Unlock()
	}
	tm.Clear()
	tm.Println("End of queue. Exitting...\n Thanks for using playmp3go")
}