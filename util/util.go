package util

import "strings"

func Snake2Camel(str string) string {
	strBytes := []byte(str)
	flag := true
	newBytes := make([]byte, 0, len(strBytes))
	for _, b := range strBytes {
		if b == '_' {
			flag = true
			continue
		}
		if flag == true {
			flag = false
			if b > 96 && b < 123 {
				b -= 32
			}
		}
		newBytes = append(newBytes, b)
	}
	return string(newBytes)
}

func TrimPrefix(tableName, prefix string) string {
	if prefix == "" {
		return tableName
	}
	if strings.HasPrefix(tableName, prefix) {
		return strings.Replace(tableName, prefix, "", 1)
	}
	return tableName
}
