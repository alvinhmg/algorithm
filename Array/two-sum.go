package ALGORITHM

import (
	"sort"
)

/**
@author: alvin
@create: 2024-12-08 15:30
title: 两数之和
题解：
	1. 暴力法：两层循环，时间复杂度O(n^2)
	2. 双指针：时间复杂度O(nlogn)，空间复杂度O(1)
		1. 先对数组进行排序
		2. 定义两个指针，分别指向数组的头和尾
		3. 遍历数组，计算sum = nums[left] + nums[right]
		4. 如果sum == target，返回下标
	3. 哈希表：时间复杂度O(1)，空间复杂度O(n)
		1. 遍历数组，将数组元素作为key，下标作为value存入map中
		2. 遍历数组，计算another = target - nums[i]
		3. 判断another是否在map中，若在，返回下标，若不在，将当前元素存入map中
*/

// 暴力法
func twoSum_1(nums []int , target int) []int {
	for i := 0; i < len(nums); i ++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 双指针
func twoSum_2 (nums []int, target int ) []int {
	// 先对数组进行排序
	sort.Ints(nums)
	// 调用go自带的快速排序方法
	
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int {left, right}		
		} else if sum < target {
			left ++
		} else {
			right--
		}
	}
	return nil
}

// 哈希表
func twoSum_3(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
