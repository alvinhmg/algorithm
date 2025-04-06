package ALGORITHM

import "sort"

/**
@author: alvin
@time: 2024-12-22 22:30
@desc:

*/

func firstBadVersion(n int) int {
	return sort.Search(n, func(i int) bool { return isBadVersion(i) })
}

func isBadVersion(_ int) bool {
	return true
}
