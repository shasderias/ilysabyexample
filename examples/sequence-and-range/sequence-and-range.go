// Use a **sequence** for discrete events (e.g. every beat for 4 beats), use
//  a **range** for continuous events (e.g. fade alpha from 1 to 0 over 0.5
// beats).  Most effects require both and Ilysa lets you nest them.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

// Use `colorful.MustParseHex()` to specify colors.
var Red = colorful.MustParseHex("#FF0000")

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	// Create events at beats 0 and 2 ...
	p.Sequence(timer.Seq([]float64{0, 2}, 0), func(ctx context.Context) {

		// ... for each beat in the sequence create 6 events from beats 0 to 0.5.
		ctx.Range(timer.Rng(0, 0.5, 6, ease.Linear), func(ctx context.Context) {

			// `ctx.NewRGBLighting()` specifies Chroma lighting events.
			ctx.NewRGBLighting(
				evt.WithLight(evt.LeftRotatingLasers),
				evt.WithLightValue(evt.LightRedOn),
				evt.WithColor(Red),
				evt.WithAlpha(ctx.T()),
			)
		})
	})

	p.Dump()
}
