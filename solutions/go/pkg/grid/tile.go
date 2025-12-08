package grid

// Meta struct for tracking tiles in a grid.
type Tile struct {
	Position Coordinate // The position of the tile.
	Value    rune       // The symbol value of the tile.
	Parent   *Tile      // The parent node for obtaining the shortest path to this node.
}

// A unified struct for tracking coordinates in a 2D grid.
type Coordinate struct {
	X int
	Y int
}

// A unified struct for tracking points in a 3D grid.
// TODO: Candidate for separation into dedicated package for working in 3D.
type Vector struct {
	X int
	Y int
	Z int
}
