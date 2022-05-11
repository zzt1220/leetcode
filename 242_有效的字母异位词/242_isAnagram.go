package _42_有效的字母异位词

/*
1、数组其实就是一个简单哈希表
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	exists := make([]int, 26)
	for i := 0; i < len(s); i++ {
		exists[s[i]-'a']++
		exists[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if exists[i] != 0 {
			return false
		}
	}
	return true
}
