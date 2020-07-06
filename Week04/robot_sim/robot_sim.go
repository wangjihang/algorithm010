package robot_sim

// 北 东 南 西 [1, 2, 3, 4]
// 时间：O(N+K)
// 空间：O(k) 存储obstacle数组
func RobotSim(commands []int, obstacles [][]int) int {
	var (
		dx                   = [4]int{0, 1, 0, -1} // 北 东 南 西
		dy                   = [4]int{1, 0, -1, 0} // 北 东 南 西
		x, y, direction, max int

		obstacleM = make(map[[2]int]struct{})
	)

	for _, vs := range obstacles {
		if len(vs) >= 2 {
			obstacleM[[2]int{vs[0], vs[1]}] = struct{}{}
		}
	}

	for _, cmd := range commands {
		if cmd == -2 {
			direction = (direction + 3) % 4 // 左边 + 3 [上一个方向]
		} else if cmd == -1 {
			direction = (direction + 1) % 4 // 右边 + 1 [下一个方向]
		} else { // 方向不变  (cx, cy) + (dx, dy)
			for i := 0; i < cmd; i++ {
				nx, ny := x+dx[direction], y+dy[direction]
				if _, ok := obstacleM[[2]int{nx, ny}]; ok { // 前面有障碍物，此action到此结束
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
