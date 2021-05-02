package astar

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 1600
	ScreenHeight = 1000
)

// Game represents a game state.
type Game struct {
	board            *Board
	boardImage       *ebiten.Image
	counter          int
	input            *Input
	ticksPerInterval int
}

// NewGame generates a new Game object.
func NewGame(h, w, movesPerSecond int) (*Game, error) {
	g := &Game{
		counter:          0,
		input:            &Input{},
		ticksPerInterval: 60 / movesPerSecond,
	}
	var err error
	g.board, err = NewBoard(h, w)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()

	if g.input.ResetPressed {
		b, err := NewBoard(g.board.height, g.board.width)
		if err != nil {
			return err
		}
		g.board = b
		g.input.CopyAndReset()
		return nil
	}

	g.counter++
	if g.counter%g.ticksPerInterval != 0 {
		return nil
	}
	g.counter = 0

	// Adjust coordinates to be relative to board.
	bw, bh := g.board.Size()
	g.input.MouseX = g.input.MouseX - (ScreenWidth-bw)/2
	g.input.MouseY = g.input.MouseY - (ScreenHeight-bh)/2

	g.board.Update(g.input.CopyAndReset())
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage = ebiten.NewImage(w, h)
	}
	screen.Fill(BackgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
