package ALGORITHM

/**
@author: alvin
@create: 2024-12-22 22:30
title: 搜索旋转排序数组,整数数组 nums 按升序排列，数组中的值 互不相同 。
思路：
	1. 从旋转点开始分成两个分别有序的数组，然后分别使用二分查找
	2. 旋转点的特点：旋转点的左边和右边都是有序的
 leetcode: https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=binary-search
 */
func rotateSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		
		if nums[left] <= nums[mid] { // 判断左边数组是否有序
			if nums[left] <= target && target < nums[mid] { // 如果左边数组有序，则判断target是否在左边数组中
				right = mid -1
			} else {
				left = mid + 1
			}
		} else { // 否则右边数组是否有序
			if nums[mid] < target && target <= nums[right] { // 如果右边数组有序，则判断target是否在右边数组中
				left = mid + 1 
			} else {
				right = mid - 1 
			}

		}
	}
	return -1
}
