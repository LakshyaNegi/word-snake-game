package snake

import (
	"snake-game/utils"
)

type Snake interface {
	GetBody() []body
	Update()
	GetHead() body
	Move(direction)
	Eat(string)
	RemoveWord(l int)
}

type snake struct {
	Body []body
}

func (s *snake) HeadDirection() direction {
	return s.GetHead().Direction
}

func (s *snake) SetHeadDirection(dir direction) {
	s.Body[0].Direction = dir
}

type body struct {
	Block     utils.Block
	Direction direction
}

func (b *body) GetPos() utils.Block {
	return b.Block
}

func (s *snake) RemoveWord(l int) {
	s.Body = s.Body[:len(s.Body)-l]
}

func NewSnake(x, y int) Snake {
	return &snake{
		Body: []body{
			{
				Block: utils.Block{
					Char:     "*",
					Position: [2]int{x, y},
				},
				Direction: Right,
			},
		},
	}
}

func (s *snake) GetBody() []body {
	return s.Body
}

func (s *snake) GetHead() body {
	return s.Body[0]
}

func (s *snake) Eat(ch string) {
	last := s.Body[len(s.Body)-1].Direction
	pos := s.Body[len(s.Body)-1].Block.Position

	switch last {
	case Up:
		pos[1]--
	case Down:
		pos[1]++
	case Left:
		pos[0]++
	case Right:
		pos[0]--
	}

	s.Body = append(s.Body, body{
		Block:     utils.Block{Char: ch, Position: pos},
		Direction: last,
	})
}

func (s *snake) Move(dir direction) {
	if len(s.Body) > 1 {
		head := s.GetHead().Block.Position
		switch s.HeadDirection() {
		case Up:
			head[1]++
		case Down:
			head[1]--
		case Left:
			head[0]--
		case Right:
			head[0]++
		}

		body1 := s.Body[1]

		if utils.PositionOverlap(head, body1.Block.Position) {
			return
		}
	}

	s.SetHeadDirection(dir)
}

func (s *snake) Update() {
	head := s.GetHead().Block.Position

	switch s.HeadDirection() {
	case Up:
		head[1]++
	case Down:
		head[1]--
	case Left:
		head[0]--
	case Right:
		head[0]++
	}

	// updating rest of the body
	for i := (len(s.Body) - 1); i >= 1; i-- {
		s.Body[i].Block.Position = s.Body[i-1].Block.Position
		s.Body[i].Direction = s.Body[i-1].Direction
	}

	// updating head position
	s.Body[0].Block.Position = head
}
