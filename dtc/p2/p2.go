package p2

// N^3
func maxSubsequenceSumV1(arr []int) (s, e, sum int) {
	length := len(arr)

	if length == 0 {
		return
	} else {
		sum = arr[1]
		e = 1
	}
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			thisSum := 0
			for k := i; k < j; k++ {
				thisSum += arr[k]
			}
			if thisSum > sum {
				sum = thisSum
				s = i
				e = j
			}
		}
	}
	return
}

// N^2
func maxSubsequenceSumV2(arr []int) (s, e, sum int) {
	length := len(arr)

	if length == 0 {
		return
	} else {
		sum = arr[1]
		e = 1
	}
	for i := 0; i < length; i++ {
		thisSum := 0
		for j := i; j < length; j++ {
			thisSum += arr[j]
			if thisSum > sum {
				sum = thisSum
				s = i
				e = j
			}
		}
	}
	return
}

// N*logN
func maxSubsequenceSumV3(arr []int, start, end int) (sum int) {
	if start >= end {
		return 0
	}
	if start+1 == end {
		return arr[start]
	}
	mid := (start + end) / 2
	suml := maxSubsequenceSumV3(arr, start, mid)
	sumr := maxSubsequenceSumV3(arr, mid+1, end)
	ls, lmaxs := 0, 0
	for i := mid; i >= start; i-- {
		ls += arr[i]
		if ls > lmaxs {
			lmaxs = ls
		}
	}
	rs, rmaxs := 0, 0
	for i := mid + 1; i < end; i++ {
		rs += arr[i]
		if rs > rmaxs {
			rmaxs = rs
		}
	}
	return max3(suml, sumr, rmaxs+lmaxs)
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func maxSubsequenceSumV4(arr []int) int {
	length := len(arr)

	thisSum, maxSum := 0, 0
	for i := 0; i < length; i++ {
		thisSum += arr[i]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
	return maxSum
}
