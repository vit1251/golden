package audio

// +build WINDOWS

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	path2 "path"
	"time"
)

type AudioManager struct {
	basePath string
}

func NewAudioManager() *AudioManager {
	am := new(AudioManager)
//	am.basePath = "C:\\WORK\\golden\\static\\audio\\modern"
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	am.basePath = path2.Join(dir, "static", "audio")
	return am
}

func (self *AudioManager) Play(name string) {

	path := path2.Join(self.basePath, name)

	f, err := os.Open(path)
	if err != nil {
		log.Printf("error: Unable to play %+v", path)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		done <- true
	})))

	<-done
}
