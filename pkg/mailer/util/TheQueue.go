package util

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

type TheQueue struct {
	frames []stream.Frame
}

func NewTheQueue() *TheQueue {
	queue := new(TheQueue)
	return queue
}

func (self TheQueue) IsEmpty() bool {
	return len(self.frames) == 0
}

func (self *TheQueue) Push(nextFrame stream.Frame) {
	self.frames = append(self.frames, nextFrame)
}

func (self *TheQueue) Pop() *stream.Frame {
	if len(self.frames) > 0 {
		result := self.frames[0]
		self.frames = self.frames[1:]
		return &result
	} else {
		return nil
	}
}

func (self TheQueue) Dump() {
	log.Printf("Queue: items = %+v", self.frames)
}
