// The `transform` package comes with many **transformers** you can use to
// construct light ID sequences. One such **transformer** is `transform.Fan()`.
// \
// \
// `transform.Fan()` accepts an integer argument `groupCount`. It then creates
// `groupCount` groups and allocates each light ID to a successive group until
// it reaches the `groupCount`th group. This process is repeated until all
// light IDs are allocated.

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
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentFitBeat, 120, "[]")
	p := ilysa.New(m)

	fbRing := transform.Light(

		// Back lasers in the FitBit environment have 30 light IDs. \
		// `1 - [1,2,3 ... 30]`
		light.NewBasic(p, evt.BackLasers),

		// `1 - [1,3,5 ... 29]` \
		// `2 - [2,4,6 ... 30]`
		transform.Fan(2),

		// `1 - [1]`\
		// `2 - [3]`\
		// `3 - [5]`\
		// ...\
		// `29 - [28]`\
		// `30 - [30]`\
		transform.DivideSingle(),
	)

	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.Light(fbRing, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithLight(evt.LightType(evt.LightRedOn)))
		})
	})

	p.Dump()
}
