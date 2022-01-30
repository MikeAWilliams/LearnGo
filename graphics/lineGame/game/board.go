package game

import (
	"fmt"

	"github.com/ungerik/go3d/vec2"
)

func NewBoard(lines []Segment) Board {
	internalLines := make([]segment, len(lines))
	for index, line := range lines {
		internalLines[index] = newSegment(vecFromPoint(line.P1), vecFromPoint(line.P2))
	}
	return Board{lines: internalLines}
}

type Board struct {
	lines []segment
}

func (b *Board) NearestPoint(source Point) (bool, *Point) {
	sourceVec := vecFromPoint(source)
	for _, line := range b.lines {
		p1Tocource := vec2.Sub(&sourceVec, &line.p1)
		fmt.Printf("p1ToSource (%v, %v)\n", p1Tocource[0], p1Tocource[1])
		fmt.Printf("line.dir (%v, %v)\n", line.dir[0], line.dir[1])
		dist := vec2.Dot(&line.dir, &p1Tocource)
		fmt.Println(dist)
		if dist > 1 {
			return false, nil
		}
		if dist < 0 {
			return false, nil
		}
	}
	return true, &Point{}
}
