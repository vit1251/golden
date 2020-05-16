package audio

// +build NO_AUDIO

import (
)

type AudioManager struct {
}

func NewAudioManager() *AudioManager {
	am := new(AudioManager)
	return am
}

func (self *AudioManager) Play(name string) {
}
