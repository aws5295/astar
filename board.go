package astar

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Board represents the game board.
type Board struct {
	width       int
	height      int
	tiles       []*Tile
	graph       *AStarGraph
	initialized bool
}

// NewBoard generates a new Board with giving a size.
func NewBoard(height, width int) (*Board, error) {
	b := &Board{
		height:      height,
		width:       width,
		tiles:       make([]*Tile, height*width),
		initialized: false,
	}

	for i := range b.tiles {
		x, y := b.Coord(i)
		b.tiles[i] = &Tile{
			kind: TypeBlank,
			x:    x,
			y:    y,
		}
	}
	b.Start().SetKind(TypeStart)
	b.End().SetKind(TypeEnd)

	return b, nil
}

// Index returns the position in the array for a coordinate.
// Board stores the grid as a 1-d array, so this helper computes the
// 1-d index based on the 2-d coordinates.
func (b *Board) Index(x, y int) int {
	return y*b.width + x
}

// Coord returns the 2-d coordinates for an array index.
// Board stores the grid as a 1-d array.
func (b *Board) Coord(index int) (int, int) {
	x := index % b.width
	y := index / b.width
	return x, y
}

// Start returns the top left tile in the board which is our starting point.
func (b *Board) Start() *Tile {
	index := b.Index(0, 0)
	return b.tiles[index]
}

// End returns the bottom right tile in thee board, which is our destination.
func (b *Board) End() *Tile {
	index := b.Index(b.width-1, b.height-1)
	return b.tiles[index]
}

func (b *Board) initialize() {
	g := NewAStarGraph(b)
	b.graph = g
	b.initialized = true
}

// addRandomWalls adds numWalls in random locations on the board.
func (b *Board) addRandomWalls(numWalls int) {
	start := b.Start()
	end := b.End()

	for i := 0; i < numWalls; i++ {
		isValidTile := false
		for !isValidTile {
			x := rand.Intn(b.width)
			y := rand.Intn(b.height)
			index := b.Index(x, y)
			tile := b.tiles[index]

			// Start and end are not valid walls.
			if tile == start || tile == end || tile.kind == TypeWall {
				continue
			}

			isValidTile = true
			tile.SetKind(TypeWall)
		}
	}
}

// IsOnBoard reeturns true if the coordinate is on the board.
// This is the relative x and y coordiantes of the mouse, not the cell in the grid.
func (b *Board) IsOnBoard(x, y int) bool {
	xOnBoard := x >= 0 && x <= (b.width*TileSize)+(b.width+1)*TileMargin
	yOnBoard := y >= 0 && y <= (b.height*TileSize)+(b.height+1)*TileMargin

	return xOnBoard && yOnBoard
}

// TileAt returns the tile at the x and y position on the board.
// This is the position of the mouse, not the x and y coordinates of the grid.
func (b *Board) TileAt(x, y int) *Tile {
	xCoord := x / (TileSize + TileMargin)
	yCoord := y / (TileSize + TileMargin)
	index := b.Index(xCoord, yCoord)

	return b.tiles[index]
}

// Update updates the board state.
func (b *Board) Update(input *Input) {
	if b.initialized {
		b.graph.Step()
		return
	}

	if input.EnterPressed {
		b.initialize()
		return
	}

	if input.RightMousePressed {
		b.addRandomWalls(10)
		return
	}

	if !input.LeftMousePressed {
		return
	}

	if b.IsOnBoard(input.MouseX, input.MouseY) {
		tile := b.TileAt(input.MouseX, input.MouseY)
		tile.TryFlipWall()
	}
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := b.width*TileSize + (b.width+1)*TileMargin
	y := b.height*TileSize + (b.height+1)*TileMargin
	return x, y
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(FrameColor)
	for j := 0; j < b.height; j++ {
		for i := 0; i < b.width; i++ {
			op := &ebiten.DrawImageOptions{}
			x := i*TileSize + (i+1)*TileMargin
			y := j*TileSize + (j+1)*TileMargin
			op.GeoM.Translate(float64(x), float64(y))
			boardImage.DrawImage(tileImage, op)
		}
	}

	for _, tile := range b.tiles {
		tile.Draw(boardImage)
	}
}
