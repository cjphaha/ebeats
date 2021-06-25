package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type countACKer func(int)

func (countACKer) AddEvent(_ beat.Event, _ bool) {}
func (fn countACKer) ACKEvents(n int)            { fn(n) }
func (countACKer) Close()                        {}