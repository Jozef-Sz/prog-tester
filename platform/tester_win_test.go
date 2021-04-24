// +build windows
package platform

import "testing"

func TestStrCmp(t *testing.T) {
	a := "jedno\nVelkeSpojeneSlovo\n"
	b := "jedno\r\nVelkeSpojeneSlovo\r\n"

	equality := StrCmp(a, b)

	if equality == false {
		t.Error("Two strings should equal.")
	}
}
