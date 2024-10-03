package users

import (
	"sync"

	"github.com/gedex/bp3d/bp3d"
)

const (
	LOCAL_USESSION = "us"
)

type Credit struct {
	Name   string
	Amount int
}

type USession struct {
	Id      string
	Packer  *bp3d.Packer
	Credits map[string]*Credit
}

// for io.Closer
func (s *USession) Close() error {
	return nil
}

func (s *USession) AffirmPacker() *bp3d.Packer {
	if nil == s.Packer {
		s.Packer = bp3d.NewPacker()

	}
	return s.Packer
}

type USessions struct {
	m sync.Map
}

func (s *USessions) GetById(id string) *USession {
	if r, ok := s.m.Load(id); ok {
		return r.(*USession)
	} else {
		return nil
	}
}
func (s *USessions) SetById(id string, aUS *USession) {
	s.m.Store(id, aUS)
}
