package go_try

type Number interface {
	~int | ~float64 | ~float32 | ~int8 | ~int32 | ~int64
}

func GetValIndex[V any](val []V, target V, compare func(v1, v2 V) bool) int {
	for i, v := range val {
		if compare(v, target) {
			return i
		}
	}
	return -1
}

// SortSince todo 实现相关的排序算法
func SortSince[V any](val []V, compare func(v1, v2 V) int) {
	if len(val) < 2 {
		return
	}
	for i := 0; i < len(val)-1; i++ {
		for j := i + 1; j < len(val); j++ {
			if compare(val[i], val[j]) > 0 {
				val[i], val[j] = val[j], val[i]
			}
		}
	}
}

func GetMapFromSlice[Val any, K comparable](val []Val, getKey func(val Val) K) map[K]Val {
	ans := make(map[K]Val)
	for _, v := range val {
		ans[getKey(v)] = v
	}
	return ans
}

func DoIt[V any](param []V, behave func(s V)) {
	for i, v := range param {
		behave(v)
		if i == len(param)-1 {
			println()
		}
	}
}

func GetKeyFromMap[Val any, K comparable](valMap map[K]Val) []K {
	ans := make([]K, 0, len(valMap))
	for k := range valMap {
		ans = append(ans, k)
	}
	return ans
}

func Max[K Number](k1, k2 K) K {
	if k1 > k2 {
		return k1
	} else {
		return k2
	}
}

func Min[K Number](k1, k2 K) K {
	if k1 < k2 {
		return k1
	} else {
		return k2
	}
}
