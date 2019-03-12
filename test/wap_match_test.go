package test

import (
	"log"
	"net/url"
	"regexp"
	"testing"
)

func TestWapMatch(t *testing.T) {

	str := `        var data = {"requestType":"SafePay","fromAppUrlScheme":"alipays","dataString":"h5_route_token=\"RZ25j33MBqtxbjX65zCa17UODGbzqgmobilecashierRZ25\"&is_h5_route=\"true\""};`

	reg := regexp.MustCompile(`(?U)var data = (.*?);`)

	result := reg.FindStringSubmatch(str)
	log.Println(result)

	s := url.QueryEscape(result[1])

	log.Println(s)

	if s != "%7B%26quot%3Bserver_param%26quot%3B%3A%26quot%3BemlkPTE5O25kcHQ9MTZlNjtjYz15%26quot%3B%7D" {

		t.Error("result not match")
	}

	l := "alipay://alipayclient/?" + s

	log.Println(l)

}
