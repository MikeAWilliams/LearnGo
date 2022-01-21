package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	pointW       = 50
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
	p1Color    color.Color
	p2Color    color.Color
	p1Points   []point
	p2Points   []point
	isP1turn   bool
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
	result.p1Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	result.p2Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	result.isP1turn = true
	return &result
}

func (g *Game) addPointToCorrectPlayer() {
	currentX, currentY := ebiten.CursorPosition()
	point := point{x: float64(currentX - pointW/2), y: float64(currentY - pointW/2)}
	if g.isP1turn {
		g.p1Points = append(g.p1Points, point)
	} else {
		g.p2Points = append(g.p2Points, point)
	}
	g.isP1turn = !g.isP1turn
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.addPointToCorrectPlayer()
	}
	return nil
}

func (g *Game) drawBoard(canvas *ebiten.Image) {
	for _, line := range g.board {
		ebitenutil.DrawLine(canvas, line.p1.x, line.p1.y, line.p2.x, line.p2.y, g.boardColor)
	}
}

func (g *Game) drawPoints(canvas *ebiten.Image) {
	for _, p := range g.p1Points {
		ebitenutil.DrawRect(canvas, p.x, p.y, pointW, pointW, g.p1Color)
	}
	for _, p := range g.p2Points {
		ebitenutil.DrawRect(canvas, p.x, p.y, pointW, pointW, g.p2Color)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
	g.drawPoints(screen)
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
