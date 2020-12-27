package demo2

func Run(nums []int, target int) []int {
	//1. 两数之和
	//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	//
	//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

	//双指针
	var j int
	var res []int
	for i := 0; i < len(nums); i++ {
		j = 1
		for i != j && j < len(nums) {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
			j++
		}
	}
	return res
}

func Run2(nums []int, target int) []int {
	//哈希表
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
