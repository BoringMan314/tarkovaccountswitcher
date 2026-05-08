package i18n

import (
	"strings"
	"sync"
)

// Canonical locale tags stored in settings and sent to the frontend.
const (
	LangEN_US = "en_US"
	LangDE_DE = "de_DE"
	LangZH_TW = "zh_TW"
	LangZH_CN = "zh_CN"
	LangJA_JP = "ja_JP"
)

var (
	currentLanguage = LangEN_US
	languageMutex   sync.RWMutex
)

// Translation keys
const (
	// Tabs
	TabAccounts = "tabAccounts"
	TabAdd      = "tabAdd"
	TabSettings = "tabSettings"

	// Add Account Tab
	AddAccountTitle    = "addAccountTitle"
	LabelAccountName   = "labelAccountName"
	PlaceholderAccName = "placeholderAccountName"
	LabelEmail         = "labelEmail"
	PlaceholderEmail   = "placeholderEmail"
	AddAccountHelp     = "addAccountHelp"
	BtnAddAccount      = "btnAddAccount"

	// Account List
	EmptyStateTitle    = "emptyStateTitle"
	EmptyStateSubtitle = "emptyStateSubtitle"
	StatusAutoLogin    = "statusAutoLogin"
	StatusLoginReq     = "statusLoginRequired"
	BtnSwitch          = "btnSwitch"
	BtnDelete          = "btnDelete"
	ConfirmDelete      = "confirmDelete"

	// Settings Tab
	SettingsTitle       = "settingsTitle"
	LabelLanguage       = "labelLanguage"
	LabelLauncherPath   = "labelLauncherPath"
	PlaceholderLauncher = "placeholderLauncherPath"
	BtnBrowse           = "btnBrowse"
	BtnSave             = "btnSave"
	LabelStreamerMode   = "labelStreamerMode"
	StreamerModeHelp    = "streamerModeHelp"

	// Theme
	LabelTheme = "labelTheme"

	// Autostart
	LabelAutoStart = "labelAutoStart"
	AutoStartHelp  = "autoStartHelp"

	// Quit
	BtnQuit = "btnQuit"

	// Switch Result Messages (used by backend)
	SwitchAutoLogin   = "switchAutoLogin"
	SwitchManualLogin = "switchManualLogin"

	// Tray Menu
	TrayOpen = "trayOpen"
	TrayQuit = "trayQuit"

	// Status Messages
	StatusFillFields      = "statusFillFields"
	StatusAccountAdded    = "statusAccountAdded"
	StatusAccountDeleted  = "statusAccountDeleted"
	StatusDeleteError     = "statusDeleteError"
	StatusLauncherRestart = "statusLauncherRestarting"
	StatusAutoLoginActive = "statusAutoLoginActive"
	StatusManualLogin     = "statusManualLogin"
	StatusError           = "statusError"
	StatusPathSaved       = "statusPathSaved"
	StatusSaveError       = "statusSaveError"
	StatusEnterPath       = "statusEnterPath"
	StatusLanguageSaved   = "statusLanguageSaved"

	// Update Notifications
	UpdateAvailableStable = "updateAvailableStable"
	UpdateAvailableBeta   = "updateAvailableBeta"
)

var translations = map[string]map[string]string{
	LangDE_DE: {
		// Tabs
		TabAccounts: "📋 Accounts",
		TabAdd:      "➕ Hinzufügen",
		TabSettings: "⚙️ Einstellungen",

		// Add Account Tab
		AddAccountTitle:    "Neuen Account hinzufügen",
		LabelAccountName:   "Account Name (z.B. \"Main\", \"Alt\")",
		PlaceholderAccName: "Main Account",
		LabelEmail:         "Email",
		PlaceholderEmail:   "your@email.com",
		AddAccountHelp:     "Nach dem Hinzufügen startet der Launcher automatisch.\nLogge dich ein - die Session wird automatisch gespeichert! ✅",
		BtnAddAccount:      "Account hinzufügen & Launcher starten",

		// Account List
		EmptyStateTitle:    "Noch keine Accounts gespeichert",
		EmptyStateSubtitle: "Füge oben deinen ersten Account hinzu",
		StatusAutoLogin:    "Auto-Login aktiv",
		StatusLoginReq:     "Login erforderlich - wird automatisch gespeichert",
		BtnSwitch:          "Wechseln",
		BtnDelete:          "Löschen",
		ConfirmDelete:      "Account wirklich löschen?",

		// Settings Tab
		SettingsTitle:       "Einstellungen",
		LabelLanguage:       "Sprache / Language",
		LabelLauncherPath:   "BSG Launcher Pfad",
		PlaceholderLauncher: `C:\Battlestate Games\BsgLauncher\BsgLauncher.exe`,
		BtnBrowse:           "Durchsuchen...",
		BtnSave:             "Speichern",
		LabelStreamerMode:   "Streamer Modus",
		StreamerModeHelp:    "Versteckt Email-Adressen mit ****",
		LabelTheme:          "Design / Theme",
		LabelAutoStart:      "Autostart mit Windows",
		AutoStartHelp:       "Startet die App automatisch beim Windows-Login",
		BtnQuit:             "Beenden",

		// Switch Result Messages
		SwitchAutoLogin:   "Launcher gestartet - Auto-Login aktiv!",
		SwitchManualLogin: "Bitte einloggen - Session wird automatisch gespeichert!",

		// Tray Menu
		TrayOpen: "Öffnen",
		TrayQuit: "Beenden",

		// Status Messages
		StatusFillFields:      "Bitte fülle alle Felder aus",
		StatusAccountAdded:    "✅ Account hinzugefügt!\n\nLauncher startet jetzt...\nBitte einloggen - Session wird automatisch gespeichert!",
		StatusAccountDeleted:  "Account gelöscht",
		StatusDeleteError:     "Fehler beim Löschen",
		StatusLauncherRestart: "Launcher wird neu gestartet...",
		StatusAutoLoginActive: "🚀 AUTO-LOGIN AKTIV!\n\nAccount: {name}\nLauncher startet automatisch eingeloggt!",
		StatusManualLogin:     "⚠️ MANUELLES LOGIN\n\nAccount: {name}\nEmail: {email}\n\nBitte einloggen - Session wird automatisch gespeichert!",
		StatusError:           "Fehler: {error}",
		StatusPathSaved:       "Launcher Pfad gespeichert!",
		StatusSaveError:       "Fehler beim Speichern",
		StatusEnterPath:       "Bitte gib einen Pfad ein",
		StatusLanguageSaved:   "Sprache gespeichert!",

		// Update Notifications
		UpdateAvailableStable: "Update verfügbar: {version} — {url}",
		UpdateAvailableBeta:   "Neue Beta verfügbar: {version} — {url}",
	},
	LangEN_US: {
		// Tabs
		TabAccounts: "📋 Accounts",
		TabAdd:      "➕ Add",
		TabSettings: "⚙️ Settings",

		// Add Account Tab
		AddAccountTitle:    "Add New Account",
		LabelAccountName:   "Account Name (e.g. \"Main\", \"Alt\")",
		PlaceholderAccName: "Main Account",
		LabelEmail:         "Email",
		PlaceholderEmail:   "your@email.com",
		AddAccountHelp:     "After adding, the launcher will start automatically.\nLog in - the session will be saved automatically! ✅",
		BtnAddAccount:      "Add Account & Start Launcher",

		// Account List
		EmptyStateTitle:    "No accounts saved yet",
		EmptyStateSubtitle: "Add your first account above",
		StatusAutoLogin:    "Auto-login active",
		StatusLoginReq:     "Login required - will be saved automatically",
		BtnSwitch:          "Switch",
		BtnDelete:          "Delete",
		ConfirmDelete:      "Really delete account?",

		// Settings Tab
		SettingsTitle:       "Settings",
		LabelLanguage:       "Language / Sprache",
		LabelLauncherPath:   "BSG Launcher Path",
		PlaceholderLauncher: `C:\Battlestate Games\BsgLauncher\BsgLauncher.exe`,
		BtnBrowse:           "Browse...",
		BtnSave:             "Save",
		LabelStreamerMode:   "Streamer Mode",
		StreamerModeHelp:    "Hides email addresses with ****",
		LabelTheme:          "Theme / Design",
		LabelAutoStart:      "Start with Windows",
		AutoStartHelp:       "Automatically start the app on Windows login",
		BtnQuit:             "Quit",

		// Switch Result Messages
		SwitchAutoLogin:   "Launcher started - Auto-login active!",
		SwitchManualLogin: "Please login - session will be saved automatically!",

		// Tray Menu
		TrayOpen: "Open",
		TrayQuit: "Quit",

		// Status Messages
		StatusFillFields:      "Please fill all fields",
		StatusAccountAdded:    "✅ Account added!\n\nLauncher starting...\nPlease login - session will be saved automatically!",
		StatusAccountDeleted:  "Account deleted",
		StatusDeleteError:     "Error deleting",
		StatusLauncherRestart: "Restarting launcher...",
		StatusAutoLoginActive: "🚀 AUTO-LOGIN ACTIVE!\n\nAccount: {name}\nLauncher starting automatically logged in!",
		StatusManualLogin:     "⚠️ MANUAL LOGIN\n\nAccount: {name}\nEmail: {email}\n\nPlease login - session will be saved automatically!",
		StatusError:           "Error: {error}",
		StatusPathSaved:       "Launcher path saved!",
		StatusSaveError:       "Error saving",
		StatusEnterPath:       "Please enter a path",
		StatusLanguageSaved:   "Language saved!",

		// Update Notifications
		UpdateAvailableStable: "Update available: {version} — {url}",
		UpdateAvailableBeta:   "Beta available: {version} — {url}",
	},
	LangZH_TW: {
		TabAccounts: "📋 帳號",
		TabAdd:      "➕ 新增",
		TabSettings: "⚙️ 設定",

		AddAccountTitle:    "新增帳號",
		LabelAccountName:   "帳號名稱（例：「主帳」、「分身」）",
		PlaceholderAccName: "主帳號",
		LabelEmail:         "電子郵件",
		PlaceholderEmail:   "your@email.com",
		AddAccountHelp:     "新增後啟動器會自動啟動。\n請登入——工作階段會自動儲存！ ✅",
		BtnAddAccount:      "新增帳號並啟動啟動器",

		EmptyStateTitle:    "尚無已儲存的帳號",
		EmptyStateSubtitle: "在上方新增你的第一個帳號",
		StatusAutoLogin:    "自動登入中",
		StatusLoginReq:     "需要登入——將自動儲存",
		BtnSwitch:          "切換",
		BtnDelete:          "刪除",
		ConfirmDelete:      "確定要刪除此帳號？",

		SettingsTitle:       "設定",
		LabelLanguage:       "語言",
		LabelLauncherPath:   "BSG 啟動器路徑",
		PlaceholderLauncher: `C:\Battlestate Games\BsgLauncher\BsgLauncher.exe`,
		BtnBrowse:           "瀏覽…",
		BtnSave:             "儲存",
		LabelStreamerMode:   "實況模式",
		StreamerModeHelp:    "以 **** 隱藏電子郵件",
		LabelTheme:          "主題",
		LabelAutoStart:      "隨 Windows 啟動",
		AutoStartHelp:       "登入 Windows 時自動啟動此程式",
		BtnQuit:             "結束",

		SwitchAutoLogin:   "啟動器已啟動——自動登入中！",
		SwitchManualLogin: "請登入——工作階段會自動儲存！",

		TrayOpen: "開啟",
		TrayQuit: "結束",

		StatusFillFields:      "請填寫所有欄位",
		StatusAccountAdded:    "✅ 已新增帳號！\n\n正在啟動啟動器…\n請登入——工作階段會自動儲存！",
		StatusAccountDeleted:  "已刪除帳號",
		StatusDeleteError:     "刪除時發生錯誤",
		StatusLauncherRestart: "正在重新啟動啟動器…",
		StatusAutoLoginActive: "🚀 自動登入啟用！\n\n帳號：{name}\n啟動器將以已登入狀態啟動！",
		StatusManualLogin:     "⚠️ 手動登入\n\n帳號：{name}\n電子郵件：{email}\n\n請登入——工作階段會自動儲存！",
		StatusError:           "錯誤：{error}",
		StatusPathSaved:       "已儲存啟動器路徑！",
		StatusSaveError:       "儲存時發生錯誤",
		StatusEnterPath:       "請輸入路徑",
		StatusLanguageSaved:   "已儲存語言！",

		UpdateAvailableStable: "有可用更新：{version} — {url}",
		UpdateAvailableBeta:   "有可用 Beta：{version} — {url}",
	},
	LangZH_CN: {
		TabAccounts: "📋 账号",
		TabAdd:      "➕ 添加",
		TabSettings: "⚙️ 设置",

		AddAccountTitle:    "添加新账号",
		LabelAccountName:   "账号名称（例如「主号」「小号」）",
		PlaceholderAccName: "主账号",
		LabelEmail:         "电子邮箱",
		PlaceholderEmail:   "your@email.com",
		AddAccountHelp:     "添加后启动器会自动启动。\n请登录——会话会自动保存！ ✅",
		BtnAddAccount:      "添加账号并启动启动器",

		EmptyStateTitle:    "尚无已保存的账号",
		EmptyStateSubtitle: "在上方添加你的第一个账号",
		StatusAutoLogin:    "自动登录已启用",
		StatusLoginReq:     "需要登录——将自动保存",
		BtnSwitch:          "切换",
		BtnDelete:          "删除",
		ConfirmDelete:      "确定要删除此账号？",

		SettingsTitle:       "设置",
		LabelLanguage:       "语言",
		LabelLauncherPath:   "BSG 启动器路径",
		PlaceholderLauncher: `C:\Battlestate Games\BsgLauncher\BsgLauncher.exe`,
		BtnBrowse:           "浏览…",
		BtnSave:             "保存",
		LabelStreamerMode:   "主播模式",
		StreamerModeHelp:    "用 **** 隐藏电子邮箱地址",
		LabelTheme:          "主题",
		LabelAutoStart:      "随 Windows 启动",
		AutoStartHelp:       "登录 Windows 时自动启动本程序",
		BtnQuit:             "退出",

		SwitchAutoLogin:   "启动器已启动——自动登录已启用！",
		SwitchManualLogin: "请登录——会话会自动保存！",

		TrayOpen: "打开",
		TrayQuit: "退出",

		StatusFillFields:      "请填写所有字段",
		StatusAccountAdded:    "✅ 已添加账号！\n\n正在启动启动器…\n请登录——会话会自动保存！",
		StatusAccountDeleted:  "已删除账号",
		StatusDeleteError:     "删除时出错",
		StatusLauncherRestart: "正在重新启动启动器…",
		StatusAutoLoginActive: "🚀 自动登录已启用！\n\n账号：{name}\n启动器将以已登录状态启动！",
		StatusManualLogin:     "⚠️ 手动登录\n\n账号：{name}\n电子邮箱：{email}\n\n请登录——会话会自动保存！",
		StatusError:           "错误：{error}",
		StatusPathSaved:       "已保存启动器路径！",
		StatusSaveError:       "保存时出错",
		StatusEnterPath:       "请输入路径",
		StatusLanguageSaved:   "已保存语言！",

		UpdateAvailableStable: "有可用更新：{version} — {url}",
		UpdateAvailableBeta:   "有可用 Beta：{version} — {url}",
	},
	LangJA_JP: {
		TabAccounts: "📋 アカウント",
		TabAdd:      "➕ 追加",
		TabSettings: "⚙️ 設定",

		AddAccountTitle:    "新しいアカウントを追加",
		LabelAccountName:   "アカウント名（例:「メイン」「サブ」）",
		PlaceholderAccName: "メインアカウント",
		LabelEmail:         "メール",
		PlaceholderEmail:   "your@email.com",
		AddAccountHelp:     "追加後、ランチャーが自動で起動します。\nログインするとセッションが自動保存されます！ ✅",
		BtnAddAccount:      "アカウントを追加してランチャーを起動",

		EmptyStateTitle:    "保存されたアカウントはありません",
		EmptyStateSubtitle: "上で最初のアカウントを追加してください",
		StatusAutoLogin:    "自動ログイン有効",
		StatusLoginReq:     "ログインが必要です（自動保存されます）",
		BtnSwitch:          "切替",
		BtnDelete:          "削除",
		ConfirmDelete:      "このアカウントを削除しますか？",

		SettingsTitle:       "設定",
		LabelLanguage:       "言語",
		LabelLauncherPath:   "BSG ランチャーのパス",
		PlaceholderLauncher: `C:\Battlestate Games\BsgLauncher\BsgLauncher.exe`,
		BtnBrowse:           "参照…",
		BtnSave:             "保存",
		LabelStreamerMode:   "ストリーマーモード",
		StreamerModeHelp:    "メールアドレスを **** で隠します",
		LabelTheme:          "テーマ",
		LabelAutoStart:      "Windows 起動時に開始",
		AutoStartHelp:       "Windows ログイン時にこのアプリを自動起動",
		BtnQuit:             "終了",

		SwitchAutoLogin:   "ランチャーを起動しました — 自動ログイン有効！",
		SwitchManualLogin: "ログインしてください — セッションは自動保存されます！",

		TrayOpen: "開く",
		TrayQuit: "終了",

		StatusFillFields:      "すべての項目を入力してください",
		StatusAccountAdded:    "✅ アカウントを追加しました！\n\nランチャーを起動中…\nログインするとセッションが自動保存されます！",
		StatusAccountDeleted:  "アカウントを削除しました",
		StatusDeleteError:     "削除できませんでした",
		StatusLauncherRestart: "ランチャーを再起動しています…",
		StatusAutoLoginActive: "🚀 自動ログイン有効！\n\nアカウント: {name}\nランチャーはログイン済みで起動します！",
		StatusManualLogin:     "⚠️ 手動ログイン\n\nアカウント: {name}\nメール: {email}\n\nログインするとセッションが自動保存されます！",
		StatusError:           "エラー: {error}",
		StatusPathSaved:       "ランチャーのパスを保存しました！",
		StatusSaveError:       "保存できませんでした",
		StatusEnterPath:       "パスを入力してください",
		StatusLanguageSaved:   "言語を保存しました！",

		UpdateAvailableStable: "アップデートあり: {version} — {url}",
		UpdateAvailableBeta:   "新しいベータ: {version} — {url}",
	},
}

// IsSupportedLocale reports whether lang is one of the canonical locale tags.
func IsSupportedLocale(lang string) bool {
	switch lang {
	case LangEN_US, LangDE_DE, LangZH_TW, LangZH_CN, LangJA_JP:
		return true
	default:
		return false
	}
}

// SetLanguage sets the entire UI language.
func SetLanguage(lang string) {
	switch lang {
	case LangEN_US, LangDE_DE, LangZH_TW, LangZH_CN, LangJA_JP:
		languageMutex.Lock()
		currentLanguage = lang
		languageMutex.Unlock()
	}
}

// GetLanguage returns the current language
func GetLanguage() string {
	languageMutex.RLock()
	defer languageMutex.RUnlock()
	return currentLanguage
}

// T returns the translation for the given key
func T(key string) string {
	languageMutex.RLock()
	defer languageMutex.RUnlock()

	if trans, ok := translations[currentLanguage]; ok {
		if val, ok := trans[key]; ok {
			return val
		}
	}

	// Fallback to English
	if trans, ok := translations[LangEN_US]; ok {
		if val, ok := trans[key]; ok {
			return val
		}
	}

	return key
}

// TF returns the translation with placeholders replaced
func TF(key string, replacements map[string]string) string {
	text := T(key)
	for placeholder, value := range replacements {
		text = strings.ReplaceAll(text, "{"+placeholder+"}", value)
	}
	return text
}
