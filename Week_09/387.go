package Week_09

//387. 字符串中的第一个唯一字符
func FirstUniqChar(s string) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	if length < 2 {
		return 0
	}
	hashMap := make(map[int32]int)
	for _, v := range s {
		if v1, ok := hashMap[v]; ok == true {
			hashMap[v] = v1 + 1
			continue
		}
		hashMap[v] = 1
	}
	for i, v := range s {
		if v1, ok := hashMap[v]; ok == true && v1 == 1 {
			return i
		}
	}
	return -1
}

func FirstUniqChar2(s string) int {
	arr := [26]int{}
	for _, char := range s {
		arr[char-'a']++
	}

	for i, v := range s {
		if arr[v-'a'] == 1 {
			return i
		}
	}
	return -1
}


