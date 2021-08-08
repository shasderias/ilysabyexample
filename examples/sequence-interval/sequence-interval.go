package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/timer"
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	// `timer.Interval()` specifies a **sequence** that is a series of beats.
	// `0, 2, 3` means 3 beats spaced 2 beats apart, with the first beat
	//  at 0, i.e. 0, 2, 4.
	p.Sequence(timer.Interval(0, 2, 3), func(ctx context.Context) {

		// `ctx.NewRotation()` specifies a base game rotation event.
		ctx.NewRotation()

		// `ctx.NewZoom()` specifies a base game zoom event.
		ctx.NewZoom()
	})

	p.Dump()
}
