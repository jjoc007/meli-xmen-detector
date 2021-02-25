package utils

const (
	ACompleteSequence = "AAAA"
	TCompleteSequence = "TTTT"
	CCompleteSequence = "CCCC"
	GCompleteSequence = "GGGG"
)

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
	Diagonal
	InverseDiagonal
)