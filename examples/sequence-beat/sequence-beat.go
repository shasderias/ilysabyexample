// Generate events by specifying a **sequence** and the events you want to
// generate for that **sequence**. There are many ways to specify a
// **sequence**, `timer.Beat()` is the simplest.
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

	// `timer.Beat()` specifies a **sequence** with a single beat. All events
	// created using ctx will have their beat set to 0.
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLighting(
			evt.WithLight(evt.BackLasers),
			evt.WithLightValue(evt.LightBlueFade),
		)
	})

	p.Dump()
}
