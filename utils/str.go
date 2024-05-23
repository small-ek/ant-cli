package utils

import "strings"

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
		return "//" + comment
	}
	return ""

}
