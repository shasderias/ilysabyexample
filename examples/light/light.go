// Just as **sequences** and **ranges** let you specify when events should be
// created, **lights** let you specify which lighting event types and light IDs
// to trigger.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/timer"
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	// You can specify lighting event types and light IDs manually like so, but
	// unless you are trying achieve a very specific effect, it's often easier
	// to work with **lights**.
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewRGBLighting(
			evt.WithLight(evt.RightRotatingLasers),
			evt.WithLightValue(evt.LightBlueOn),
			evt.WithLightID(lightid.New(1, 2, 3)),
		)
	})

	// `light.NewBasic()` takes two arguments, an Ilysa project and a light
	// event type, and returns a **basic light**. **Basic lights** don't set
	// the lightID field when used and are essentially equal to a base game
	// lighting event.
	leftLasers := light.NewBasic(p, evt.LeftRotatingLasers)

	p.Sequence(timer.Beat(0), func(ctx context.Context) {

		// Use a **light** by calling `ctx.Light()` from within a **sequence**
		// or a **range** context.
		ctx.Light(leftLasers, func(ctx context.LightContext) {

			// All lighting events created using `ctx` will have their event
			// type appropriately set.
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightBlueOn),
			)
		})
	})

	p.Dump()
}
