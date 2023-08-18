package utils

import (
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func CalculateEPaySign(mapData map[string]string, secret string) string {
	// sort keys
	keys := lo.Keys(mapData)
	sort.Strings(keys)

	combinedData := ""
	for _, k := range keys {
		if k == "sign" || k == "sign_type" || lo.IsEmpty(mapData[k]) {
			continue
		}

		combinedData += fmt.Sprintf("%s=%s&", k, mapData[k])
	}

	combinedData = strings.TrimSuffix(combinedData, "&")

	return MD5String(combinedData + secret)
}

func CheckEPaySign(mapData map[string]string, secret string, sign string) bool {
	return CalculateEPaySign(mapData, secret) == sign
}
