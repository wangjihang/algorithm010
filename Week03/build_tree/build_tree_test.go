package build_tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildTree(t *testing.T) {
	Convey("TestBuildTree", t, func() {
		BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	})
}
