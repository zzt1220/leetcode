package _49_两个数组的交集

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	var res []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			//把值放在结果里面后，删除map中的值，防止重复值加入
			delete(m, v)
		}
	}
	return res
}
