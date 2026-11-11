package domain

// bestMove выбирает ход компьютера, перебирая возможные продолжения партии.
// При одинаковой оценке берётся первая свободная клетка.
func bestMove(board Board) (Board, bool) {
	bestScore := -100
	best := board
	found := false

	for row := range board {
		for col := range board[row] {
			if board[row][col] != Empty {
				continue
			}
			next := board
			next[row][col] = Computer
			score := minimax(next, false, 1)
			if !found || score > bestScore {
				bestScore = score
				best = next
				found = true
			}
		}
	}

	return best, found
}

func minimax(board Board, computerTurn bool, depth int) int {
	finished, winner := result(board)
	if finished {
		switch winner {
		case Computer:
			return 10 - depth
		case Human:
			return depth - 10
		default:
			return 0
		}
	}

	if computerTurn {
		best := -100
		for row := range board {
			for col := range board[row] {
				if board[row][col] == Empty {
					next := board
					next[row][col] = Computer
					score := minimax(next, false, depth+1)
					if score > best {
						best = score
					}
				}
			}
		}
		return best
	}

	best := 100
	for row := range board {
		for col := range board[row] {
			if board[row][col] == Empty {
				next := board
				next[row][col] = Human
				score := minimax(next, true, depth+1)
				if score < best {
					best = score
				}
			}
		}
	}
	return best
}

func result(board Board) (finished bool, winner int) {
	lines := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}
	for _, line := range lines {
		a := board[line[0]/3][line[0]%3]
		b := board[line[1]/3][line[1]%3]
		c := board[line[2]/3][line[2]%3]
		if a != Empty && a == b && b == c {
			return true, a
		}
	}
	for _, row := range board {
		for _, cell := range row {
			if cell == Empty {
				return false, Empty
			}
		}
	}
	return true, Empty
}
