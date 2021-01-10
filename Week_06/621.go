package Week_06

import "math"

//
//给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
//
//然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//
//你需要计算完成所有任务所需要的 最短时间 。

func leastInterval(tasks []byte, n int) int {
	var minTime int
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}

	nextValid := make([]int, 0, len(cnt))
	rest := make([]int, 0, len(cnt))
	for _, c := range cnt {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}

	for range tasks {
		minTime++
		minNextValid := math.MaxInt64
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return minTime
}
