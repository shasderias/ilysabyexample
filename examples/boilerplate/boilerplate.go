// Click the gopher at the top right hand of each code snippet to run and edit
// the snippet in your browser. \
// \
// Before getting started, you need to provide Ilysa with a map. Ilysa uses
// data in the map - specifically the map's environment, BPM and the position
// of the BPM blocks to do its work.
package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

func main() {
	// For the purposes of these tutorials, we do not have access to the
	// filesystem so we use `beatsaber.NewMockMap()` to create a fake map.
	// For running Ilysa on your computer, refer to the following link for the
	// boilerplate to load a map from disk:
	// [Getting Started](https://github.com/shasderias/ilysa/blob/master/examples/getting-started/main.go).
	m, _ := beatsaber.NewMockMap(

		// Origins Environment, 120 BPM, no BPM blocks.
		beatsaber.EnvironmentOrigins, 120, "[]")

	// Create an Ilysa project that uses the mock map.
	p := ilysa.New(m)

	// Code to specify events goes here.

	// Print the JSON for the events specified.
	p.Dump()
}
