package validate_test

import (
	"testing"

	"github.com/wusphinx/awesome-validate"
)

func TestValidateUrlOK(t *testing.T) {
	cards := []string{
		"http://www.google.com",
		"http://www.cloudflare.com",
		"http://www.github.com",
		"http://www.v2ex.com",
		"https://www.netflix.com",
		"http://www.facebook.com",
	}
	for _, card := range cards {
		if ok := validate.ValidateUrl(card); !ok {
			t.Errorf("may be not correct url:%s ", card)
		}
	}
}

func TestValidateUrlFail(t *testing.T) {
	cards := []string{
		" http://www.google.com",
		"Http://www.cloudflare.com",
		"Zttp://www.github.com",
		"http:www.v2ex.com",
		"://www.netflix.com",
		"http:///www.facebook.com",
		"",
	}
	for _, card := range cards {
		if ok := validate.ValidateUrl(card); ok {
			t.Errorf("should not correct url:%s ", card)
		}
	}
}
