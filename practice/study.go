package practice

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
func letterCombinations(digits string) []string {
	arr := make(map[string][]string)
	arr["2"] = []string{"a", "b", "c"}
	arr["3"] = []string{"d", "e", "f"}
	arr["4"] = []string{"g", "h", "i"}
	arr["5"] = []string{"j", "k", "l"}
	arr["6"] = []string{"m", "n", "o"}
	arr["7"] = []string{"p", "q", "r", "s"}
	arr["8"] = []string{"t", "u", "v"}
	arr["9"] = []string{"w", "x", "y", "z"}

	ans := []string{""}
	for _, v := range digits {
		lastArr := make([]string, len(ans))
		copy(lastArr, ans)
		ans = nil
		for _, _v := range arr[string(v)] {
			for _, pre := range lastArr {
				ans = append(ans, pre+_v)
			}
		}
	}
	return ans
}
