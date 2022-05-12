package go_try

func GetMapValue[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	~int | ~float64 | ~float32 | ~int8 | ~int32 | ~int64
}

func AddValue[K Number](k1, k2 K) K {
	return k1 + k2
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
