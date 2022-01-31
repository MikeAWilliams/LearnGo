package game

import (
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
	// todo initilize this more intelligently
	var closestPoint *vec2.T
	var closestDist float32
	for _, line := range b.lines {
		p1Tocource := vec2.Sub(&sourceVec, &line.p1)
		dist := vec2.Dot(&line.dir, &p1Tocource) / line.dirDotDir
		if dist > 1 || dist < 0 {
			continue
		}
		toCandidate := line.dir.Scaled(dist)
		candidate := vec2.Add(&line.p1, &toCandidate)
		sourceToCandidate := vec2.Sub(&candidate, &sourceVec)
		distFromSource := sourceToCandidate.Length()
		if nil == closestPoint || distFromSource < closestDist {
			closestDist = distFromSource
			closestPoint = &candidate
		}
	}
	if nil == closestPoint {
		return false, &Point{}
	}
	result := pointFromVec(*closestPoint)
	return true, &result
}
