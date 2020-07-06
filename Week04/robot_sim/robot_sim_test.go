package robot_sim

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRobotSim(t *testing.T) {
	Convey("TestRobotSim", t, func() {
		Convey("commands = [4,-1,3], obstacles = []", func() {
			So(RobotSim([]int{4, -1, 3}, nil), ShouldEqual, 25)
		})
	})
}
