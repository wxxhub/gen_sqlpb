package gen

import (
	"fmt"
	"github.com/wxxhub/gen_sqlpb/internal/xstring"
	"html/template"
	"strings"
)

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{
		"ItemIndex":      ItemIndex,
		"AddNote":        AddNote,
		"StringCamel":    StringCamel,
		"StringLowFirst": StringLowFirst,
		"StringToLow":    strings.ToLower,
		"StringToUpper":  strings.ToUpper,
		"JoinAnd":        JoinAnd,
		"ToCamelJoinAnd": ToCamelJoinAnd,
		"RpcLine":        RpcLine,
		"StringJoin":     StringJoin,
	}
}

func GetTemplateFuncList() template.FuncMap {
	return funcMap
}

func AddTemplateFunc(addFuncMap template.FuncMap) {
	for k, v := range addFuncMap {
		funcMap[k] = v
	}
}

// ItemIndex 加一操作
func ItemIndex(index int) int {
	return index + 1
}

// AddNote 注释格式
func AddNote(comment string) string {
	if len(comment) > 0 {
		return "// " + comment
	}
	return ""
}

// StringCamel 将字符串转为驼峰形式
func StringCamel(str string) string {
	return xstring.ToCamelWithStartUpper(str)
}

// StringLowFirst 获取字符串第一个字符的小写
func StringLowFirst(str string) string {
	return strings.ToLower(string(str[0]))
}

func StringJoin(strs ...string) string {
	r := ""
	for _, item := range strs {
		r += item
	}
	return r
}

// JoinAnd 用And将字符串拼接
func JoinAnd(strs []string) string {
	return strings.Join(strs, "And")
}

// ToCamelJoinAnd 用And将字符串拼接
func ToCamelJoinAnd(strs []string) string {
	t := make([]string, len(strs))
	for index := range strs {
		t[index] = xstring.ToCamelWithStartUpper(strs[index])
	}
	return strings.Join(t, "And")
}

func RpcLine(funName string) string {
	return fmt.Sprintf("rpc %s(%sReq) returns (%sRes)", funName, funName, funName)
}
