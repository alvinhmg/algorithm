package ALGORITHM

/**
@author: alvin
@create: 2024-12-08 15:30
title: 二分查找-寻找比目标字母大的最小字母
题解：
	1. 二分查找：时间复杂度O(logn)，空间复杂度O(1)
	2. 二分查找的变种，有序字母数组中查找比指定字母大对应的最小字母，需要考虑字母不存在的情况
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 循环结束后，left指向的是比target大的最小字母
	if left < len(letters){
		return letters[left] // 说明存在，返回下一个字母
	} else {
		return letters[0] // 不存在返回首字母
	}
}
