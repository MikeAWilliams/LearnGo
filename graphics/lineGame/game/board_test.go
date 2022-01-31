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

func TestNearestPoint_BoardIsSquare(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(1), Y: float64(0)}
	p3 := game.Point{X: float64(1), Y: float64(1)}
	p4 := game.Point{X: float64(0), Y: float64(1)}
	p5 := game.Point{X: float64(0), Y: float64(0)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}, {P1: p2, P2: p3}, {P1: p3, P2: p4}, {P1: p4, P2: p5}})

	//right side
	testPoint := game.Point{X: float64(1.1), Y: float64(0.5)}
	found, result := testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 1.0, result.X)
	require.Equal(t, 0.5, result.Y)
	testPoint.X = float64(0.9)
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 1.0, result.X)
	require.Equal(t, 0.5, result.Y)

	//left side
	testPoint = game.Point{X: float64(0.1), Y: float64(0.5)}
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.0, result.X)
	require.Equal(t, 0.5, result.Y)
	testPoint.X = float64(-0.1)
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.0, result.X)
	require.Equal(t, 0.5, result.Y)

	//top side
	testPoint = game.Point{X: float64(0.5), Y: float64(1.1)}
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.5, result.X)
	require.Equal(t, 1.0, result.Y)
	testPoint.Y = float64(0.9)
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.5, result.X)
	require.Equal(t, 1.0, result.Y)

	//bottom side
	testPoint = game.Point{X: float64(0.5), Y: float64(0.1)}
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.5, result.X)
	require.Equal(t, 0.0, result.Y)
	testPoint.Y = float64(-0.1)
	found, result = testObject.NearestPoint(testPoint)
	require.True(t, found)
	require.Equal(t, 0.5, result.X)
	require.Equal(t, 0.0, result.Y)
}

func TestNearestPoint_BoardIsNonUnitSingleLineUp_TestPointIsOnSegment(t *testing.T) {
	p1 := game.Point{X: float64(0), Y: float64(0)}
	p2 := game.Point{X: float64(0), Y: float64(100)}
	testObject := game.NewBoard([]game.Segment{{P1: p1, P2: p2}})

	// above
	found, _ := testObject.NearestPoint(game.Point{X: float64(0), Y: float64(101)})
	require.False(t, found)

	// below
	found, _ = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(-0.1)})
	require.False(t, found)

	// inside
	found, rp := testObject.NearestPoint(game.Point{X: float64(0), Y: float64(0.0)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 0.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(100)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 100.0, rp.Y)
	found, rp = testObject.NearestPoint(game.Point{X: float64(0), Y: float64(50)})
	require.True(t, found)
	require.Equal(t, 0.0, rp.X)
	require.Equal(t, 50.0, rp.Y)
}
