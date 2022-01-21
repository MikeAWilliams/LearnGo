package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	emptyImage   = ebiten.NewImage(3, 3)
	circlePoints = []ebiten.Vertex{}
)

type point struct {
	x float64
	y float64
}

type line struct {
	p1 point
	p2 point
}

type Game struct {
	board      []line
	boardColor color.Color
}

func NewGame() *Game {
	result := Game{}
	// make board a square that is a little smaller than the screen
	const offset = 50
	ul := point{x: offset, y: offset}
	ur := point{x: screenWidth - offset, y: offset}
	ll := point{x: offset, y: screenHeight - offset}
	lr := point{x: screenWidth - offset, y: screenHeight - offset}
	result.board = append(result.board, line{p1: ul, p2: ur})
	result.board = append(result.board, line{p1: ur, p2: lr})
	result.board = append(result.board, line{p1: lr, p2: ll})
	result.board = append(result.board, line{p1: ll, p2: ul})
	result.boardColor = color.White
	return &result
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) drawBoard(canvas *ebiten.Image) {
	for _, line := range g.board {
		ebitenutil.DrawLine(canvas, line.p1.x, line.p1.y, line.p2.x, line.p2.y, g.boardColor)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("lineGame")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
