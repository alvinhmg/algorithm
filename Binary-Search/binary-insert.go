package ALGORITHM

/**
@author: alvin
@create: 2024-12-08 15:30
title: 二分查找-插入指定位置
题解：
	1. 二分查找：时间复杂度O(logn)，空间复杂度O(1)
	2. 二分查找的变种，查找指定元素的插入位置
	3. 二分查找的变种，查找指定元素的插入位置，需要考虑元素不存在的情况
	4. 二分查找的变种，查找指定元素的插入位置，需要考虑元素存在的情况
*/

func binary_insert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 // 缩小左边界
		} else {
			right = mid - 1 // 缩小右边界
		}
	}
	return left
}
