package config

import "syscall"

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultUILanguage = kernel32.NewProc("GetUserDefaultUILanguage")
)

// Windows LANGID primary language IDs and common full LANGIDs.
const (
	langPrimaryGerman   = 0x07
	langPrimaryChinese  = 0x04
	langPrimaryJapanese = 0x11

	// Chinese — traditional locales (sublanguage in high word)
	langChineseTraditionalTW = 0x0404 // Taiwan
	langChineseTraditionalHK = 0x0C04 // Hong Kong SAR
	langChineseTraditionalMO = 0x1404 // Macau SAR
)

// GetSystemLanguage returns a canonical locale tag from the Windows UI language.
func GetSystemLanguage() string {
	ret, _, _ := procGetUserDefaultUILanguage.Call()
	id := uint16(ret)
	primary := id & 0x3FF
	switch primary {
	case langPrimaryGerman:
		return "de_DE"
	case langPrimaryJapanese:
		return "ja_JP"
	case langPrimaryChinese:
		switch id {
		case langChineseTraditionalTW, langChineseTraditionalHK, langChineseTraditionalMO:
			return "zh_TW"
		default:
			return "zh_CN"
		}
	default:
		return "en_US"
	}
}
