// Program haiku is a simple generator of haiku-style names for whatever your
// naming needs.
package main

import (
	"flag"
	"fmt"

	"github.com/cfunkhouser/haiku"
)

var (
	withNumber = flag.Bool("number", false, "If true, include a number in the haiku")
)

func main() {
	flag.Parse()

	if *withNumber {
		fmt.Println(haiku.Custom(haiku.WithComponents(haiku.Adjective, haiku.Noun, haiku.Number)))
		return
	}
	fmt.Println(haiku.Simple())
}
