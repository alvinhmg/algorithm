package ALGORITHM
/**
@author: alvin
@time: 2024-12-22 22:30
@desc:	寻找K个最接近的元素
思路：
	二分查找，找到窗口的左边界，然后返回窗口内的元素
leetcode: https://leetcode.cn/problems/find-k-closest-elements/?envType=study-plan-v2&envId=binary-search
*/

func findKClosestElements(nums []int, k int, x int) []int {
	left, right := 0, len(nums)-k // 定义左右边界，注意右边界的最大值为 len(nums)-k
	for left < right {
		mid := left + (right-left)>>1 // 定义中间位置
		if x-nums[mid] > nums[mid+k]-x { // 比较中点和后边边界的距离
			// 如果中点距离比后边边界距离大，说明答案在右边，更新左边界
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left : left+k] // 返回窗口内的元素
}
