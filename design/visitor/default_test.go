package visitor

import "testing"

func TestTool(t *testing.T) {
	toolApplication()
	/*
		=== RUN   TestTool
		extract ppt a.ppt
		compress ppt a.ppt
		A ppt a.ppt

		extract word b.word
		compress word b.word
		A word b.word

		extract execl c.execl
		compress execl c.execl
		A execl c.execl

		extract ppt d.ppt
		compress ppt d.ppt
		A ppt d.ppt

		extract execl e.execl
		compress execl e.execl
		A execl e.execl

		--- PASS: TestTool (0.00s)
	*/
}
