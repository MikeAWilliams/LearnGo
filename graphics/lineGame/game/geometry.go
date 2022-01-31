package game

import (
	"github.com/ungerik/go3d/vec2"
)

type Point struct {
	X float64
	Y float64
}

func vecFromPoint(point Point) vec2.T {
	return vec2.T{float32(point.X), float32(point.Y)}
}

func pointFromVec(vec vec2.T) Point {
	return Point{X: float64(vec[0]), Y: float64(vec[1])}
}

type Segment struct {
	P1 Point
	P2 Point
}

func newSegment(p1, p2 vec2.T) segment {
	oneToTwo := vec2.Sub(&p2, &p1)
	dirDotDir := vec2.Dot(&oneToTwo, &oneToTwo)
	return segment{p1: p1, p2: p2, dir: oneToTwo, dirDotDir: dirDotDir}
}

type segment struct {
	p1        vec2.T
	p2        vec2.T
	dir       vec2.T
	dirDotDir float32
}
