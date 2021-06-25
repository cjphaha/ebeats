package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type ackerList []beat.ACKer

func (l ackerList) AddEvent(event beat.Event, published bool) {

}

func (l ackerList) ACKEvents(n int) {

}

func (l ackerList) Close() {

}
