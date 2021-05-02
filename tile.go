package astar

import (
	"image/color"
	"log"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	TileSize   = 40
	TileMargin = 2
)

// TileType is used to dtermine how to render the tile.
// Some types are immutable - like TypeStart and TypeEnd.
// SetKind() should be used to change a Tile's Type.
type TileType int

const (
	TypeBlank TileType = iota
	TypeStart
	TypeEnd
	TypeWall
	TypeOpen
	TypeClosed
)

var (
	tileImage      = ebiten.NewImage(TileSize, TileSize)
	mplusSmallFont font.Face
)

func init() {
	tileImage.Fill(color.White)
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Tile contains the information necssary to render a tile on the board.
// The color and contents are determined based on the kind.
type Tile struct {
	kind  TileType
	value string
	// isPath is separate from kind. isPath represents whether the tile is part of
	// the current path based on the current run of A*.
	isPath bool
	x      int
	y      int
}

func (t *Tile) Color() color.RGBA {
	if t.isPath {
		return LightBlue
	}

	switch t.kind {
	case TypeBlank:
		return White
	case TypeStart:
		return Orange
	case TypeEnd:
		return Orange
	case TypeWall:
		return Black
	case TypeClosed:
		return Red
	case TypeOpen:
		return Green
	}

	return White
}

func (t *Tile) SetKind(kind TileType) {
	if t.kind == TypeStart || t.kind == TypeEnd {
		return
	}
	t.kind = kind
}

func (t *Tile) TryFlipWall() {
	if t.kind == TypeWall {
		t.SetKind(TypeBlank)
		return
	}

	if t.kind == TypeBlank {
		t.SetKind(TypeWall)
		return
	}
}

func (t *Tile) Text() string {
	switch t.kind {
	case TypeStart:
		return "Start"
	case TypeEnd:
		return "End"
	}

	return t.value
}

// HeuristicDistanceFrom is the h() function for A*.
// Here we use the pythagorean distance.
func (t *Tile) HeuristicDistanceFrom(other *Tile) float64 {
	a := float64(t.x - other.x)
	b := float64(t.y - other.y)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// NewTile creates a new Tile object.
func NewTile(x, y int) *Tile {
	return &Tile{
		kind: TypeBlank,
		x:    x,
		y:    y,
	}
}

// Draw draws the current tile to the given boardImage.
func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.x, t.y
	op := &ebiten.DrawImageOptions{}
	x := i*TileSize + (i+1)*TileMargin
	y := j*TileSize + (j+1)*TileMargin
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.Scale(colorToScale(t.Color()))
	boardImage.DrawImage(tileImage, op)

	bound, _ := font.BoundString(mplusSmallFont, t.Text())
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	x = x + (TileSize-w)/2
	y = y + (TileSize-h)/2 + h
	text.Draw(boardImage, t.Text(), mplusSmallFont, x, y, Black)
}
