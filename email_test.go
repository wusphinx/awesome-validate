package validate_test

import (
	"testing"

	"github.com/wusphinx/awesome-validate"
)

func TestValidateEmailOK(t *testing.T) {
	// 网上搜的一些银行、外交部、教育部门的公开邮箱
	emails := []string{
		"fgjc@stats.gov.cn",
		"zfsswdc@stats.gov.cn",
		"info@stats.gov.cn",
		"rssjdc@stats.gov.cn",
		"12391@moe.edu.cn",
		"sunlili090901@163.com",
		"whjy1111@163.com",
		"xzfjyj@tonglu.gov.cn",
		"jjz83986003@163.com",
		"jyjjsgh9852@163.com",
		"lwp1982918@163.com",
		"zsedu@126.com",
		"service.la@bankofchina.com",
		"info@centralbank.cw",
		"admin@mfa.gov.cn",
		"Chinaemb_et@mfa.gov.cn",
	}

	for _, email := range emails {
		if ok := validate.ValidateEmail(email); !ok {
			t.Errorf("may be not correct email:%s ", email)
		}
	}
}

func TestValidateEmailFail(t *testing.T) {
	emails := []string{
		"23e23@",
		"@okkn",
		"&sdf@@",
		"no safsdf @11",
		"**&@com.11",
		"**&@银河",
		"😂hello@163.com",
		"hello@🎉.com",
		"",
	}

	for _, email := range emails {
		if ok := validate.ValidateEmail(email); ok {
			t.Errorf("should not correct email:%s ", email)
		}
	}
}
