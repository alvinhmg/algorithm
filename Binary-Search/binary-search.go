package ALGORITHM

/**
@author: alvin
@create: 2024-12-08 15:30
title: 二分查找
题解：
	1. 二分查找：时间复杂度O(logn)，空间复杂度O(n)

二分查找的前提：
	1. 目标函数单调性（单调递增或者递减）
	2. 存在上下界（bounded）
	3. 能够通过索引访问（index accessible）

二分查找的模板：
	1. 初始化left和right
	2. while left <= right
	3. 计算mid
	4. 判断mid是否等于target
	5. 判断mid是否小于target
	6. 判断mid是否大于target
	7. 返回-1

*/

// @param target 目标值
// @return 目标值的下标
// @return -1 没有找到
// 左闭右闭区间版本  [left, right]

func BinarySearch_basic1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 左闭右开区间版本 【left， right）
func BinarySearch_basic2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

// 查找左边界的二分查找	
func binarySearch_findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // 搜索区间变为[mid+1, right]
		} else if nums[mid] > target {
			right = mid - 1 // 搜索区间变为[left, mid-1]
		} else if nums[mid] == target {
			result = mid  // 记录位置，继续向左搜索
			right = mid - 1 // !! 别返回，锁定左侧区间继续搜索
		}
	}
	return result
}

// 查找右边界的二分查找

func binarySearch_findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right{
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			result = mid // 记录位置，继续向右搜索
			left = mid + 1 //!! 别返回，锁定右侧区间继续搜索
		}
	}
	return result
}
