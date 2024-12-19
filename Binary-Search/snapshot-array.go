package ALGORITHM

import "sort"

/**
@author: alvin
@create: 2024-12-19 22:29
@title: 快照数组
@desc: 实现支持下列接口的「快照数组」- SnapshotArray：
SnapshotArray(int length) - 初始化一个与指定长度相等的 类数组 的数据结构。初始时，每个元素都等于 0。
void set(index, val) - 会将指定索引 index 处的元素设置为 val。
int snap() - 获取该数组的快照，并返回快照的编号 snap_id（快照号是调用 snap() 的总次数减去 1）。
int get(index, snap_id) - 根据指定的 snap_id 选择快照，并返回该快照指定索引 index 的值。
@leetcode: https://leetcode.cn/problems/snapshot-array/description/?envType=study-plan-v2&envId=binary-search

@solution: 二分查找的实际使用
*/

// 快照数组
type SnapshotArray struct{
	SnapID int  // 当前的快照ID
	Data []map[int]int // 每一个索引存储一个map[sanp_id]val 
}

// 构造函数
func SnapshotArrayConstructor(length int) SnapshotArray {
	data := make([]map[int]int, length) // 初始化数组
	for i :=0; i < length; i ++ {
		data[i] = map[int]int{} // 初始化每一个索引map
	}
	return SnapshotArray{
		SnapID: 0,
		Data: data,
	}
}


// Set 设置指定索引的值
func (s *SnapshotArray) Set(index int, val int)  {
	s.Data[index][s.SnapID] = val // 设置指定索引index 对应的快照ID的val
}

// Snap 创建快照并返回当前快照的ID
func(s *SnapshotArray) Snap() int {
	s.SnapID++  // 增加快照ID
	return s.SnapID - 1
}

// Get 获取指定索引index 在指定快照ID snap_id 下的值
func (s *SnapshotArray) Get(snap_id int, index int ) int {
	records:= s.Data[index]
	if val, ok := records[snap_id]; ok {
		return val
	}
	// 如果没有找到，返回最近一次快照的值，使用二分查找
	keys := make([]int, 0, len(records))
	for k := range records {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	pos := sort.Search(len(keys), func(i int) bool {
		return keys[i]> snap_id
	})
	if pos == 0 {
		return 0
	}
	return records[keys[pos-1]] // 返回最近一次快照的值
}