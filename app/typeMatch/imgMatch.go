package typeMatch

import (
	"regexp"
)

func GetImgSlice(content string, matchCount int, types []string) (resultSet []string) {
	resultSet = []string{}
	typesLength := len(types)
	var compile *regexp.Regexp
	for i := 0; i < typesLength ; i++ {
		regexStr := "src=\".*\\." + types[i]
		compile = regexp.MustCompile(regexStr)
		typeResultSet := compile.FindAllString(content, matchCount)
		resultSet = append(resultSet, typeResultSet...)
	}
	for index, item := range(resultSet)  {
		resultSet[index] = item[5:]
	}
	return resultSet
}
