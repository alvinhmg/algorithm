package ALGORITHM

import (
	"math"
)

/**
 * @author alvin
 * @date 2024/12/29
 * @title
 53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
@leetcode: URL_ADDRESS@leetcode: https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-interview-53

@solution:
	1. 分治法
	2. 动态规划
	3. 前缀和
*/

// 1、分治法
// 思路：1、先把数组平均分成左右两部分，2、最大子数组可能在左半部分，也可能在右半部分，也可能跨越左右两部分。3、递归求解左右两部分的最大子数组和。4、比较三种情况的最大值。
func maxSumArray(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) int {
	if left > right { // 递归终止条件
		return math.MinInt64 // 返回最小值
	}

	mid := left + (right - left)/2 // 定义中间位置
	// 分别计算左右半部分的最大子数组和
	leftSum := helper(nums, left, mid-1)
	rightSum := helper(nums, mid+1, right)

	leftMaxSuffixSum, rightMaxProfixSum  := 0, 0
	total := 0
	// 计算左半部分最大后缀和
	for i := mid - 1; i >= left; i-- {
		total += nums[i]
		if total > leftMaxSuffixSum {
			leftMaxSuffixSum = total
		}
	}

	// 计算右半部分最大前缀和
	total = 0
	for i := mid + 1; i <= right; i++ {
		total += nums[i]
		if total > rightMaxProfixSum {
			rightMaxProfixSum = total
		}
	}
	// 计算跨越左右两部分的最大子数组之和
	crossSum := leftMaxSuffixSum + rightMaxProfixSum + nums[mid]

	// 返回三种情况的最大值
	return max(crossSum, max(leftSum, rightSum))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2、动态规划
// 思路：
// 1、定义一个数组 dp，dp[i] 表示以 nums[i] 结尾的最大子数组和。
// 2、状态转移方程：dp[i] = max(dp[i-1] + nums[i], nums[i])。
// 3、返回 dp 数组中的最大值。
func maxSumArray_Dynamic(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i]) // 比较当前元素和当前元素与当前和，取最大值作为当前值
		maxSum = max(maxSum, currentSum) // 更新最大值
	}
	return maxSum
}

// 3、前缀和
// 思路：
// 1、定义一个变量 sum，表示当前的前缀和。
// 2、定义一个变量 maxSum，表示最大子数组和。
// 3、遍历数组，对于每个元素，更新 sum 和 maxSum。
// 4、返回 maxSum。

func maxSumArray_PrefixSum(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0

	// 计算前缀和
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1] // 前缀和等于前一个前缀和加上当前元素
	}
	minPrefix := 0 // 定义最小前缀和
	maxSum :=math.MinInt64
	// 遍历数组，找到最大子数组和
	for i := 1; i <= len(nums); i++ {
		maxSum = max(maxSum, prefixSum[i]-minPrefix) // 比较当前前缀和减去最小前缀和，取最大值作为当前值
		minPrefix = min(minPrefix, prefixSum[i]) // 更新最小前缀和
	}
	return maxSum
}	


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}