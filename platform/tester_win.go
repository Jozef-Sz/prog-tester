// +build windows

package platform

func GetExeName(name string) string {
	return name + ".exe"
}
