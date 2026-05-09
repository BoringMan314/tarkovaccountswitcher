@echo off
setlocal EnableExtensions
cd /d "%~dp0"

if not exist "v2\wails.json" (
  echo [build_win10] FAIL: missing v2\wails.json ^(run from repo root^)
  goto :end_fail
)

set "EXE_NAME=Tarkov Account Switcher.exe"
set "ROOT_EXE=%~dp0%EXE_NAME%"
set "OUT_DIR=%~dp0v2\build\bin"
set "OUT_EXE=%OUT_DIR%\%EXE_NAME%"

echo [build_win10] Wails Win10/amd64: %ROOT_EXE% ^(from v2\build\bin^)

taskkill /F /IM "%EXE_NAME%" /T >nul 2>&1
if exist "%ROOT_EXE%" del /f /q "%ROOT_EXE%" 2>nul

where go >nul 2>&1
if errorlevel 1 (
  echo [build_win10] FAIL: go not in PATH ^(install Go 1.23+^)
  goto :end_fail
)

where wails >nul 2>&1
if errorlevel 1 (
  echo [build_win10] FAIL: wails not in PATH
  echo [build_win10] install: go install github.com/wailsapp/wails/v2/cmd/wails@latest
  goto :end_fail
)

pushd "v2" >nul
if errorlevel 1 (
  echo [build_win10] FAIL: cannot cd to v2
  goto :end_fail
)

echo [build_win10] using:
go version
wails version

go mod tidy
if errorlevel 1 (
  echo [build_win10] FAIL: go mod tidy
  popd >nul
  goto :end_fail
)

go run sync_version.go
if errorlevel 1 (
  echo [build_win10] FAIL: sync_version.go
  popd >nul
  goto :end_fail
)

wails build -platform windows/amd64
set "BUILD_RC=%ERRORLEVEL%"
popd >nul
if not "%BUILD_RC%"=="0" (
  echo [build_win10] FAIL: wails build
  goto :end_fail
)

if not exist "%OUT_EXE%" (
  echo [build_win10] FAIL: missing %OUT_EXE%
  goto :end_fail
)

copy /Y "%OUT_EXE%" "%ROOT_EXE%" >nul
if errorlevel 1 (
  echo [build_win10] FAIL: copy to repo root
  goto :end_fail
)

if not exist "%ROOT_EXE%" (
  echo [build_win10] FAIL: missing %ROOT_EXE%
  goto :end_fail
)

echo [build_win10] OK: %ROOT_EXE%
goto :end_ok

:end_fail
if /i "%~1"=="nopause" exit /b 1
echo.
pause
exit /b 1

:end_ok
if /i "%~1"=="nopause" exit /b 0
echo.
pause
exit /b 0
