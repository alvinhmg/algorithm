package ALGORITHM

import "sort"

/**
@author: alvin
@create: 2024-12-17 21:30
title: 基于时间的键值存储
desc: 设计一个基于时间的键值数据结构，该结构可以在不同时间戳存储对应同一个键的多个值，并针对特定时间戳检索键对应的值。
leetcode: https://leetcode.cn/problems/time-based-key-value-store/description/?envType=study-plan-v2&envId=binary-search
题解：
	1. 基于时间的键值存储：时间复杂度O(logn)，空间复杂度O(n)
	2. 题目要求：
		1. 存储键值对
		2. 查找键值对
		3. 查找键值对的时间戳
		4. 查找键值对的时间戳的前一个时间戳
*/

type TimeValue struct {
	value     string
	timestamp int
}

type TimeMap struct {
	stroe map[string][]TimeValue // 存储键值对
}

func Constructor() TimeMap {
	return TimeMap{
		stroe: make(map[string][]TimeValue, 0), // 初始化存储
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.stroe[key] = append(m.stroe[key], TimeValue{ // 先根据key获取到value, 该value是一个切片, 然后将新的value添加到切片中
		value:     value,
		timestamp: timestamp,
	})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	values, exist := m.stroe[key] // 先根据key获取到value, 该value是一个切片
	if !exist || len(values) == 0 {
		return "" // 如果不存在, 则返回空字符串
	}
	// 二分查找
	idx := sort.Search(len(values), func(i int) bool { // 二分查找的条件是: 时间戳是否大于等于timestamp
		return values[i].timestamp >= timestamp
	})
	if idx == 0 { // 如果idx等于0, 则说明没有找到
		return ""
	}
	return values[idx-1].value // 如果idx不等于0, 则说明找到了, 则返回idx-1的value, 因为idx是大于等于timestamp的
}
