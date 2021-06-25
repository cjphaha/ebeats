package beat

type ACKer interface {

	AddEvent(event Event, published bool)

	ACKEvents(n int)

	Close()
}