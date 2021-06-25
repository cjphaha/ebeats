package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type clientOnlyACKer struct {

}

func (c *clientOnlyACKer) AddEvent(event beat.Event, published bool) {

}

func (c *clientOnlyACKer) ACKEvents(n int) {

}

func (c *clientOnlyACKer) Close() {

}
