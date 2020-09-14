package p6

// 保龄球计分小程序
type game struct {
	itsCurrentFrame   int
	firstThrowInFrame bool
	scorer            *scorer
}

func NewGame() *game {
	return &game{
		firstThrowInFrame: true,
		itsCurrentFrame:   1,
		scorer:            NewScorer(),
	}
}

func (g *game) GetScore() int {
	return g.ScoreForFrame(g.itsCurrentFrame)
}

func (g *game) Add(pins int) {
	g.scorer.addThrow(pins)
	g.adjustCurrentFrame(pins)
}

func (g *game) ScoreForFrame(theFrame int) int {

	return g.scorer.scoreForFrame(theFrame)
}

func (g *game) GetCurrentFrame() int {

	return g.itsCurrentFrame
}

func (g *game) adjustCurrentFrame(pins int) {

	if g.lastBallInFrame(pins) {
		g.firstThrowInFrame = true
		g.advanceFrame()
	} else {
		g.firstThrowInFrame = false
	}
}

// 全中判断
func (g *game) strike(pins int) bool {
	return g.firstThrowInFrame && pins == 10
}

// 本轮最后一次
func (g *game) lastBallInFrame(pins int) bool {
	return g.strike(pins) || !g.firstThrowInFrame
}

func (g *game) advanceFrame() {
	if g.itsCurrentFrame < 10 {
		g.itsCurrentFrame++
	}
}

func (g *game) adjustFrameForStrike(pins int) bool {
	if pins == 10 {
		g.advanceFrame()
		return true
	}
	return false
}
