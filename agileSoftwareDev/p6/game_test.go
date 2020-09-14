package p6

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGame(t *testing.T) {
	g := NewGame()

	g.Add(5)
	g.Add(4)
	Convey("two throws", t, func() {
		So(9, ShouldEqual, g.GetScore())
		So(9, ShouldEqual, g.ScoreForFrame(1))
	})

	g.Add(7)
	g.Add(2)
	Convey("four throws", t, func() {
		So(18, ShouldEqual, g.GetScore())
		So(18, ShouldEqual, g.ScoreForFrame(2))
	})

	g = NewGame()
	g.Add(3)
	g.Add(7)
	g.Add(3)
	g.Add(2)
	Convey("simple spare 补中", t, func() {
		So(13, ShouldEqual, g.ScoreForFrame(1))
		So(18, ShouldEqual, g.ScoreForFrame(2))
	})

	g = NewGame()
	g.Add(10)
	g.Add(3)
	g.Add(6)
	Convey("全中", t, func() {
		So(19, ShouldEqual, g.ScoreForFrame(1))
		So(28, ShouldEqual, g.GetScore())
		g.Add(10)
		g.Add(5)
		g.Add(4)
		So(56, ShouldEqual, g.GetScore())
	})

	g = NewGame()
	for i := 0; i < 12; i++ {
		g.Add(10)
	}
	Convey("", t, func() {
		So(300, ShouldEqual, g.GetScore())
	})
}
