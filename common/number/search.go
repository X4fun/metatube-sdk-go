package number

import (
	"fmt"
	"regexp"
)

func BuildMovieSearchKeyword(s string) string {

	s = Trim(s)
	// 使用正则表达式提取 "abc" 和 "123" 部分
	re := regexp.MustCompile(`([a-zA-Z]+)[-_]?(\d+)[a-zA-Z]*`)
	matches := re.FindStringSubmatch(s)

	if len(matches) < 3 {
		// 如果没有匹配成功，返回原字符串或其他提示
		return ""
		//return s // 返回原始字符串
	}

	// 拼接成所需的格式
	return fmt.Sprintf("%s-%s", matches[1], matches[2])
}
