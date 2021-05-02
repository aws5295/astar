# astar

astar is a visual representation of the [A\* algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm) for finding the shortest path between two nodes.

## Installation

This project uses a 2d game library called [ebiten](https://ebiten.org/).
Follow the [instructions](https://ebiten.org/documents/install.html) for installing ebiten for your platform which covers installing:

- Go
- C compiler (for certain platforms)
- Dependencies

## CLI

The tool supports several command line options:

- `-h <N>`, `-height <N>` the height of the graph
- `-w <N>`, `-width <N>` the width of the graph
- `-f <N>`, `-frequency <N>` how often to render the graph solution

Example invocation:
`go run ./cmd/ -h 20 -w 20 -f 5`

## Usage

When the grid opens there will be no walls. The following actions are supported:

- `Left-Click` on a cell to flip it from wall to empty or vice-versa
- `Right-Click` anywhere to randomly add 10 cells to the grid
- Press `Enter` to start running the algorithm
- Press `R` to stop the algorithm and reset the grid

## Legend

| Color  | Meaning |
| ------------- | ------------- |
| Orange  | Start or End node  |
| Black  | Wall  |
| White  | Unexplored node  |
| Red  | Closed Node - full explored  |
| Green  | Open Node - fringe - potential node to explore  |
| Blue  | Shortest path to currently exploring node  |

The value printed in each cell is its `f(n)` value - which is an estimate of the cost to extend that node to the End.

## Examples

| Description   | Example |
| ------------- | ------------- |
| Direct path  | ![astarDirect](https://user-images.githubusercontent.com/12830359/116815614-581ae380-ab2c-11eb-896d-118fc5d60af1.gif)  |
| Path with walls  | ![astarLarge](https://user-images.githubusercontent.com/12830359/116815898-bdbb9f80-ab2d-11eb-8f5d-61237ff453de.gif) |
| No path - exhaustive search |   ![astarNoPath](https://user-images.githubusercontent.com/12830359/116815746-ed1ddc80-ab2c-11eb-920e-2414a5fe46db.gif)  |
