package config

import "syscall"

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultUILanguage = kernel32.NewProc("GetUserDefaultUILanguage")
)

// Windows LANGID primary language IDs
const (
	langPrimaryGerman = 0x07
)

// GetSystemLanguage returns "de" or "en" based on the Windows UI language.
// Falls back to "en" for any non-German locale.
func GetSystemLanguage() string {
	ret, _, _ := procGetUserDefaultUILanguage.Call()
	if uint16(ret)&0x3FF == langPrimaryGerman {
		return "de"
	}
	return "en"
}
