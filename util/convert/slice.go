package convert

// 切片

// UniqueSlice 切片去重
func UniqueSlice[K comparable](languages []K) []K {
	result := make([]K, 0, len(languages))
	temp := map[K]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Remove 移除切片满足条件的元素
func Remove(sl []interface{}, f func(v1 interface{}) bool) []interface{} {
	for k, v := range sl {
		if f(v) {
			sl[k] = sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			return sl
		}
	}
	return sl
}

// RemoveSlice 移除切片中指定元素
func RemoveSlice[K comparable](src []K, sub K) []K {
	for k, v := range src {
		if v == sub {
			copy(src[k:], src[k+1:])
			return src[:len(src)-1]
		}
	}
	return src
}

// DifferenceSlice 比较两个切片，返回他们的差集
func DifferenceSlice[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, item := range s1 {
		m[item] = true
	}

	var diff []T
	for _, item := range s2 {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// InSlice 元素是否存在于切片中
func InSlice[K comparable](slice []K, key K) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}
