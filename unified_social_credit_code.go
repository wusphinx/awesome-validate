/*
规则参照 https://zh.wikisource.org/zh-hans/GB_32100-2015_法人和其他组织统一社会信用代码编码规则

Note: 发现统一社会代码如：p91350213M0001Q7Q47 不符合编码规则，但是通过天眼查能查到该企业
https://www.tianyancha.com/search?key=91350213M0001Q7Q47
官方的说法是: 您好，官网数据都是由当地登记管理部门进行回传，建议您咨询发证机关。
http://www.cods.org.cn/portal/publish/questionList/questionList_1_1_questionDetail_3_1.html

*/
package validate

import (
	"fmt"
	"regexp"
)

var unifiedSocialCreditCodeRegexp *regexp.Regexp
var code2value map[string]int
var positionWeights []int

func init() {
	first := `(1|5|9|Y)`
	second := `(1|2|3|9)`
	thirdToEight := `(\d{6})`
	nineToEighteen := `([0-9A-Z^IOZSV]{10})`
	str := fmt.Sprintf(`^%s%s%s%s$`, first, second, thirdToEight, nineToEighteen)
	unifiedSocialCreditCodeRegexp = regexp.MustCompile(str)
	code2value = map[string]int{
		`0`: 0,
		`1`: 1,
		`2`: 2,
		`3`: 3,
		`4`: 4,
		`5`: 5,
		`6`: 6,
		`7`: 7,
		`8`: 8,
		`9`: 9,
		`A`: 10,
		`B`: 11,
		`C`: 12,
		`D`: 13,
		`E`: 14,
		`F`: 15,
		`G`: 16,
		`H`: 17,
		`J`: 18,
		`K`: 19,
		`L`: 20,
		`M`: 21,
		`N`: 22,
		`P`: 23,
		`Q`: 24,
		`R`: 25,
		`T`: 26,
		`U`: 27,
		`W`: 28,
		`X`: 29,
		`Y`: 30,
	}
	positionWeights = []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
}

func ValidateUnifiedSocialCreditCode(unifiedSocialCreditCode string) bool {
	if !unifiedSocialCreditCodeRegexp.MatchString(unifiedSocialCreditCode) {
		return false
	}

	var r int
	for index, i := range unifiedSocialCreditCode[:17] {
		c, ok := code2value[string(i)]
		if !ok {
			return false
		}
		r += c * positionWeights[index]
	}
	r = 31 - r%31
	if r == 31 {
		r = 0
	}
	if r == 30 {
		r = code2value[`Y`]
	}

	checkCode, ok := code2value[string(unifiedSocialCreditCode[17])]
	if !ok {
		return false
	}

	if r != checkCode {
		return false
	}

	return true
}
