package epay

import (
	"epay-cli/pkg/utils"
)

// GenerateParams 生成加签参数
func GenerateParams(params map[string]string, secret string) map[string]string {
	params["sign"] = utils.CalculateEPaySign(params, secret)
	params["sign_type"] = "MD5"
	return params
}
