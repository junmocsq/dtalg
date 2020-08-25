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
		tmp := s[i : i+lp]
		h := 0
		if lastHash == 0 {
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
