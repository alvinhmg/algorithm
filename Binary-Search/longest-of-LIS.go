package ALGORITHM

/**
 * @author alvin
 * @date 2023/05/15
 300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
@leetcode: https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=binary-search
@思路：
	1、动态规划
	2、二分查找
	3、动态规划 + 二分查找
 */
 func longestOfLIS_Dynamic(nums []int) int {
	dp := make([]int, len(nums))
	// 初始化
	for i := range dp {
		dp[i] = 1 // 每个元素都可以单独成为子序列，长度为1
	}

	// 状态转移
	for i := 0; i < len(nums); i ++ {
		for j := 0; j < i; j ++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1) // 如果nums[i] > nums[j]，则可以将nums[i]加入到以nums[j]结尾的子序列中，长度加1
			}
		}
	}
	ret := 0
	for _, v := range dp {
		ret = max(ret, v) // 找到最长的子序列长度
	}
	return ret
}


// 二分查找  参考纸牌游戏 https://labuladong.online/algo/dynamic-programming/longest-increasing-subsequence/#%E4%BA%8C%E3%80%81%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE%E8%A7%A3%E6%B3%95
func longestOfLIS_Binary(nums []int) int {
	top := make([]int, len(nums)) // 牌堆
	piles := 0 // 牌堆数
	for i := 0; i < len(nums); i++ {
 		poker := nums[i] // 要处理的扑克牌
		// 二分查找插入位置, 查找左侧边界的二分查找
		left, right := 0, piles 
		for left < right {
			mid := left + (right - left)/2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1 
			} else {
				right = mid
			}
		}

		if left == piles { // 没有找到合适的牌堆，新建一堆
			piles++ 
		}
		top[left] = poker // 把这张牌放到牌堆顶
	}
	return piles // 牌堆数就是 LIS 长度
}

func longestOfLIS_DynamicAndBinary(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 动态规划数组
	dp := []int{}

	for _, num := range nums {
		// 二分查找 num 在 dp 中的位置
		left, right := 0, len(dp)
		for left < right {
			mid := left + (right-left)/2
			if dp[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 如果找到的位置等于 dp 长度，表示 num 是新的最大值
		if left < len(dp) {
			dp[left] = num
		} else {
			dp = append(dp, num)
		}
	}

	return len(dp)
}