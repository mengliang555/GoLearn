package go_try

import (
	"testing"
)

func FuzzMaxAndInt(f *testing.F) {
	f.Add(1, 2)
	f.Fuzz(func(t *testing.T, k1, k2 int) {
		max := Max[int](k1, k2)
		min := Min[int](k1, k2)
		if max < min {
			t.Error("code error")
		}
	})
}

func FuzzMaxAndFloat(f *testing.F) {
	f.Add(float32(1.0), float32(2.1))
	f.Fuzz(func(t *testing.T, k1, k2 float32) {
		max := Max[float32](k1, k2)
		min := Min[float32](k1, k2)
		if max < min {
			t.Error("code error")
		}
	})
}
