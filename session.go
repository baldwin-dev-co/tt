package main

import (
	"time"
)

type Session struct {
	start *time.Time
	pauses []*time.Time
	end *time.Time
}

func NewSesh() *Session {
	now := time.Now()
	return &Session{
		start: &now,
		pauses: make([]*time.Time, 0),
	}
}

func (sesh *Session) Pause() {
	now := time.Now()
	sesh.pauses = append(sesh.pauses, &now)
}

func (sesh *Session) Stop() {
	now := time.Now()
	sesh.end = &now
}

