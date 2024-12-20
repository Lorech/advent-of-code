package grid

// Meta struct for tracking tiles in a grid.
type Tile struct {
	Position Coordinates // The position of the tile.
	Value    rune        // The symbol value of the tile.
	Parent   *Tile       // The parent node for obtaining the shortest path to this node.
}

// A unified struct for tracking coordinates in a grid.
type Coordinates struct {
	X int
	Y int
}
