package _6_最小覆盖子串

import "math"

/*
1、滑动窗口
*/
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	minLen := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(s); r++ {
		if ori[s[r]] > 0 {
			cnt[s[r]]++

			for l <= r {
				if _, ok := ori[s[l]]; !ok {
					l++
					continue
				}
				if check() {
					if r-l+1 < minLen {
						minLen = r - l + 1
						ansL, ansR = l, r+1
					}
					cnt[s[l]] -= 1
					l++
				} else {
					break
				}
			}
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
