package game_test

import (
	"testing"

	"github.com/MikeAWilliams/LearnGo/tree/master/graphics/lineGame/game"
	"github.com/stretchr/testify/require"
)

func TestNearestPointBoardIsSingleLineUp(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(0), Y: float64(1)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	found, _ := testObject.NearestPoint(game.Point{X: float64(0), Y: float64(3)})
	require.False(t, found)

	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(1.1)})
	require.False(t, found)

	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(-0.1)})
	require.False(t, found)

	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(0.0)})
	require.True(t, found)

	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(1.0)})
	require.True(t, found)

	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(0.5)})
	require.True(t, found)
}
