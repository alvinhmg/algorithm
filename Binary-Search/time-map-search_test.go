package ALGORITHM

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	t.Log(timeMap.Get("foo", 1))
	t.Log(timeMap.Get("foo", 3))
	timeMap.Set("foo", "bar2", 4)
	t.Log(timeMap.Get("foo", 4))
	t.Log(timeMap.Get("foo", 5))
}
