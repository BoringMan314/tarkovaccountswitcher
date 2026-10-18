package launcher

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"tarkov-account-switcher/internal/config"
)

// UpdateLauncherAccount updates the launcher settings with a new account email and clears session
func UpdateLauncherAccount(email string) error {
	paths := config.GetPaths()
	launcherDataPath := filepath.Dir(paths.LauncherSettingsPath)

	// Read existing settings or create new
	var settings map[string]any

	data, err := os.ReadFile(paths.LauncherSettingsPath)
	if err == nil {
		json.Unmarshal(data, &settings)
	}

	if settings == nil {
		settings = make(map[string]any)
	}

	// Update login email and settings
	settings["login"] = email
	settings["saveLogin"] = true
	settings["keepLoggedIn"] = true
	settings["tempFolder"] = paths.TempFolder

	// CRITICAL: Delete session tokens to force fresh login
	delete(settings, "at")
	delete(settings, "rt")
	delete(settings, "atet")

	// Ensure directory exists
	if err := os.MkdirAll(launcherDataPath, 0755); err != nil {
		return err
	}

	// Write settings
	settingsData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(paths.LauncherSettingsPath, settingsData, 0644); err != nil {
		return err
	}

	clearLauncherSessionFiles(launcherDataPath)

	return nil
}

// clearLauncherSessionFiles removes separate launcher auth caches that can
// override the credentials in settings. This matters when Steam and the
// standalone launcher have both written to the shared launcher data folder.
func clearLauncherSessionFiles(launcherDataPath string) {
	sessionFiles := []string{
		"user.json",
		"session",
		"token",
		".session",
		"auth",
		"auth.json",
		"login",
		"login.json",
	}

	for _, file := range sessionFiles {
		filePath := filepath.Join(launcherDataPath, file)
		os.Remove(filePath) // Ignore errors
	}
}

// RestoreLauncherSession restores a saved launcher session
func RestoreLauncherSession(sessionData json.RawMessage) error {
	paths := config.GetPaths()

	// Parse saved session
	var savedSession map[string]any
	if err := json.Unmarshal(sessionData, &savedSession); err != nil {
		return err
	}

	// Read existing launcher settings to preserve game state
	var existingSettings map[string]any

	data, err := os.ReadFile(paths.LauncherSettingsPath)
	if err == nil {
		json.Unmarshal(data, &existingSettings)
	}

	if existingSettings == nil {
		existingSettings = make(map[string]any)
		// Copy all from saved session as base
		for k, v := range savedSession {
			existingSettings[k] = v
		}
	}

	// Remove leftover Steam/launcher auth caches before writing the selected
	// account's saved tokens. Otherwise the launcher can prefer stale cached
	// credentials and prompt for the password again.
	clearLauncherSessionFiles(filepath.Dir(paths.LauncherSettingsPath))

	// Remove auth-related fields left by Steam or another launcher session.
	// The settings schema can change between launcher versions, so preserve
	// non-auth preferences while dropping fields whose names indicate identity
	// or credentials.
	for field := range existingSettings {
		if isLauncherAuthSetting(field) {
			delete(existingSettings, field)
		}
	}

	// Restore the selected account's auth fields and selected game, while
	// keeping non-auth game state and UI preferences from the current settings.
	authFields := []string{"login", "at", "rt", "atet", "keepLoggedIn", "saveLogin", "selectedGame", "sysInfCheck"}
	for _, field := range authFields {
		if val, ok := savedSession[field]; ok {
			existingSettings[field] = val
		}
	}

	// ALWAYS use our own temp folder
	existingSettings["tempFolder"] = paths.TempFolder

	// Write settings
	settingsData, err := json.MarshalIndent(existingSettings, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(paths.LauncherSettingsPath, settingsData, 0644); err != nil {
		return err
	}

	// Restore ingame background (EnvironmentUiType) to Game.ini
	if envType, ok := savedSession["environmentUiType"].(string); ok && envType != "" {
		WriteEnvironmentUiType(envType)
	}

	return nil
}

func isLauncherAuthSetting(field string) bool {
	key := strings.ToLower(field)
	for _, marker := range []string{"auth", "token", "session", "credential", "identity", "steam", "password", "provider", "login", "account"} {
		if strings.Contains(key, marker) {
			return true
		}
	}
	return key == "user" || strings.HasPrefix(key, "user_") || strings.HasSuffix(key, "_user")
}

// GetGameSettingsPath returns the path to the EFT Game.ini file
func GetGameSettingsPath() string {
	appData := os.Getenv("APPDATA")
	return filepath.Join(appData, "Battlestate Games", "Escape from Tarkov", "Settings", "Game.ini")
}

// ReadEnvironmentUiType reads the current EnvironmentUiType (ingame background) from Game.ini
func ReadEnvironmentUiType() string {
	data, err := os.ReadFile(GetGameSettingsPath())
	if err != nil {
		return ""
	}

	var gameSettings map[string]any
	if err := json.Unmarshal(data, &gameSettings); err != nil {
		return ""
	}

	if envType, ok := gameSettings["EnvironmentUiType"].(string); ok {
		return envType
	}
	return ""
}

// WriteEnvironmentUiType writes the EnvironmentUiType (ingame background) to Game.ini
func WriteEnvironmentUiType(envType string) error {
	if envType == "" {
		return nil // Nothing to restore
	}

	gameSettingsPath := GetGameSettingsPath()

	// Read existing settings
	data, err := os.ReadFile(gameSettingsPath)
	if err != nil {
		return err // Game.ini must exist
	}

	var gameSettings map[string]any
	if err := json.Unmarshal(data, &gameSettings); err != nil {
		return err
	}

	// Update EnvironmentUiType
	gameSettings["EnvironmentUiType"] = envType

	// Write back
	settingsData, err := json.MarshalIndent(gameSettings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(gameSettingsPath, settingsData, 0644)
}
