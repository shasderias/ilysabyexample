// Within a **sequence**, the context value `ctx` contains values useful for
// creating lighting effects. A full list of available values can be viewed
// [here](/context-reference).
package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/timer"
)

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	p.Sequence(timer.Seq([]float64{0, 1.5, 2.5, 3, 4}, 6), func(ctx context.Context) {
		fmt.Printf("%4.2f %4.2f %1d %4.2f %t\n",

			// `ctx.B()` returns the current beat.\
			// \
			// `ctx.T()` returns the current position in the **sequence**, on a 0 to 1 scale.\
			// \
			// `ctx.Ordinal()` returns the ordinal number of the current beat, i.e. 0, 1, 2, 3... etc.\
			// \
			// `ctx.SeqNextBOffset()` returns the number of beats to the next beat in the current sequence.
			// As a special case, when using `timer.Seq()` to specify a sequence, on the last beat,
			// `ctx.SeqNextBOffset()` returns the difference between the last beat and the ghost beat. \
			// \
			// `ctx.Last()` returns true if this is the last beat in the sequence.
			ctx.B(), ctx.T(), ctx.Ordinal(), ctx.SeqNextBOffset(), ctx.Last())
	})
}
