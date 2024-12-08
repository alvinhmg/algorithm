package ALGORITHM

import "sort"

/**
@author: alvin
@create: 2024-12-08 15:30
title: 三数之和
题解：
	1. 暴力法：三层循环，时间复杂度O(n^3)
	2. 双指针法：时间复杂度O(n^2)，空间复杂度O(1)
		1. 先对数组进行排序
		2. 定义两个指针，遍历排序后的数组，使用两个指针，一个指针指向排序后数组的头，一个指针指向数组的尾
		3.遍历数组，计算 sum = nums[i] + nums[left] + nums[right]
		4. 如果sum == target，返回下标
		5. 如果sum < target，left++
		6. 如果sum > target，right--
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序
	ret := make([][]int, 0) // 初始化结果数组
	for i :=0; i < len(nums)-2; i++ {  // 遍历数组，先固定一个数 nums[i]
		if nums[i] > 0 {
			break // 如果nums[i] > 0，说明后面的数都大于0，不可能相加等于0，直接退出循环
		}
		if i > 0 && nums[i] == nums[i-1] { // 如果当前数与前一个数相等，说明已经遍历过，跳过
			continue // 去重
		}
		left, right := i+1, len(nums)-1 // 
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})  // 找到一组解，加入结果数组
				for left < right && nums[left] == nums[left+1] { // 去重，跳过左指针重复的结果
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 去重，跳过有指针的结果
					right--
				}
				// 移动指针，继续寻找下一组解
				left++
				right--
			}
		}
	}
	return ret
}