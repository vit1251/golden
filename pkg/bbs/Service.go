package bbs

import (
	"github.com/gliderlabs/ssh"
	"github.com/vit1251/golden/pkg/registry"
)

type BoardService struct {
	registry *registry.Container
}

func NewBoardService(r *registry.Container) *BoardService {
	s := new(BoardService)
	s.registry = r
	return s
}

func (s *BoardService) handleSession(sess ssh.Session) {
	var connState = NewConnState(s.registry)
	connState.SetSession(sess)
	connState.ProcessConnection()
}

func (s *BoardService) Start() {
	ssh.ListenAndServe("127.0.0.1:2222", s.handleSession)
}
