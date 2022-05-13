package go_try

import (
	"fmt"
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

func FuzzMaxAndInt1(f *testing.F) {
	f.Add(10, 2)
	f.Fuzz(func(t *testing.T, k1, k2 int) {
		max := Max[int](k1, k2)
		min := Min[int](k1, k2)
		if max < min {
			t.Error("code error")
		}
	})
}

func FuzzMaxAndFloat1(f *testing.F) {
	f.Add(float32(3.0), float32(2.1))
	f.Fuzz(func(t *testing.T, k1, k2 float32) {
		max := Max[float32](k1, k2)
		min := Min[float32](k1, k2)
		if max < min {
			t.Error("code error")
		}
	})
}

type testStruct struct {
	name string
	age  int32
}

func TestGetMapFromSlice(t *testing.T) {
	val := []testStruct{{"hello", 12}, {"world", 13}}
	ans := GetMapFromSlice[testStruct, string](val, func(v testStruct) string { return v.name })
	for key, val := range ans {
		fmt.Printf("key:%s val:%v", key, val)
	}
}

func TestGetKeyFromMap(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6}
	ans := GetMapFromSlice[int, int](target, func(val int) int { return val * val })
	fi := GetKeyFromMap[int, int](ans)
	fmt.Println(fi)
}

func TestDoIt(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6}
	DoIt[int](target, func(v int) { println(v) })
}

func TestGetValIndex(t *testing.T) {
	target := []int{1 + 1, 2 + 4, 3 + 6, 4 + 8, 5 + 12, 6 + 15}
	print(GetValIndex[int](target, 6, func(v1, v2 int) bool { return v1 == v2 }))
}

func TestGetValIndex1(t *testing.T) {
	target := []int{1 + 1, 2 + 4, 3 + 6, 4 + 8, 5 + 12, 6 + 15}
	print(GetValIndex[int](target, 100, func(v1, v2 int) bool { return v1 == v2 }))
}

func TestSort(t *testing.T) {
	target := []int{1 + 11, 2 + 1, 3 + 2, 4 + 8, 5 + 12, 6 + 2}
	DoIt[int](target, func(v int) { fmt.Printf("%d\t", v) })
	SortSince[int](target, func(v1, v2 int) int { return v1 - v2 })
	DoIt[int](target, func(v int) { fmt.Printf("%d\t", v) })
}

func TestSort2(t *testing.T) {
	target := []int{1 + 11}
	DoIt[int](target, func(v int) { fmt.Printf("%d\t", v) })
	SortSince[int](target, func(v1, v2 int) int { return v1 - v2 })
	DoIt[int](target, func(v int) { fmt.Printf("%d\t", v) })
}
