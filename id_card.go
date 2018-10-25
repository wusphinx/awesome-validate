/*
规则参照 https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码

Note:这里实现了身份证的校验规则，包括出生日期的校验，行政区划代码存在不确定性，故而没有加入校验
*/

package validate

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var idCardRegexp *regexp.Regexp
var index2weight map[int]int

func init() {
	str := `^\d{17}[0-9X]$`
	idCardRegexp = regexp.MustCompile(str)
	index2weight = map[int]int{
		1:  1,
		2:  2,
		3:  4,
		4:  8,
		5:  5,
		6:  10,
		7:  9,
		8:  7,
		9:  3,
		10: 6,
		11: 1,
		12: 2,
		13: 4,
		14: 8,
		15: 5,
		16: 10,
		17: 9,
		18: 7,
	}
}

func ValidateIDCard(idCard string) bool {
	if !idCardRegexp.MatchString(idCard) {
		return false
	}
	if _, err := time.Parse("20060102", string(idCard[6:14])); err != nil {
		return false
	}

	var sum int
	for index, item := range idCard[:17] {
		value, _ := strconv.Atoi(string(item))
		sum += index2weight[18-index] * value
	}
	checkCode := (12 - sum%11) % 11
	if checkCode < 10 {
		return fmt.Sprintf("%d", checkCode) == string(idCard[17])
	}

	return string(idCard[17]) == "X"
}
