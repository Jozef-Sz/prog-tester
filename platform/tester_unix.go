// +build !windows

package platform

func GetExeName(name string) string {
	return "./" + name
}

func StrCmp(a, b string) bool {
	return a == b
}
