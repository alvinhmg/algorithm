package ALGORITHM

import (
	"sort"
)

/**
@author: alvin
@create: 2024-12-08 15:30
title: 四数之和
题解：
	1. 回溯法：时间复杂度O(n^4)，空间复杂度O(n)
*/

func fourSum_backtrack(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums) // 排序
	backtrack(nums, target, []int{}, 0, &ret)
	return ret
}

func backtrack(nums []int, target int, tempList []int, start int, ret *[][]int) {
	if len(tempList) > 4 { // 剪枝条件1： 超过4个元素直接返回
		return
	}
	if target == 0 && len(tempList) == 4 { // 剪枝条件2： 找到一组解，加入结果数组
		combination := make([]int, len(tempList))
		copy(combination, tempList) // 复制切片，避免引用传递
		*ret = append(*ret, combination) 
		return
	}

	for i := start; i < len(nums); i++ { // 遍历数组，从start开始
		if i > start && nums[i] == nums[i-1] { // 去重，跳过重复的结果
			continue
		}
		if len(tempList) == 0 && nums[i] > target {// 剪枝条件3： 如果当前数大于target，说明后面的数都大于target，不可能相加等于target，直接退出循环
			break
		}
		if len(tempList) == 0 && nums[i]+nums[len(nums)-1]*3 < target { // 剪枝条件4： 如果当前数加上后面的3个数都小于target，说明当前数加上后面的3个数都小于target，不可能相加等于target，直接退出循环
			continue
		}

		tempList = append(tempList, nums[i]) // 添加当前数到临时列表
		backtrack(nums, target-nums[i], tempList, i+1, ret) // 递归，继续寻找下一组解
		tempList = tempList[:len(tempList)-1] // 回溯，移除当前数，继续寻找下一组解
	}
}


/**
@author: alvin
@create: 2024-12-08 15:30
title: 四数之和
题解：
	1. 将问题分解为若干个子问题，将问题转化为Nsum问题，即将四个数之后先转为twoSum(a, threeSum(A)),再将threeSum(A)转为twoSum(b, twoSum(B))问题
	2. 再使用双指针解决twoSum()：时间复杂度O(n^3)，空间复杂度O(1)
		1. 先对数组进行排序
		2.采用分治法将四数之和转为Nsum问题
*/

func fourSum_fenzhi(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	findNsum(nums, target, 4, []int{},  &ret)
	return ret
}

func findNsum(nums []int, target int, n int, tempList []int, ret *[][]int) {
	if n < 2 || len(nums) < n { // 剪枝条件1： 如果n小于2，或者数组长度小于n，直接返回
		return 
	}

	// two-sum case
	if n == 2 {
		l, r := 0, len(nums) -1
		for l < r {
			if nums[l] + nums[r] == target {
				*ret = append(*ret, append(append([]int{}, tempList...), nums[l], nums[r])) // 添加当前数到临时列表
				l ++
				r -- 
				for l < r && nums[l] == nums[l-1] { // 去重，跳过重复的结果
					l ++
				}
				for l < r && nums[r] == nums[r+1] { // 去重，跳过重复的结果
					r--
				}
			} else if nums[l] + nums[r] < target {
				l ++
			} else {
				r --
			}
		}
	} else {
		for i := 0 ; i < len(nums)-1; i ++ {
			if i == 0 || (i > 0 && nums[i] == nums[i-1]) {
				findNsum(nums[i+1:], target-nums[i], n-1, append(tempList, nums[i]), ret) // 递归，继续寻找下一组解
			}
		}
	}
}
