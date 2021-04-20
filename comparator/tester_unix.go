// +build !windows

package comparator

func GetExeName(name string) string {
	return "./" + name
}
