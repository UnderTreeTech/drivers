package drivers

import "strings"

// ParseDSNAddr 解析DSN中的IP+PORT，该方法仅适用于@(IP+PORT)/格式的DSN，如非此类型的请自定义解析方法
func ParseDSNAddr(dsn string) (addr string) {
	atIdx := strings.Index(dsn, "@")
	if atIdx == -1 {
		return
	}
	slot := strings.Trim(dsn[atIdx:], "@")
	addrs := strings.Split(slot, "/")
	if len(addrs) == 0 {
		return
	}
	addr = addrs[0]
	return
}
