package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	emptyImage = ebiten.NewImage(3, 3)
	circlePoints = []ebiten.Vertex{}
)

func init() {
	emptyImage.Fill(color.White)
	circlePoints = genUnitCircle(16)
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
}

func (g *Game) Update() error {
	g.vertices = generateCircle(200.0, 200.0, 100.0, 1.0, 0.0, 0.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	indices := []uint16{}
	vertexCount := len(g.vertices) -1
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
	ebiten.SetWindowTitle("Polygons (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}