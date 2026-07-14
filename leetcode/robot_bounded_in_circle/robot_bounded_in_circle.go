package robot_bounded_in_circle

const (
	UP int = iota
	LEFT
	DOWN
	RIGHT
)

func isRobotBounded(instructions string) bool {
	x, y, direction := 0, 0, 0

	for i := range instructions {
		if instructions[i] == 'L' {
			direction = (direction + 1) % 4
			continue
		} else if instructions[i] == 'R' {
			direction = (direction + 3) % 4
			continue
		}

		switch direction {
		case LEFT:
			x--
		case RIGHT:
			x++
		case UP:
			y++
		case DOWN:
			y--
		}
	}

	return x == 0 && y == 0 || direction != UP
}
