package alipay

import (
	"encoding/base64"
	"sort"
	"strings"
	"github.com/smartwalle/alipay/encoding"
	"crypto"
)

func verifySign(param map[string]string, publicKey []byte) (ok bool, err error) {
	sign, err := base64.StdEncoding.DecodeString(param["sign"])
	if err != nil {
		return false, err
	}

	var keys = make([]string, 0, 0)
	for key, value := range param {
		if key == "sign" || key == "sign_type" {
			continue
		}
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param[key])
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")

	err = encoding.VerifyPKCS1v15([]byte(s), sign, publicKey, crypto.SHA256)

	if err != nil {
		return false, err
	}
	return true, nil
}
