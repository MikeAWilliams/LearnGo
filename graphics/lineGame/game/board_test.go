package game_test

import (
	"testing"

	"github.com/MikeAWilliams/LearnGo/tree/master/graphics/lineGame/game"
	"github.com/stretchr/testify/require"
)

func TestNearestPoint_BoardIsSingleLineUp_TestPointIsOnSegment(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(0), Y: float64(1)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	// above
	found, _ := testObject.NearestPoint(game.Point{X: float64(0), Y: float64(3)})
	require.False(t, found)
	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(1.1)})
	require.False(t, found)

	// below
	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(-0.1)})
	require.False(t, found)

	// inside
	found, rp := testObject.NearestPoint(game.Point{X: float64(0), Y: float64(0.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(1.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 1.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(0.5)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.5, rp.Y)
}

func TestNearestPoint_BoardIsSingleLineUp_TestPointIsRightOfSegment(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(0), Y: float64(1)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	// above
	found, _ := testObject.NearestPoint(game.Point{X: float64(1), Y: float64(3)})
	require.False(t, found)
	found, _ = testObject.NearestPoint(game.Point{X: float64(1), Y: float64(1.1)})
	require.False(t, found)

	// below
	found, _ = testObject.NearestPoint(game.Point{X: float64(1), Y: float64(-0.1)})
	require.False(t, found)

	// inside
	found, rp := testObject.NearestPoint(game.Point{X: float64(1), Y: float64(0.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(1), Y: float64(1.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 1.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(1), Y: float64(0.5)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.5, rp.Y)
}

func TestNearestPoint_BoardIsSingleLineUp_TestPointIsLeftOfSegment(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(0), Y: float64(1)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	// above
	found, _ := testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(3)})
	require.False(t, found)
	found, _ = testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(1.1)})
	require.False(t, found)

	// below
	found, _ = testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(-0.1)})
	require.False(t, found)

	// inside
	found, rp := testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(0.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(1.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 1.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(-1), Y: float64(0.5)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.5, rp.Y)
}

func TestNearestPoint_BoardIsSingleLineRight(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(1), Y: float64(0)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	// above
	found, _ := testObject.NearestPoint(game.Point{X: float64(3), Y: float64(0)})
	require.False(t, found)
	found, _ = testObject.NearestPoint(game.Point{X: float64(1.1), Y: float64(0)})
	require.False(t, found)

	// below
	found, _ = testObject.NearestPoint(game.Point{X: float64(-0.1), Y: float64(0)})
	require.False(t, found)

	// inside
	found, rp := testObject.NearestPoint(game.Point{X: float64(0.0), Y: float64(0.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(1.0), Y: float64(0)})
	require.True(t, found)
	require.Equal(t, 1.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(0.5), Y: float64(0)})
	require.True(t, found)
	require.Equal(t, 0.5, rp.X)
	require.Equal(t, 0.0, rp.Y)
}
