package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	maxRadius    = 100
	minRadius    = 1
)

var (
	emptyImage   = ebiten.NewImage(3, 3)
	circlePoints = []ebiten.Vertex{}
)

func init() {
	emptyImage.Fill(color.White)
	circlePoints = genUnitCircle(40)
}

func genUnitCircle(num int) []ebiten.Vertex {
	const (
		centerX = 0
		centerY = 0
		r       = 1
	)

	vs := []ebiten.Vertex{}
	for i := 0; i < num; i++ {
		rate := float64(i) / float64(num)
		cr := 0.0
		cg := 0.0
		cb := 0.0
		vs = append(vs, ebiten.Vertex{
			DstX:   float32(r*math.Cos(2*math.Pi*rate)) + centerX,
			DstY:   float32(r*math.Sin(2*math.Pi*rate)) + centerY,
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(cr),
			ColorG: float32(cg),
			ColorB: float32(cb),
			ColorA: 1,
		})
	}

	vs = append(vs, ebiten.Vertex{
		DstX:   centerX,
		DstY:   centerY,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	})

	return vs
}

func generateCircle(x, y, rad, r, g, b float32) []ebiten.Vertex {
	result := make([]ebiten.Vertex, len(circlePoints))
	for index, pt := range circlePoints {
		result[index].DstX = rad*pt.DstX + x
		result[index].DstY = rad*pt.DstY + y
		result[index].ColorR = r
		result[index].ColorG = g
		result[index].ColorB = b
		result[index].ColorA = 1
	}
	return result
}

type Game struct {
	vertices []ebiten.Vertex
	radius   float32
	x        float32
	y        float32
	vx       float32
	vy       float32
	r        float32
	g        float32
	b        float32
}

func pegColorValue(value float32) float32 {
	if value > 1 {
		return 1
	}
	if value < 0 {
		return 0
	}
	return value
}

func pegRadius(value float32) float32 {
	if value > maxRadius {
		return maxRadius
	}
	if value < minRadius {
		return minRadius
	}
	return value
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.r -= 0.1
		g.r = pegColorValue(g.r)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.r += 0.1
		g.r = pegColorValue(g.r)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.g -= 0.1
		g.g = pegColorValue(g.g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.g += 0.1
		g.g = pegColorValue(g.g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.b -= 0.1
		g.b = pegColorValue(g.b)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		g.b += 0.1
		g.b = pegColorValue(g.b)
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.radius -= 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.radius += 1
	}
	g.radius = pegRadius(g.radius)
	g.x += g.vx
	g.y += g.vy
	if g.x > screenWidth-g.radius || g.x < g.radius {
		g.vx *= -1
	}
	if g.y > screenHeight-g.radius || g.y < g.radius {
		g.vy *= -1
	}

	g.vertices = generateCircle(g.x, g.y, g.radius, g.r, g.g, g.b)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	indices := []uint16{}
	vertexCount := len(g.vertices) - 1
	for i := 0; i < vertexCount; i++ {
		indices = append(indices, uint16(i), uint16(i+1)%uint16(vertexCount), uint16(vertexCount))
	}
	screen.DrawTriangles(g.vertices, indices, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Circle")
	if err := ebiten.RunGame(&Game{vx: 1, vy: 1, x: 100, y: 100, radius: 50, r: 1, g: 1, b: 1}); err != nil {
		log.Fatal(err)
	}
}
