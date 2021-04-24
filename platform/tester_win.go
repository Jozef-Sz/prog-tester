// +build windows

package platform

import "strings"

func GetExeName(name string) string {
	return name + ".exe"
}

func StrCmp(a, b string) bool {
	return strings.ReplaceAll(a, "\r\n", "\n") == strings.ReplaceAll(b, "\r\n", "\n")
}
