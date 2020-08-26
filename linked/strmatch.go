package linked

// Brute Force 中文叫作暴力匹配算法，也叫朴素匹配算法
type BruteForce struct {
}

func NewBF() *BruteForce {
	return &BruteForce{}
}

func (b BruteForce) Match(str, pattern string) int {
	p := []byte(pattern)
	s := []byte(str)
	end := len(s) - len(p)
	for i := 0; i <= end; i++ {
		flag := true
		for k := 0; k < len(p); k++ {
			if p[k] != s[i+k] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

type RabinKarp struct {
	powArr []int
}

// RK 算法的思路是这样的：我们通过哈希算法对主串中的 n-m+1 个子串分别求哈希值，然后逐个与模式串的哈希值比较大小。如果某个子串的哈希值与模式串相等，那就说明对应的子串和模式串匹配了
func NewRk() *RabinKarp {
	return &RabinKarp{
		powArr: make([]int, 0),
	}
}

func (r *RabinKarp) Match(str, pattern string) int {
	s := []byte(str)
	p := []byte(pattern)
	lp := len(p)
	hp := r.hash(p)
	lastHash := 0
	for i := 0; i <= len(s)-len(p); i++ {
		h := 0
		if lastHash == 0 {
			tmp := s[i : i+lp]
			h = r.hash(tmp)
		} else {
			// 通过上一个子串的hash值快速计算下一个子串的hash值
			h := (lastHash-int(s[i-1])*r.pow(lp-1))*26 + int(s[i+lp-1])
			lastHash = h
		}
		if h == hp {
			return i
		}
	}
	return -1
}

func (r *RabinKarp) hash(s []byte) int {
	sum := int(s[0] - 'a')
	for i := 1; i < len(s); i++ {
		sum = (sum + int(s[i]-'a')) * 26
	}
	return sum
}

func (r *RabinKarp) pow(i int) int {
	for n := len(r.powArr); n <= i; n++ {
		if n == 0 {
			r.powArr = append(r.powArr, 1)
		} else {
			r.powArr = append(r.powArr, r.powArr[n-1]*26)
		}
	}
	return r.powArr[i]
}

// BM 算法包含两部分，分别是坏字符规则（bad character rule）和好后缀规则（good suffix shift）。
type BoyerMoore struct {
	ascii  []int
	suffix []int
	prefix []bool
}

func NewBM() *BoyerMoore {
	return &BoyerMoore{}
}

// https://time.geekbang.org/column/article/71525

func (b *BoyerMoore) genBc(pattern string) {
	b.ascii = make([]int, 256)
	for i := 1; i < 256; i++ {
		b.ascii[i-1] = -1
	}
	for k, v := range pattern {
		index := int(v)
		b.ascii[index] = k
	}
}

func (b *BoyerMoore) Match(str, pattern string) int {
	b.genBc(pattern)
	b.genGs(pattern)
	s := []byte(str)
	p := []byte(pattern)
	lp := len(pattern)
	ls := len(s)
	for i := lp - 1; i < ls; {
		xi := -1
		si := -1
		for k := lp - 1; k >= 0; lp-- {
			// 坏字符匹配
			diff := lp - 1 - k
			if p[k] != s[i-diff] {
				xi = k
				si = b.ascii[s[i-lp+1+k]]
				break
			}
		}
		if xi >= 0 {
			x := xi - si
			// 好后缀匹配
			y := b.MoveGs(xi, lp)
			if x > y {
				i += x
			} else {
				i += y
			}
		} else {
			return i - (lp - 1)
		}
	}

	return -1
}

// 后缀数组
func (b *BoyerMoore) genGs(pattern string) {
	l := len(pattern)
	b.suffix = make([]int, l)
	b.prefix = make([]bool, l)
	for i := 0; i < l; i++ {
		b.suffix[i] = -1
	}
	p := []byte(pattern)
	for i := 0; i < l-1; i++ {
		j := i
		k := 0
		for k >= 0 && p[j] == p[l-1-k] {
			j--
			k++
			b.suffix[k] = j + 1
		}
		if j == -1 {
			b.prefix[k] = true
		}
	}
}

// 好后缀
func (b *BoyerMoore) MoveGs(j, m int) int {
	k := m - 1 - j
	if b.suffix[k] != -1 {
		return j - b.suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if b.prefix[m-r] {
			return r
		}
	}
	return m
}
