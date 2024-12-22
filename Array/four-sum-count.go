package ALGORITHM

/**
@author: alvin
@create: 2024-12-22 21:30
title: 四数相加II
desc: 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
来源：力扣（LeetCode）

思路：
	1. 四数相加II：时间复杂度O(n^2)，空间复杂度O(n^2)
	2. 将四个数之和转为两个数之和，然后使用哈希表存储两个数之和的结果，然后遍历哈希表，找到两个数之和的结果为0的个数
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 定义一个哈希表，存储两个数之和的结果
	m := make(map[int]int) // key: 两个数之和，value: 两个数之和的个数
	// 定义一个变量，存储结果
	var res int
	// 遍历前二个数组
	for _, v1 := range nums1  {
		for _, v2 := range nums2 {
			m[v1+v2] ++ // 将两个数之和的结果存储到哈希表中
		}
	}
	// 遍历后二个数组
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			target := -(v3+v4)  // 这里是找 v1 + v2 + v3 + v4 = 0，所以是 -(v3+v4)
			if count, ok := m[target]; ok {
				res += count
			}
		}
	}
	return res // 返回结果
}