package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")

	// `p` represents an Ilysa project. You call methods on it to specify the
	// events to generate.
	p := ilysa.New(m)

	// At beat 1 ...
	p.Sequence(timer.Beat(1), func(ctx context.Context) {

		// ... create a left rotating lasers blue fade event, ...
		ctx.NewLighting(
			// *Event options are contained in the `evt` package, and all start with "With".*
			evt.WithLight(evt.LeftRotatingLasers),
			evt.WithLightValue(evt.LightBlueFade),
		)

		// ... and a laser speed event.
		ctx.NewLaser(
			evt.WithIntValue(3),
		)
	})

	p.Dump()
}
