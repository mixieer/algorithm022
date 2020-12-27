package demo1

func Run(s string, t string) bool {
	//242. 有效的字母异位词
	//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
	if len(s) != len(t) {
		return false
	}

	tab1 := make(map[int32]int)
	for _, v := range s {
		tab1[v-'a']++
	}
	for _, v := range t {
		key := v - 'a'
		tab1[key]--
		if tab1[key] < 0 {
			return false
		}
	}
	return true
}
