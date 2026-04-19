package config

import (
	"os"
	"syscall"
	"unsafe"
)

const registryKey = `Software\Microsoft\Windows\CurrentVersion\Run`
const appName = "TarkovAccountSwitcher"

var (
	advapi32        = syscall.NewLazyDLL("advapi32.dll")
	regOpenKeyExW   = advapi32.NewProc("RegOpenKeyExW")
	regSetValueExW  = advapi32.NewProc("RegSetValueExW")
	regDeleteValueW = advapi32.NewProc("RegDeleteValueW")
	regCloseKey     = advapi32.NewProc("RegCloseKey")
)

const (
	hkeyCurrentUser = 0x80000001
	keyWrite        = 0x20006
	regSZ           = 1
)

// ApplyAutoStart sets or removes the Windows autostart registry entry
func ApplyAutoStart(enabled bool) error {
	if enabled {
		return setAutoStartRegistry()
	}
	return removeAutoStartRegistry()
}

func setAutoStartRegistry() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	keyPath, _ := syscall.UTF16PtrFromString(registryKey)
	valueName, _ := syscall.UTF16PtrFromString(appName)
	valueData, _ := syscall.UTF16FromString(`"` + exePath + `"`)

	var hKey uintptr
	ret, _, _ := regOpenKeyExW.Call(hkeyCurrentUser, uintptr(unsafe.Pointer(keyPath)), 0, keyWrite, uintptr(unsafe.Pointer(&hKey)))
	if ret != 0 {
		return syscall.Errno(ret)
	}
	defer regCloseKey.Call(hKey)

	dataBytes := len(valueData) * 2 // UTF-16 = 2 bytes per char
	ret, _, _ = regSetValueExW.Call(hKey, uintptr(unsafe.Pointer(valueName)), 0, regSZ, uintptr(unsafe.Pointer(&valueData[0])), uintptr(dataBytes))
	if ret != 0 {
		return syscall.Errno(ret)
	}

	return nil
}

func removeAutoStartRegistry() error {
	keyPath, _ := syscall.UTF16PtrFromString(registryKey)
	valueName, _ := syscall.UTF16PtrFromString(appName)

	var hKey uintptr
	ret, _, _ := regOpenKeyExW.Call(hkeyCurrentUser, uintptr(unsafe.Pointer(keyPath)), 0, keyWrite, uintptr(unsafe.Pointer(&hKey)))
	if ret != 0 {
		return nil // Key doesn't exist, nothing to remove
	}
	defer regCloseKey.Call(hKey)

	regDeleteValueW.Call(hKey, uintptr(unsafe.Pointer(valueName)))
	return nil
}
