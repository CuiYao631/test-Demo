package main

func main() {
	//给定一个字符串,找出其中不含有重复字符的 最长子串的长度
	str := "abcabcbb"
	//思路: 用一个map来存储每个字符最后出现的位置,然后遍历字符串,如果当前字符在map中存在,则更新start的值,如果不存在,则将当前字符和位置存入map中,并更新max的值
	//start表示当前子串的起始位置,max表示最长子串的长度
	start, max := 0, 0
	//用一个map来存储每个字符最后出现的位置
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		//如果当前字符在map中存在,则更新start的值
		if _, ok := m[str[i]]; ok {
			start = maxInt(start, m[str[i]]+1)
		}
		//将当前字符和位置存入map中
		m[str[i]] = i
		//更新max的值
		max = maxInt(max, i-start+1)
	}
	//返回最长子串的长度

}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
