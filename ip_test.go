package validate_test

import (
	"testing"

	"github.com/wusphinx/awesome-validate"
)

func TestValidateIPOK(t *testing.T) {
	ips := []string{
		"183.18.124.0",
		"69.63.184.14",
		"8.7.198.45",
		"208.101.60.87",
		"198.41.214.162",
		"1.1.1.1",
		"8.8.8.8",
		"2001:da8:8000:1:202:120:2:100",
		"2001:da8:8000:1:202:120:2:101",
		"2001:da8:202:10::36",
		"2001:da8:208:10::6",
		"2001:da8:202:10::37",
		"2001:dc7:1000::1",
		"2001:4860:4860::8888",
	}

	for _, ip := range ips {
		if ok := validate.ValidateIP(ip); !ok {
			t.Errorf("may be not correct ip:%s ", ip)
		}
	}
}

func TestValidateIPFail(t *testing.T) {
	ips := []string{
		"1183.18.124.0",
		"-1.63.184.14",
		"8.7.1",
		"208.101.60.87.1",
		",198.41.214.162",
		"a1.1.1.1",
		"8.8.8.8:da8",
		"2001:da8:80001:1:202:120:2:100",
		"",
	}

	for _, ip := range ips {
		if ok := validate.ValidateIP(ip); ok {
			t.Errorf("should not correct ip:%s ", ip)
		}
	}
}
