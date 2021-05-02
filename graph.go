package astar

import (
	"fmt"
	"math"
)

// TableEntry is a row in the A* table.
type TableEntry struct {
	Node *Tile
	// Distance from start.
	G *float64
	// Heuristic distance from end.
	H float64
	// For tracking the path to how we got here.
	PreviousVertex *Tile
}

// F is the A* distance function - g() + h().
func (te *TableEntry) F() float64 {
	if te.G == nil {
		return te.H
	}
	return *te.G + te.H
}

func FloatPtr(f float64) *float64 {
	return &f
}

type AStarGraph struct {
	Board   *Board
	Entries map[*Tile]*TableEntry

	OpenNodes   map[*Tile]struct{}
	ClosedNodes map[*Tile]struct{}
	CurrentNode *Tile
}

func NewAStarGraph(b *Board) *AStarGraph {
	g := &AStarGraph{
		Board:       b,
		OpenNodes:   map[*Tile]struct{}{},
		ClosedNodes: map[*Tile]struct{}{},
	}

	entries := map[*Tile]*TableEntry{}
	for _, tile := range b.tiles {
		entries[tile] = &TableEntry{
			Node:           tile,
			G:              nil,
			H:              tile.HeuristicDistanceFrom(b.End()),
			PreviousVertex: nil,
		}
	}
	g.Entries = entries

	// Initialize the graph - open the neighbors of start.
	for _, neighbor := range g.Neighbors(b.Start()) {
		g.OpenNodes[neighbor.Tile] = struct{}{}
		neighbor.Tile.SetKind(TypeOpen)
		entry := g.Entries[neighbor.Tile]
		entry.G = &neighbor.Cost
		entry.PreviousVertex = b.Start()
		entry.Node.value = fmt.Sprintf("%.1f", entry.F())
	}

	return g
}

// MarkPath resets the isPath value for every tile in the grid.
// It then paints the path to the CurrentNode in our graph traversal.
// This is very inefficient, and not necessary ffor solving A*.
// Instead we do this to give the user a visual representation of what the
// algorithm is "doing".
func (a *AStarGraph) MarkPath() {
	for _, tile := range a.Board.tiles {
		tile.isPath = false
	}

	entry := a.Entries[a.CurrentNode]
	tile := entry.PreviousVertex

	for tile != nil && tile != a.Board.Start() {
		tile.isPath = true
		entry := a.Entries[tile]
		tile = entry.PreviousVertex
	}
}

// Step explores the next "Open" node with the lowest F-value.
// If that node is End, then we have reached the end and found the shortest path.
func (a *AStarGraph) Step() {
	if a.CurrentNode == a.Board.End() {
		a.MarkPath()
		return
	}

	// Get the open node with the lowest f-value.
	smallestF := math.MaxFloat64
	var tile *Tile
	for t := range a.OpenNodes {
		entry := a.Entries[t]
		if entry.F() < smallestF {
			smallestF = entry.F()
			tile = t
		}
	}

	a.CurrentNode = tile
	currentEntry := a.Entries[tile]
	if a.CurrentNode == a.Board.End() || a.CurrentNode == nil {
		// We have found the shortest path.
		return
	}

	for _, neighbor := range a.Neighbors(tile) {
		// Open the node.
		a.OpenNodes[neighbor.Tile] = struct{}{}
		neighbor.Tile.SetKind(TypeOpen)

		// Update the entries table if the new f-value is smaller.
		entry := a.Entries[neighbor.Tile]
		if entry.G == nil || entry.F() > *currentEntry.G+neighbor.Cost {
			entry.G = FloatPtr(*currentEntry.G + neighbor.Cost)
			entry.Node.value = fmt.Sprintf("%.1f", entry.F())
			entry.PreviousVertex = tile
		}
	}

	// Close the tile.
	delete(a.OpenNodes, tile)
	a.ClosedNodes[tile] = struct{}{}
	tile.SetKind(TypeClosed)

	// Draw the path on the map.
	a.MarkPath()
}

// EdgeTo represents the edge from a tile to its valid neighbors.
type EdgeTo struct {
	Tile *Tile
	Cost float64
}

// Neighbors returns all neighbors to the tile that are not a wall or closed.
// Neighbors in cardinal directions have a cost of 1 (Lef, Right, Up, Down).
// Neighbors in corners have a cost of sqrt2 (TopRight, TopLeft, BotRight, BotLeft).
// Walls and closed (fully visited) nodes are not returned.
func (a *AStarGraph) Neighbors(tile *Tile) []*EdgeTo {
	result := []*EdgeTo{}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			neighbor, ok := a.isNeighbor(tile.x+i, tile.y+j, tile)
			if ok {
				cost := math.Sqrt2
				if i == 0 || j == 0 {
					cost = 1.0
				}
				result = append(result, &EdgeTo{Tile: neighbor, Cost: cost})
			}
		}
	}

	return result
}

// isNeighbor takes a tile, and a relative position to the tile.
// If the neighbor is on the board, not closed, and not a wall we return it.
func (a *AStarGraph) isNeighbor(i, j int, tile *Tile) (*Tile, bool) {
	// Neighbor is off the grid.
	if i < 0 || i >= a.Board.width {
		return nil, false
	}
	if j < 0 || j >= a.Board.height {
		return nil, false
	}

	// Neighbor is the tile itself.
	if tile.x == i && tile.y == j {
		return nil, false
	}

	index := a.Board.Index(i, j)
	neighbor := a.Board.tiles[index]

	if neighbor.kind == TypeWall {
		return nil, false
	}

	if _, closed := a.ClosedNodes[neighbor]; closed {
		return nil, false
	}
	return neighbor, true
}
