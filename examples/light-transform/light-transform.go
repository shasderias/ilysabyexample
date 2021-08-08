// Ilysa lets you apply **transforms** to lights. The typical Ilysa workflow
// involves creating **basic lights** and transforming them to achieve the
// results you want.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	// Origins environment back lasers.
	backLasers := light.NewBasic(p, evt.BackLasers)

	// `transform.Light()` accepts a light and one or more **transformers**, and
	// returns the light with the transforms applied.
	backLasersDiv := transform.Light(backLasers,

		// `transform.DivideSingle()` is a **transformer** that divides each
		// of a light's light IDs into individual light IDs. In the Origins
		// environment, back lasers have 4 light IDs and we can therefore
		// pictue `backLasers` like so: \
		// \
		// `1 - [1, 2, 3, 4]` \
		// \
		// and we can picture `backLasersDiv` like so: \
		// \
		// `1 - [1]`\
		// `2 - [2]`\
		// `3 - [3]`\
		// `4 - [4]`\
		transform.DivideSingle(),
	)

	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.Light(backLasersDiv, func(ctx context.LightContext) {
			// Note that a lighting event is created for each light ID.
			ctx.NewRGBLighting(evt.WithLight(evt.LightType(evt.LightRedOn)))
		})
	})

	// With a little more work, we can make the light IDs light up in order.
	p.Sequence(timer.Beat(2), func(ctx context.Context) {
		ctx.Light(backLasersDiv, func(ctx context.LightContext) {
			// Save the event generated to the variable `e`.
			e := ctx.NewRGBLighting(evt.WithLight(evt.LightType(evt.LightRedOn)))

			// All events generated by Ilysa have the `Apply()` method. The
			// `Apply()` method takes one or more options (the same options
			// you use when creating events) and applies them to the event. \
			// \
			// The option `evt.WithBOffset()` offsets the event's time by the
			// specified amount. \
			// \
			// `ctx.LightIDT()` is like `ctx.T()`, but for light IDs.
			e.Apply(evt.WithBOffset(ctx.LightIDT()))
		})
	})

	p.Dump()
}
