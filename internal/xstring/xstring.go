package xstring

import "strings"

func ToCamelWithStartUpper(str string) string {
	r := ""
	strs := strings.Split(str, "_")
	for _, item := range strs {
		r += strings.ToUpper(item[0:1])
		r += item[1:]
	}

	return r
}
