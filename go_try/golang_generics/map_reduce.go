package go_try

type MapWithFilter[T comparable, V any] struct {
	val map[T]V
}

func NewMapWithFilter[T comparable, V any](val map[T]V) *MapWithFilter[T, V] {
	return &MapWithFilter[T, V]{val: val}
}

func (m *MapWithFilter[K, V]) GetValSet() []V {
	ans := make([]V, 0, len(m.val))
	for _, v := range m.val {
		ans = append(ans, v)
	}
	return ans
}

func (m *MapWithFilter[K, V]) GetKeySet() []K {
	ans := make([]K, 0, len(m.val))
	for k := range m.val {
		ans = append(ans, k)
	}
	return ans
}

// for filter

func (m *MapWithFilter[K, V]) FilterByKeyMethod(judge func(key K) bool) *MapWithFilter[K, V] {
	for k := range m.val {
		if !judge(k) {
			delete(m.val, k)
		}
	}
	return m
}

func (m *MapWithFilter[K, V]) FilterByValMethod(judge func(value V) bool) *MapWithFilter[K, V] {
	for k, v := range m.val {
		if !judge(v) {
			delete(m.val, k)
		}
	}
	return m
}
