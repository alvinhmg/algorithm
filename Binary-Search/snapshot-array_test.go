package ALGORITHM

import "testing"

func TestSnapshotArray(t *testing.T) {
	snap := SnapshotArrayConstructor(3)
	snap.Set(0, 5)
	snap.Snap()
	snap.Set(0, 6)
	snap.Get(0, 0)
}
