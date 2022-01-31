package main

import (
	"image/color"
	"log"

	"github.com/MikeAWilliams/LearnGo/tree/master/graphics/lineGame/game"
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

type Game struct {
	board       game.Board
	boardToDraw []game.Segment
	boardColor  color.Color
	p1Color     color.Color
	p2Color     color.Color
	p1Points    []game.Point
	p2Points    []game.Point
	isP1turn    bool
}

func NewGame() *Game {
	result := Game{}
	// make board a square that is a little smaller than the screen
	const offset = 50
	ul := game.Point{X: offset, Y: offset}
	ur := game.Point{X: screenWidth - offset, Y: offset}
	ll := game.Point{X: offset, Y: screenHeight - offset}
	lr := game.Point{X: screenWidth - offset, Y: screenHeight - offset}
	result.boardToDraw = append(result.boardToDraw, game.Segment{P1: ul, P2: ur})
	result.boardToDraw = append(result.boardToDraw, game.Segment{P1: ur, P2: lr})
	result.boardToDraw = append(result.boardToDraw, game.Segment{P1: lr, P2: ll})
	result.boardToDraw = append(result.boardToDraw, game.Segment{P1: ll, P2: ul})
	result.board = game.NewBoard(result.boardToDraw)
	result.boardColor = color.White
	result.p1Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	result.p2Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	result.isP1turn = true
	return &result
}

func (g *Game) addPointToCorrectPlayer() {
	currentX, currentY := ebiten.CursorPosition()
	point := game.Point{X: float64(currentX - pointW/2), Y: float64(currentY - pointW/2)}
	found, boardPoint := g.board.NearestPoint(point)
	if !found {
		return
	}
	if g.isP1turn {
		g.p1Points = append(g.p1Points, *boardPoint)
	} else {
		g.p2Points = append(g.p2Points, *boardPoint)
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
	for _, line := range g.boardToDraw {
		ebitenutil.DrawLine(canvas, line.P1.X, line.P1.Y, line.P2.X, line.P2.Y, g.boardColor)
	}
}

func (g *Game) drawPoints(canvas *ebiten.Image) {
	for _, p := range g.p1Points {
		ebitenutil.DrawRect(canvas, p.X, p.Y, pointW, pointW, g.p1Color)
	}
	for _, p := range g.p2Points {
		ebitenutil.DrawRect(canvas, p.X, p.Y, pointW, pointW, g.p2Color)
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
