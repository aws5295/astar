package main

import (
	"flag"
	"log"

	"github.com/aws5295/astar"
	"github.com/hajimehoshi/ebiten/v2"
)

var height int
var width int
var frequency int

func init() {
	const (
		defaultHeight = 20
		heightUsage   = "board height"

		defaultWidth = 20
		widthUsage   = "board width"

		defaultFrequency = 5
		frequencyUsage   = "number of moves per second"
	)

	flag.IntVar(&height, "height", defaultHeight, heightUsage)
	flag.IntVar(&height, "h", defaultHeight, heightUsage+" (shorthand)")

	flag.IntVar(&width, "width", defaultWidth, widthUsage)
	flag.IntVar(&width, "w", defaultWidth, widthUsage+" (shorthand)")

	flag.IntVar(&frequency, "frequency", defaultFrequency, frequencyUsage)
	flag.IntVar(&frequency, "f", defaultFrequency, frequencyUsage+" (shorthand)")
}

func main() {
	flag.Parse()
	if height <= 0 || width <= 0 {
		log.Fatalln("Invalid height:", height, " or width:", width)
	}
	if frequency < 1 || frequency > 60 {
		log.Fatalln("Frequency must be between 1 and 60: ", frequency)
	}

	game, err := astar.NewGame(height, width, frequency)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(astar.ScreenWidth, astar.ScreenHeight)
	ebiten.SetWindowTitle("A*")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
