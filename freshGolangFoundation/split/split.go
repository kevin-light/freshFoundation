package split

import "strings"

// Split 将s按照sep进行分割，返回一个字符串的切片
func Split(s, sep string)(ret []string)  {
	idx := strings.Index(s,sep)
	for idx > -1{
		ret = append(ret,s[:idx])
		s = s[idx+len(sep):]		// 中文测试
		//s = s[idx+1:]				// a b c测试
		idx = strings.Index(s,sep)
	}
	ret = append(ret,s)
	return
}
