// A common effect is to delay successive light on events for each light ID
// to achieve a "ripple" effect that looks like the light is in motion.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

var (
	Red     = colorful.MustParseHex("#FF0000")
	Green   = colorful.MustParseHex("#00FF00")
	Blue    = colorful.MustParseHex("#0000FF")
	RGBGrad = gradient.New(Red, Blue, Green)
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	l := transform.Light(light.NewBasic(p, evt.BackLasers),
		transform.DivideSingle(),
	)

	p.Range(timer.Rng(2, 3, 10, ease.Linear), func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {

			// This effect can be achieved manually.
			e := ctx.NewRGBLighting(
				evt.WithLight(evt.BackLasers),
				evt.WithLightValue(evt.LightRedOn),
			)

			// Shift the light on event for the 1st light ID by 0 beats,
			// the light on event for the 2nd light ID by 0.2 beats, the
			// light on event for the 3rd light ID by 0.4 beats, etc.
			e.Apply(evt.WithBOffset(float64(ctx.LightIDOrdinal()) * 0.2))

			// The above line can be replaced with this.
			fx.Ripple(ctx, e, 0.2)

			// Alternatively, you can use `fx.RippleT()` which delays
			// each successive event by an amount such that the last event
			// is triggered 1 beat later than the first event.
			fx.RippleT(ctx, e, 1)

		})
	})
}
