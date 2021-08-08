// This example documents all the possible options for the methods that create
// Chroma events. For Chroma specific details on what the options do, please refer
// to https://github.com/Aeroluna/Chroma.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/timer"
)

var (
	Red = colorful.MustParseHex("#FF0000")
)

func main() {
	m, _ := beatsaber.NewMockMap(
		beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	p.Sequence(timer.Beat(0), func(ctx context.Context) {

		// `ctx.NewRGBLighting()` creates Chroma RGB lighting events.
		ctx.NewRGBLighting(
			evt.WithLight(evt.BackLasers),
			evt.WithLight(evt.RingLights),
			evt.WithLight(evt.LeftRotatingLasers),
			evt.WithLight(evt.RightRotatingLasers),
			evt.WithLight(evt.CenterLights),

			evt.WithLightValue(evt.LightOff),
			evt.WithLightValue(evt.LightBlueOn),
			evt.WithLightValue(evt.LightBlueFlash),
			evt.WithLightValue(evt.LightBlueFade),
			evt.WithLightValue(evt.LightRedOn),
			evt.WithLightValue(evt.LightRedFlash),
			evt.WithLightValue(evt.LightRedFade),

			evt.WithColor(Red),

			evt.WithAlpha(3.3),

			// This option is rarely used. See the section on Lights.
			evt.WithLightID(lightid.New(1, 2, 4)),

			// Shifts the event forward by 0.2 beats.
			evt.WithBOffset(0.2),

			// Included for completeness, specify timing using `ctx.Sequence()` or `ctx.Range()` instead.
			evt.WithBeat(0.2),
		)

		// `ctx.NewPreciseLaser()` creates Chroma Precise Laser events.
		ctx.NewPreciseLaser(
			evt.WithLockPosition(true),

			evt.WithPreciseLaserSpeed(3),

			evt.WithLaserDirection(chroma.Clockwise),
			evt.WithLaserDirection(chroma.CounterClockwise),

			// Shifts the event forward by 0.2 beats.
			evt.WithBOffset(0.2),

			// Included for completeness, specify timing using `ctx.Sequence()` or `ctx.Range()` instead.
			evt.WithBeat(0.2),
		)

		// `ctx.NewPreciseRotation()` creates Chroma Precise Rotation events.
		ctx.NewPreciseRotation(
			evt.WithNameFilter("SmallLaneTrackRings"),

			evt.WithReset(true),

			evt.WithRotation(1.0),

			evt.WithRotationStep(45),

			evt.WithProp(1.2),

			evt.WithRotationSpeed(3.0),

			evt.WithRotationDirection(chroma.Clockwise),
			evt.WithRotationDirection(chroma.CounterClockwise),

			evt.WithCounterSpin(true),

			// Shifts the event forward by 0.2 beats.
			evt.WithBOffset(0.2),

			// Included for completeness, specify timing using `ctx.Sequence()` or `ctx.Range()` instead.
			evt.WithBeat(0.2),
		)

		// `ctx.NewPreciseRotation()` creates Chroma Precise Zoom events.
		ctx.NewPreciseZoom(
			evt.WithZoomStep(2.0),

			// Shifts the event forward by 0.2 beats.
			evt.WithBOffset(0.2),

			// Included for completeness, specify timing using `ctx.Sequence()` or `ctx.Range()` instead.
			evt.WithBeat(0.2),
		)
	})

	p.Dump()
}
