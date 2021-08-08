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
	p := ilysa.New(m)

	// `timer.Seq()` specifies a **sequence** that is an arbitrary sequence
	// of beats. `timer.Seq()` also takes a ghost beat (5) that does not form
	// part of the **sequence**. We will see why a ghost beat is useful later.
	p.Sequence(timer.Seq([]float64{0, 1.5, 2, 4}, 5), func(ctx context.Context) {

		// `ctx.NewLaser()` specifies a laser speed event.
		ctx.NewLaser(
			evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithIntValue(1),
		)
	})

	p.Dump()
}
