package linked

// https://github.com/wangzheng0822/algo/blob/master/go/17_skiplist/skiplist.go
// 跳表

type skipListNode struct {
	val      interface{}
	score    int
	level    int
	forwards []*skipListNode
}

func skiptest() {


}
