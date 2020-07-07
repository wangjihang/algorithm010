package algorithm

// 北 东 南 西
func robotSim(commands []int, obstacles [][]int) int {
	var (
		x, y, direction, max int
		dx, dy               = []int{0, 1, 0, -1}, []int{1, 0, -1, 0} // 北 东 南 西
		obstacleM            = make(map[[2]int]struct{})
	)
	for _, vs := range obstacles {
		if len(vs) > 1 {
			obstacleM[[2]int{vs[0], vs[1]}] = struct{}{}
		}
	}

	for _, cmd := range commands {
		if cmd == -1 {
			direction = (direction + 1) % 4
		} else if cmd == -2 {
			direction = (direction + 3) % 4
		} else {
			for i := 0; i < cmd; i++ {
				nx, ny := x+dx[direction], y+dy[direction]
				if _, ok := obstacleM[[2]int{nx, ny}]; ok {
					break
				}
				x, y = nx, ny
				if ans := x*x + y*y; ans > max {
					max = ans
				}
			}
		}
	}
	return max
}
