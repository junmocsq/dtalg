package linked

type eightQueen struct {
	res [][8]int
}

func NewEightQueen() *eightQueen {
	return &eightQueen{res: make([][8]int, 0)}
}

func (e *eightQueen) EightQueens() [][8]int {
	res := [8]int{}
	e.eightQueens(res, 0)
	return e.res
}

func (e *eightQueen) eightQueens(res [8]int, row int) {
	if row == 8 {
		e.res = append(e.res, res)
		return
	}
	for column := 0; column < 8; column++ {
		if e.checkNode(res, row, column) {
			res[row] = column
			e.eightQueens(res, row+1)
		}
	}
}
func (e *eightQueen) checkNode(result [8]int, row, column int) bool {
	ok := true
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			ok = false
			return ok
		}
	}
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if result[i] == j {
			ok = false
			return ok
		}
	}
	for i, j := row-1, column+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if result[i] == j {
			ok = false
			return ok
		}
	}
	return ok
}
