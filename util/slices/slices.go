package slices

// IntersectInt64s 计算两个int64切片的交集
func IntersectInt64s(slice1, slice2 []int64) []int64 {
	if len(slice1) == 0 || len(slice2) == 0 {
		return []int64{}
	}

	// 创建映射以提高查找效率
	set := make(map[int64]bool, len(slice2))
	for _, id := range slice2 {
		set[id] = true
	}

	// 查找交集
	result := make([]int64, 0)
	for _, id := range slice1 {
		if set[id] {
			result = append(result, id)
		}
	}
	return result
}

// IntersectStrings 计算两个字符串切片的交集
func IntersectStrings(slice1, slice2 []string) []string {
	if len(slice1) == 0 || len(slice2) == 0 {
		return []string{}
	}

	// 创建映射以提高查找效率
	set := make(map[string]bool, len(slice2))
	for _, str := range slice2 {
		set[str] = true
	}

	// 查找交集
	result := make([]string, 0)
	for _, str := range slice1 {
		if set[str] {
			result = append(result, str)
		}
	}
	return result
}

// ContainsInt64 检查int64切片是否包含指定元素
func ContainsInt64(slice []int64, item int64) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsString 检查字符串切片是否包含指定元素
func ContainsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// UniqueInt64s 去除int64切片中的重复元素
func UniqueInt64s(slice []int64) []int64 {
	if len(slice) <= 1 {
		return slice
	}

	seen := make(map[int64]bool)
	result := make([]int64, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// UniqueStrings 去除字符串切片中的重复元素
func UniqueStrings(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
