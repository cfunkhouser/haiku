package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/cfunkhouser/haiku"
)

var (
	withNumber = flag.Bool("number", false, "If true, include a number in the haiku")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	if *withNumber {
		fmt.Println(haiku.SimpleWithNumber())
		return
	}
	fmt.Println(haiku.Simple())
}
