package visitor

import (
	"fmt"
)

type visitor interface {
	visit(f resourceFile)
}

// 提取数据
type extractor struct {
}

func NewExtractor() visitor {
	return &extractor{}
}

func (v *extractor) visit(f resourceFile) {
	if _, ok := f.(*pptFile); ok {
		fmt.Println("extract ppt", f.getFilePath())
	}

	if _, ok := f.(*wordFile); ok {
		fmt.Println("extract word", f.getFilePath())
	}

	if _, ok := f.(*execlFile); ok {
		fmt.Println("extract execl", f.getFilePath())
	}
}

// 压缩数据
type compressor struct {
}

func NewCompressor() visitor {
	return &compressor{}
}

func (v *compressor) visit(f resourceFile) {
	if _, ok := f.(*pptFile); ok {
		fmt.Println("compress ppt", f.getFilePath())
	}

	if _, ok := f.(*wordFile); ok {
		fmt.Println("compress word", f.getFilePath())
	}

	if _, ok := f.(*execlFile); ok {
		fmt.Println("compress execl", f.getFilePath())
	}
}

// A访问者
type A struct {
}

func NewA() visitor {
	return &A{}
}

func (v *A) visit(f resourceFile) {
	if _, ok := f.(*pptFile); ok {
		fmt.Println("A ppt", f.getFilePath())
	}

	if _, ok := f.(*wordFile); ok {
		fmt.Println("A word", f.getFilePath())
	}

	if _, ok := f.(*execlFile); ok {
		fmt.Println("A execl", f.getFilePath())
	}
}
