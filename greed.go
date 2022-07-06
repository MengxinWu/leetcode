package main

import (
	"sort"
)

/*
有一群孩子和一堆饼干，每个孩子有一个饥饿度，每个饼干都有一个大小。每个孩子只能吃最多一个饼干，
且只有饼干的大小大于孩子的饥饿度时，这个孩子才能吃饱。求解最多有多少孩子可以吃饱。

输入两个数组，分别代表孩子的饥饿度和饼干的大小。输出最多有多少孩子可以吃饱的数量。
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	num := 0
	pos := 0
	for idxHunger := 0; idxHunger < len(g); idxHunger++ {
		for idxCookie := pos; idxCookie < len(s); idxCookie++ {
			pos = idxCookie
			if s[idxCookie] >= g[idxHunger] {
				num++
				pos++
				break
			}
		}
	}
	return num
}
