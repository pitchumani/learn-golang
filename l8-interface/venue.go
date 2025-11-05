package main

type Performer interface {
	Perform()
}

func PerformAtVenue(p Performer) {
	p.Perform()
}
