package service

import "luban/utils"

type SysDict struct {
	Id    int64
	Key   string
	Value string
}

func GetValueByKey(key string) string {
	m := utils.QuerySingle("SELECT value FROM sys_dict WHERE key = ?", key)
	if len(m) == 1 {
		return m["value"].(string)
	}
	return ""
}
