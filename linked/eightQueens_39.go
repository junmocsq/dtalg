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

type bag struct {
	total   int
	currMax int
	res     [][]int
}

// 0-1 背包问题有很多变体，我这里介绍一种比较基础的。我们有一个背包，背包总的承载重量是 Wkg。
// 现在我们有 n 个物品，每个物品的重量不等，并且不可分割。我们现在期望选择几件物品，
// 装载到背包中。在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？
func NewBagZeroOne(total int) *bag {
	return &bag{total: total, res: make([][]int, 0)}
}
func (b *bag) BagZeroOne(arr []int) []int {
	b.bagZeroOne(arr, []int{}, 0, 0)
	return b.res[len(b.res)-1]
}

func (b *bag) bagZeroOne(arr, prefixArr []int, index int, currTotal int) {
	if currTotal > b.total {
		return
	}
	if currTotal == b.total {
		b.currMax = currTotal
		b.res = append(b.res, prefixArr)
		return
	}
	length := len(arr)
	if index == length {
		// 没有数据时 取最大
		if currTotal > b.currMax {
			b.currMax = currTotal
			b.res = append(b.res, prefixArr)
		}
		return
	}
	if currTotal+arr[index] > b.total {
		if currTotal > b.currMax {
			b.currMax = currTotal
			b.res = append(b.res, prefixArr)
		}
	} else {
		b.bagZeroOne(arr, append(prefixArr, arr[index]), index+1, currTotal+arr[index])
	}
	b.bagZeroOne(arr, prefixArr, index+1, currTotal)
}

type pattern struct {
	matched bool
	patt    []byte
}

// 正则表达式 匹配 *?
func NewPattern(pat string) *pattern {
	return &pattern{patt: []byte(pat)}
}

func (p *pattern) match(str string) bool {
	p.matched = false
	p.rmatch([]byte(str), 0, 0)
	return p.matched
}

func (p *pattern) rmatch(chars []byte, ti, pi int) {
	if p.matched {
		return
	}
	if pi == len(p.patt) {
		if ti == len(chars) {
			p.matched = true
		}
		return
	}
	if p.patt[pi] == '*' { // 匹配任意字符
		for k := 0; k < len(chars)-ti; k++ {
			p.rmatch(chars, ti+k, pi+1)
		}
	} else if p.patt[pi] == '?' { // 匹配0或一个
		p.rmatch(chars, ti, pi+1)
		p.rmatch(chars, ti+1, pi+1)
	} else if ti < len(chars) && p.patt[pi] == chars[ti] {
		p.rmatch(chars, ti+1, pi+1)
	}
}
