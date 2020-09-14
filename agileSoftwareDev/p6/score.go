package p6

type scorer struct {
	itsThrows       [21]int
	itsCurrentThrow int
	ball            int
}

func NewScorer() *scorer {
	return &scorer{}
}

func (s *scorer) addThrow(pins int) {
	s.itsThrows[s.itsCurrentThrow] = pins
	s.itsCurrentThrow++
}

func (s *scorer) scoreForFrame(theFrame int) int {
	score := 0
	s.ball = 0
	for currentFrame := 0; currentFrame < theFrame; currentFrame++ {
		if s.strike() {
			score += 10 + s.nextTwoBallsForStrike()
			s.ball++
		} else if s.spare() {
			score += 10 + s.nextBallForSpare()
			s.ball += 2
		} else {
			score += s.twoBallsInFrame()
			s.ball += 2
		}
	}
	return score
}

// 保龄球全中判断
func (s *scorer) strike() bool {
	return s.itsThrows[s.ball] == 10
}

func (s *scorer) nextTwoBallsForStrike() int {
	return s.itsThrows[s.ball+1] + s.itsThrows[s.ball+2]
}

// 保龄球补中
func (s *scorer) spare() bool {
	return s.itsThrows[s.ball]+s.itsThrows[s.ball+1] == 10
}

func (s *scorer) nextBallForSpare() int {
	return s.itsThrows[s.ball+2]
}

func (s *scorer) twoBallsInFrame() int {
	return s.itsThrows[s.ball] + s.itsThrows[s.ball+1]
}
