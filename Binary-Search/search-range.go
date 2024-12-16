package ALGORITHM

/**
@author: alvin
@create: 2024-12-16 22:30
@description: 在指定数组中查找给定元素的起始位置和结束位置
@param: nums  []int  数组   target  int  目标元素
@return: []int  起始位置和结束位置
@note:
	1. 二分查找：时间复杂度O(logn)，空间复杂度O(1)
	2. 二分查找的变种，查找指定元素的起始位置和结束位置，需要考虑元素不存在的情况
	3. 分别利用二分查找左右边界查找指定元素的起始位置和结束位置
@leetcode: https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=binary-search
*/
func searchRange(nums []int, target int) []int {
	return []int{binarySearch_findLeftBound(nums, target), binarySearch_findRightBound(nums, target)}
}