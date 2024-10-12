package utils

import (
	"strings"
)

// ToCamelCase 驼峰转换
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

// Many2Many 多对多表名
func Many2Many(table1, table2 string) string {
	return table1 + "_" + table2
}

func GetTag(Required int) string {
	if Required == 1 {
		return `binding:"required"`
	}
	return ""

}
func GetComment(comment string) string {
	if comment != "" {
		return "//" + RemoveNewlines(comment)
	}
	return ""

}
func RemoveNewlines(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\r\n", ""), "\n", "")
}

// ToCamelCaseLower 驼峰转换，开头小写
func ToCamelCaseLower(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(parts[i])
		} else {
			parts[i] = strings.Title(parts[i])
		}
	}
	return strings.Join(parts, "")
}

// ToKebabCase 将下划线小写转换成连字符命名
func ToKebabCase(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}
