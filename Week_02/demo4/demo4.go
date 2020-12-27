package demo4

import (
	"sort"
)

//49. 字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

//方法一：排序
func Run(strs []string) [][]string {
	hashTables := map[string][]string{}
	var res [][]string
	for _, v := range strs {
		noSort := v
		b_v := []byte(v)
		sort.Slice(b_v, func(i, j int) bool {
			return b_v[i] < b_v[j]
		})
		sortStr := string(b_v)
		hashTables[sortStr] = append(hashTables[sortStr], noSort)
	}
	for _, v := range hashTables {
		res = append(res, v)
	}
	return res
}

//计数
func Run2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
