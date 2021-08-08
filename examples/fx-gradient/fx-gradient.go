// The `fx` package collects common lighting patterns.
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

	p.Sequence(timer.Beat(1), func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			// `fx.Gradient()` displays a gradient over a light.
			fx.Gradient(ctx, RGBGrad)
		})
	})

	p.Range(timer.Rng(2, 3, 10, ease.Linear), func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			// `fx.ColorSweep()` animates a gradient over a light.
			// The second parameter (0.2) determines how quickly the gradient moves.
			fx.ColorSweep(ctx, 0.2, RGBGrad)
		})
	})

	p.Dump()
}
