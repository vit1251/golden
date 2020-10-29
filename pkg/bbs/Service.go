package bbs

import (
	"github.com/gliderlabs/ssh"
	"go.uber.org/dig"
)

type Service struct {
	c *dig.Container
}

func NewService(c *dig.Container) *Service {
	s := new(Service)
	s.c = c
	return s
}

func (s *Service) handleSession(sess ssh.Session) {
	var connState = NewConnState(s.c)
	connState.SetSession(sess)
	connState.ProcessConnection()
}

func (s *Service) Start() {
	ssh.ListenAndServe("127.0.0.1:2222", s.handleSession)
}
