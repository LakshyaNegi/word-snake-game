package game

import (
	"fmt"
	"log"
	"math/rand"
	"snake-game/screen"
	"snake-game/snake"
	"time"

	"snake-game/utils"
	"strings"

	"github.com/mattn/go-tty"
)

const (
	Selection = 3
	X         = 50
	Y         = 20
	WordLen   = 3
)

type Game interface {
	IsRunning() bool
	Draw()
	Before()
	Update()
	Score() int
}

type game struct {
	isRunning bool
	fps       int
	score     int
	screen    screen.Screen
	letters   []utils.Block
	busyPos   map[int]map[int]bool
	words     []string
	selection int
	snake     snake.Snake
}

func NewGame() Game {
	words := []string{}
	letters := []utils.Block{}
	busyPos := make(map[int]map[int]bool)

	// snake starting pos
	busyPos[X/2] = map[int]bool{}
	busyPos[X/2][Y/2] = true

	for range Selection {
		gen := utils.GenerateWord(WordLen)
		words = append(words, gen)

		for _, c := range strings.Split(gen, "") {
			pos := utils.RandomPosition(X, Y)
			for {
				if !busyPos[pos[0]][pos[1]] && pos[0] != X/2 && pos[1] != Y/2 {
					break
				}

				pos = utils.RandomPosition(X, Y)
			}

			busyPos[pos[0]] = make(map[int]bool)
			busyPos[pos[0]][pos[1]] = true

			letters = append(letters, *utils.NewBlock(c, pos[0], pos[1]))
		}
	}

	return &game{
		isRunning: true,
		fps:       150,
		screen:    screen.NewScreen(),
		words:     words,
		letters:   letters,
		busyPos:   busyPos,
		selection: rand.Intn(Selection),
		snake:     snake.NewSnake(X/2, Y/2),
	}
}

func (g *game) IsRunning() bool {
	return g.isRunning
}

func (g *game) Score() int {
	return g.score
}

func (g *game) Draw() {
	// clear screen
	g.screen.Clear()

	// draw horizontal walls
	for i := range X + 1 {
		// log.Printf("horizontal %d", i)
		g.screen.DrawAt("#", i, 0)
		g.screen.DrawAt("#", i, Y)
	}

	// draw vertical walls
	for i := range Y + 1 {
		// log.Printf("vertical %d", i)
		g.screen.DrawAt("#", 0, i)
		g.screen.DrawAt("#", X, i)
	}

	// draw letters
	for _, block := range g.letters {
		// log.Println(block)
		g.screen.DrawAt(block.Char, block.Position[0], block.Position[1])
	}

	// draw snake
	for _, body := range g.snake.GetBody() {
		g.screen.DrawAt(body.GetPos().Char, body.GetPos().Position[0], body.GetPos().Position[1])
	}

	// print words
	word := "*" + g.words[g.selection]
	g.screen.DrawAt(fmt.Sprintf("Word: %s", word), 0, Y+1)

	// flush on screen
	g.screen.Flush()

	time.Sleep(time.Duration(g.fps) * time.Millisecond)
}

func (g *game) Before() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		ch, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		// process input
		// UP, DOWN, RIGHT, LEFT == [A, [B, [C, [D
		// we ignore the escape character [
		switch ch {
		case 'A':
			g.snake.Move(snake.Down)
		case 'B':
			g.snake.Move(snake.Up)
		case 'C':
			g.snake.Move(snake.Right)
		case 'D':
			g.snake.Move(snake.Left)
		}
	}
}

func (g *game) Update() {
	g.snake.Update()

	// wall collision detection
	head := g.snake.GetHead()

	if head.GetPos().Position[0] == 1 ||
		head.GetPos().Position[0] == X ||
		head.GetPos().Position[1] == 1 ||
		head.GetPos().Position[1] == Y {
		g.isRunning = false
	}

	// eat letter
	for i := range g.letters {
		if utils.PositionOverlap(head.GetPos().Position, g.letters[i].Position) {
			// add letter to body
			g.snake.Eat(g.letters[i].Char)

			// remove letter from game
			g.letters = append(g.letters[:i], g.letters[i+1:]...)

			break
		}
	}

	// check if word is complete
	currWord := g.words[g.selection]
	snakeWord := ""

	for _, body := range g.snake.GetBody() {
		snakeWord += body.Block.Char
	}

	if len(snakeWord) > WordLen {
		snakeWord = snakeWord[len(snakeWord)-len(currWord):]
	}

	if currWord == snakeWord {
		g.snake.RemoveWord(WordLen)
		g.score += 1
		g.ResetStage()
	}
}

func (g *game) ResetStage() {
	words := []string{}
	letters := []utils.Block{}
	busyPos := make(map[int]map[int]bool)

	// snake starting pos
	busyPos[X/2] = map[int]bool{}
	busyPos[X/2][Y/2] = true

	for range Selection {
		gen := utils.GenerateWord(WordLen)
		words = append(words, gen)

		for _, c := range strings.Split(gen, "") {
			pos := utils.RandomPosition(X, Y)
			for {
				if !busyPos[pos[0]][pos[1]] && pos[0] != X/2 && pos[1] != Y/2 {
					break
				}

				pos = utils.RandomPosition(X, Y)
			}

			busyPos[pos[0]] = make(map[int]bool)
			busyPos[pos[0]][pos[1]] = true

			letters = append(letters, *utils.NewBlock(c, pos[0], pos[1]))
		}
	}

	g.words = words
	g.letters = letters
	g.busyPos = busyPos
	g.selection = rand.Intn(Selection)
	g.fps -= 10
}
