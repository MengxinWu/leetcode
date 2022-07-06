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

/*
一群孩子站成一排，每一个孩子有自己的评分。现在需要给这些孩子发糖果，规则是如果一个孩子的评分比自己身旁的一个孩子要高，
那么这个孩子就必须得到比身旁孩子更多的糖果；所有孩子至少要有一个糖果。求解最少需要多少个糖果。

输入是一个数组，表示孩子的评分。输出是最少糖果的数量。
*/

func candy(ratings []int) int {
	var num = make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		num[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			num[i+1] = num[i] + 1
		}
	}
	for j := len(ratings) - 1; j > 0; j-- {
		if ratings[j-1] > ratings[j] {
			num[j-1] = getMax(num[j]+1, num[j-1])
		}
	}
	total := 0
	for _, cnt := range num {
		total += cnt
	}
	return total
}

func getMax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，
能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	num := 0
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			num = 1
		}
		if num >= n {
			return true
		}
		return false
	}
	for i := 0; i < len(flowerbed); i++ {
		// 种过了
		if flowerbed[i] == 1 {
			i++
			continue
		}
		// 头
		if i == 0 {
			if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				num++
			}
			continue
		}
		// 尾
		if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				num++
			}
			continue
		}
		// 中间
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			num++
		}
	}
	if num >= n {
		return true
	}
	return false
}
