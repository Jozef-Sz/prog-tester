// +build !windows

package tester

func GetExeName(name string) string {
	return "./" + name
}
