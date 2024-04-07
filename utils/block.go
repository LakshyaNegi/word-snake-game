package utils

type Block struct {
	Char     string
	Position position
}

func NewBlock(c string, x, y int) *Block {
	return &Block{
		Char:     c,
		Position: [2]int{x, y},
	}
}
