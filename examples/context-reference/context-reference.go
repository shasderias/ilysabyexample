// This example documents all the possible options for the methods that create
// Chroma events. For Chroma specific details on what the options do, please refer
// to https://github.com/Aeroluna/Chroma.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
)

var (
	Red = colorful.MustParseHex("#FF0000")
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	l := light.NewBasic(p, evt.BackLasers)

	p.Sequence(timer.Seq([]float64{0, 1, 2, 4}, 5), func(ctx context.Context) {

		// `SeqT` is the current time in the current sequence, on a 0-1 scale. As a special case, SeqT is 1 if the
		// current sequence has only 1 step.
		ctx.SeqT()

		// `SeqOrdinal` is the ordinal number of the current iteration of the current sequence, starting from 0.
		ctx.SeqOrdinal()

		// `SeqLen` is the number of steps in the current sequence.
		ctx.SeqLen()

		// `SeqNextB` is the beat number of the next step in the current sequence.
		ctx.SeqNextB()

		// `SeqNextBOffset` is the difference in beats between the current step and the next step.
		// As a special case, `SeqNextBOffset` returns the difference between the last beat and the ghost beat, if this is the last step.
		ctx.SeqNextBOffset()

		// `SeqPrevB` is the beat number of the previous step in the current sequence.
		ctx.SeqPrevB()

		// `SeqPrevBOffset` is the difference in beats between the current step and the previous step.
		ctx.SeqPrevBOffset()

		// `SeqFirst` is true if this is the first step in the current sequence.
		ctx.SeqFirst()

		// `SeqLast` is true if this is the last step in the current sequence.
		ctx.SeqLast()

		ctx.Range(timer.Rng(0, 0.5, 6, ease.Linear), func(ctx context.Context) {

			// `B` is the current beat.
			ctx.B()

			// `T` is the current time in the current range (or sequence), on a 0-1 scale. As a special case, `T` is 1 if the
			// current range only has 1 step.
			ctx.T()

			// `Ordinal` is the ordinal number of the current step in the current range (or sequence), starting from 0.
			ctx.Ordinal()

			// `StartB` is the starting beat of the current range.
			ctx.StartB()

			// `EndB` is the ending beat of the current range.
			ctx.EndB()

			// `Duration` is the duration of the current range.
			ctx.Duration()

			// `First` is true if this is the first step in the current range.
			ctx.First()

			// `Last` is true if this is the last step in the current range.
			ctx.Last()

			ctx.Light(l, func(ctx context.LightContext) {

				// `LightIDT` is the position of the current light ID for the current light, on a 0-1 scale.
				ctx.LightIDT()

				// `LightIDOrdinal` is the ordinal number of current light ID.
				ctx.LightIDOrdinal()

				// `LightIDLen` is the number of light IDs the current light has.
				ctx.LightIDLen()

				// `LightIDCur` is the current light ID (effectively, `LightIDOrdinal` + 1).
				ctx.LightIDCur()

			})
		})
	})
	p.Dump()
}
