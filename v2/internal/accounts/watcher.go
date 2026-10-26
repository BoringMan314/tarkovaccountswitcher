package accounts

import (
	"strings"
	"sync"
	"time"

	"tarkov-account-switcher/internal/launcher"
)

var (
	watchMu         sync.Mutex
	watchDeadline   time.Time
	watchRunning    bool
	watchGeneration uint64
	activeAccount   string

	// SessionCapturedCallback is called when a session is captured.
	SessionCapturedCallback func(accountID string)
)

// WatchLauncherSession watches briefly for a new or rotated session. The
// launcher settings file is shared by the official and Steam launchers, so
// capture only while the switcher-started configured launcher is the sole
// launcher installation running.
func WatchLauncherSession(accountID string) {
	watchMu.Lock()
	watchGeneration++
	activeAccount = accountID
	watchDeadline = time.Now().Add(5 * time.Minute)
	if !watchRunning {
		watchRunning = true
		go watchLauncherSessionLoop()
	}
	watchMu.Unlock()
}

// StopLauncherSessionWatch prevents a previous launch from being associated
// with the account selected by the next switch.
func StopLauncherSessionWatch() {
	watchMu.Lock()
	watchGeneration++
	activeAccount = ""
	watchMu.Unlock()
}

// CurrentAccountID reports the account associated with the active launcher.
func CurrentAccountID() string {
	watchMu.Lock()
	defer watchMu.Unlock()
	return activeAccount
}

func watchLauncherSessionLoop() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		watchMu.Lock()
		if time.Now().After(watchDeadline) {
			watchRunning = false
			watchMu.Unlock()
			return
		}
		accountID, generation := activeAccount, watchGeneration
		watchMu.Unlock()

		if accountID == "" || !launcher.IsOnlyConfiguredLauncherRunning() || !launcher.IsSwitcherStartedLauncherRunning() {
			continue
		}
		settings, login, err := readLauncherAuth()
		if err != nil || settings == nil {
			continue
		}
		account, err := GetAccountByID(accountID)
		if err != nil || account == nil || !strings.EqualFold(strings.TrimSpace(login), strings.TrimSpace(account.Email)) {
			continue
		}

		watchMu.Lock()
		if activeAccount != accountID || watchGeneration != generation {
			watchMu.Unlock()
			continue
		}
		capturedID := storeLauncherSession(settings, accountID)
		watchMu.Unlock()
		if capturedID != "" && SessionCapturedCallback != nil {
			SessionCapturedCallback(capturedID)
		}
	}
}
