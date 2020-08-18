package p5

type HashFunc interface {
	Hash(string, int) int
}

type horner struct {
}

func NewHorner() *horner {
	return &horner{}
}
func (h *horner) Hash(ele string, size int) int {
	ret := 0
	for _, v := range ele {
		ret = (ret << 5) + int(v)
	}
	return ret % size
}

type charapter struct {
}

func NewChapter() *charapter {
	return &charapter{}
}

func (c *charapter) Hash(ele string, size int) int {
	arr := []byte(ele)
	return int(arr[0]) % size
}
