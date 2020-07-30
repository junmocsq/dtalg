package linked

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleInitSort(t *testing.T) {
	s := SingleInit()
	s.AddToTail(1)
	s.AddToTail(2)
	s.AddToTail(3)
	s.AddToTail(4)
	Convey("升序或无序", t, func() {
		So(s.CheckSort(), ShouldEqual, 1)
		s.AddToTail(3)
		So(s.CheckSort(), ShouldEqual, 0)
	})

	s = SingleInit()
	s.AddToTail(4)
	s.AddToTail(3)
	s.AddToTail(2)
	s.AddToTail(1)
	Convey("降序或类型不同", t, func() {
		So(s.CheckSort(), ShouldEqual, -1)
		s.AddToTail("李斯")
		So(s.CheckSort(), ShouldEqual, 0)
	})
}
